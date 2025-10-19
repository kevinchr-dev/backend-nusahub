#!/bin/bash

# Test script untuk verifikasi API deployment
# Run setelah docker-compose up

echo "=========================================="
echo "Testing Web3 Crowdfunding API Deployment"
echo "=========================================="

BASE_URL="http://localhost:3000/api/v1"

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test counter
PASSED=0
FAILED=0

# Function to test endpoint
test_endpoint() {
    local name=$1
    local url=$2
    local method=$3
    local data=$4
    
    echo -e "\n${YELLOW}Testing: $name${NC}"
    echo "URL: $method $url"
    
    if [ -z "$data" ]; then
        response=$(curl -s -w "\n%{http_code}" -X $method "$url")
    else
        response=$(curl -s -w "\n%{http_code}" -X $method "$url" \
            -H "Content-Type: application/json" \
            -d "$data")
    fi
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" -ge 200 ] && [ "$http_code" -lt 300 ]; then
        echo -e "${GREEN}✓ PASSED${NC} (HTTP $http_code)"
        echo "Response: $body" | jq '.' 2>/dev/null || echo "$body"
        PASSED=$((PASSED + 1))
        return 0
    else
        echo -e "${RED}✗ FAILED${NC} (HTTP $http_code)"
        echo "Response: $body"
        FAILED=$((FAILED + 1))
        return 1
    fi
}

# Wait for API to be ready
echo -e "\n${YELLOW}Waiting for API to be ready...${NC}"
for i in {1..30}; do
    if curl -s "$BASE_URL/health" > /dev/null 2>&1; then
        echo -e "${GREEN}✓ API is ready!${NC}"
        break
    fi
    echo -n "."
    sleep 1
done

# 1. Health Check
test_endpoint "Health Check" "$BASE_URL/health" "GET"

# 2. Get All Projects (should return empty array initially)
test_endpoint "Get All Projects" "$BASE_URL/projects" "GET"

# 3. Create Project
echo -e "\n${YELLOW}Creating test project...${NC}"
PROJECT_RESPONSE=$(curl -s -X POST "$BASE_URL/projects" \
    -H "Content-Type: application/json" \
    -d '{
        "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
        "title": "Test Deployment Project",
        "description": "Testing API deployment",
        "developer_name": "Test Developer",
        "genre": "RPG",
        "game_type": "web3"
    }')

PROJECT_ID=$(echo $PROJECT_RESPONSE | jq -r '.id')

if [ "$PROJECT_ID" != "null" ] && [ -n "$PROJECT_ID" ]; then
    echo -e "${GREEN}✓ Project created successfully${NC}"
    echo "Project ID: $PROJECT_ID"
    PASSED=$((PASSED + 1))
else
    echo -e "${RED}✗ Failed to create project${NC}"
    FAILED=$((FAILED + 1))
    PROJECT_ID=""
fi

if [ -n "$PROJECT_ID" ]; then
    # 4. Get Project by ID
    test_endpoint "Get Project by ID" "$BASE_URL/projects/$PROJECT_ID" "GET"
    
    # 5. Add Investor
    test_endpoint "Add Investor" "$BASE_URL/projects/$PROJECT_ID/investors" "POST" \
        '{"wallet_address":"0x1234567890abcdef1234567890abcdef12345678"}'
    
    # 6. Get Investors
    test_endpoint "Get Investors" "$BASE_URL/projects/$PROJECT_ID/investors" "GET"
    
    # 7. Add External Link
    test_endpoint "Add External Link" "$BASE_URL/projects/$PROJECT_ID/links" "POST" \
        '{"name":"Instagram","url":"https://instagram.com/testproject"}'
    
    # 8. Get External Links
    test_endpoint "Get External Links" "$BASE_URL/projects/$PROJECT_ID/links" "GET"
    
    # 9. Create Comment
    test_endpoint "Create Comment" "$BASE_URL/projects/$PROJECT_ID/comments" "POST" \
        '{"author_wallet_address":"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb","content":"Great project!"}'
    
    # 10. Get Comments
    test_endpoint "Get Comments" "$BASE_URL/projects/$PROJECT_ID/comments" "GET"
    
    # 11. Update Project
    test_endpoint "Update Project" "$BASE_URL/projects/$PROJECT_ID" "PATCH" \
        '{"title":"Updated Test Project"}'
fi

# 12. Create User Profile
test_endpoint "Create/Update User Profile" "$BASE_URL/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb" "PUT" \
    '{"username":"testuser","email":"test@example.com"}'

# 13. Get User Profile
test_endpoint "Get User Profile" "$BASE_URL/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb" "GET"

# Summary
echo -e "\n=========================================="
echo -e "Test Summary"
echo -e "=========================================="
echo -e "Passed: ${GREEN}$PASSED${NC}"
echo -e "Failed: ${RED}$FAILED${NC}"
echo -e "Total:  $((PASSED + FAILED))"

if [ $FAILED -eq 0 ]; then
    echo -e "\n${GREEN}✓ All tests passed! API is working correctly!${NC}"
    echo -e "\n${YELLOW}Next steps:${NC}"
    echo "1. Access Swagger UI: http://localhost:3000/docs/"
    echo "2. Test more endpoints interactively"
    echo "3. Integrate with your frontend"
    exit 0
else
    echo -e "\n${RED}✗ Some tests failed. Check the logs above.${NC}"
    echo -e "\n${YELLOW}Troubleshooting:${NC}"
    echo "1. Check API logs: docker-compose logs api"
    echo "2. Verify database connection in .env"
    echo "3. Ensure database is accessible"
    exit 1
fi
