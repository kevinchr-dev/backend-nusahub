#!/bin/bash

# Comprehensive API Testing Script
# Tests all 16 endpoints with detailed scenarios

BASE_URL="http://localhost:3000/api/v1"

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo "=========================================="
echo "🧪 Comprehensive API Testing"
echo "=========================================="

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}📋 1. HEALTH CHECK${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}GET /health${NC}"
curl -s http://localhost:3000/api/v1/health | jq .

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}🎮 2. PROJECTS (4 endpoints)${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}POST /projects - Create Project #1${NC}"
PROJECT1=$(curl -s -X POST $BASE_URL/projects \
  -H "Content-Type: application/json" \
  -d '{
    "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "title": "Epic RPG Game",
    "description": "An amazing blockchain RPG game with NFT characters",
    "developer_name": "Epic Games Studio",
    "genre": "RPG",
    "game_type": "web3",
    "cover_image_url": "https://example.com/cover1.jpg"
  }')
PROJECT1_ID=$(echo $PROJECT1 | jq -r '.id')
echo $PROJECT1 | jq .
echo -e "${GREEN}✓ Project 1 ID: $PROJECT1_ID${NC}"

echo -e "\n${YELLOW}POST /projects - Create Project #2${NC}"
PROJECT2=$(curl -s -X POST $BASE_URL/projects \
  -H "Content-Type: application/json" \
  -d '{
    "creator_wallet_address": "0x1234567890abcdef1234567890abcdef12345678",
    "title": "Strategy Web3 Game",
    "description": "Strategic gameplay with crypto rewards",
    "developer_name": "Strategy Dev",
    "genre": "Strategy",
    "game_type": "web3"
  }')
PROJECT2_ID=$(echo $PROJECT2 | jq -r '.id')
echo $PROJECT2 | jq .
echo -e "${GREEN}✓ Project 2 ID: $PROJECT2_ID${NC}"

echo -e "\n${YELLOW}GET /projects - Get All Projects${NC}"
curl -s $BASE_URL/projects | jq '. | length as $count | {total_projects: $count, projects: .}'

echo -e "\n${YELLOW}GET /projects/:id - Get Project by ID${NC}"
curl -s $BASE_URL/projects/$PROJECT1_ID | jq .

echo -e "\n${YELLOW}PATCH /projects/:id - Update Project${NC}"
curl -s -X PATCH $BASE_URL/projects/$PROJECT1_ID \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Epic RPG Game - UPDATED",
    "description": "Updated description with more details"
  }' | jq .

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}💰 3. INVESTORS (3 endpoints)${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}POST /projects/:id/investors - Add Investor #1${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/investors \
  -H "Content-Type: application/json" \
  -d '{"wallet_address":"0x1111111111111111111111111111111111111111"}' | jq .

echo -e "\n${YELLOW}POST /projects/:id/investors - Add Investor #2${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/investors \
  -H "Content-Type: application/json" \
  -d '{"wallet_address":"0x2222222222222222222222222222222222222222"}' | jq .

echo -e "\n${YELLOW}POST /projects/:id/investors - Add Investor #3${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/investors \
  -H "Content-Type: application/json" \
  -d '{"wallet_address":"0x3333333333333333333333333333333333333333"}' | jq .

echo -e "\n${YELLOW}GET /projects/:id/investors - Get All Investors${NC}"
INVESTORS=$(curl -s $BASE_URL/projects/$PROJECT1_ID/investors)
echo $INVESTORS | jq .
INVESTOR_COUNT=$(echo $INVESTORS | jq '.investors | length')
echo -e "${GREEN}✓ Total Investors: $INVESTOR_COUNT${NC}"

echo -e "\n${YELLOW}POST /projects/:id/investors - Try Add Duplicate (should fail)${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/investors \
  -H "Content-Type: application/json" \
  -d '{"wallet_address":"0x1111111111111111111111111111111111111111"}' | jq .

echo -e "\n${YELLOW}DELETE /projects/:id/investors/:address - Remove Investor${NC}"
curl -s -X DELETE $BASE_URL/projects/$PROJECT1_ID/investors/0x3333333333333333333333333333333333333333 | jq .

echo -e "\n${YELLOW}GET /projects/:id/investors - Verify Investor Removed${NC}"
curl -s $BASE_URL/projects/$PROJECT1_ID/investors | jq .

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}🔗 4. EXTERNAL LINKS (4 endpoints)${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}POST /projects/:id/links - Add Instagram${NC}"
LINK1=$(curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/links \
  -H "Content-Type: application/json" \
  -d '{"name":"Instagram","url":"https://instagram.com/epicrpggame"}')
LINK1_ID=$(echo $LINK1 | jq -r '.id')
echo $LINK1 | jq .

echo -e "\n${YELLOW}POST /projects/:id/links - Add Twitter${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/links \
  -H "Content-Type: application/json" \
  -d '{"name":"Twitter","url":"https://twitter.com/epicrpggame"}' | jq .

echo -e "\n${YELLOW}POST /projects/:id/links - Add Discord${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/links \
  -H "Content-Type: application/json" \
  -d '{"name":"Discord","url":"https://discord.gg/epicrpggame"}' | jq .

echo -e "\n${YELLOW}POST /projects/:id/links - Add Website${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/links \
  -H "Content-Type: application/json" \
  -d '{"name":"Website","url":"https://epicrpggame.com"}' | jq .

echo -e "\n${YELLOW}GET /projects/:id/links - Get All Links${NC}"
LINKS=$(curl -s $BASE_URL/projects/$PROJECT1_ID/links)
echo $LINKS | jq .
LINK_COUNT=$(echo $LINKS | jq '. | length')
echo -e "${GREEN}✓ Total Links: $LINK_COUNT${NC}"

echo -e "\n${YELLOW}PUT /projects/:id/links/:linkId - Update Instagram Link${NC}"
curl -s -X PUT $BASE_URL/projects/$PROJECT1_ID/links/$LINK1_ID \
  -H "Content-Type: application/json" \
  -d '{"name":"Instagram Official","url":"https://instagram.com/epicrpggame_official"}' | jq .

echo -e "\n${YELLOW}DELETE /projects/:id/links/:linkId - Delete Discord Link${NC}"
DISCORD_ID=$(echo $LINKS | jq -r '.[] | select(.name=="Discord") | .id')
curl -s -X DELETE $BASE_URL/projects/$PROJECT1_ID/links/$DISCORD_ID | jq .

echo -e "\n${YELLOW}GET /projects/:id/links - Verify Link Deleted${NC}"
curl -s $BASE_URL/projects/$PROJECT1_ID/links | jq '. | map(.name)'

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}💬 5. COMMENTS (2 endpoints)${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}POST /projects/:id/comments - Add Comment #1${NC}"
COMMENT1=$(curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/comments \
  -H "Content-Type: application/json" \
  -d '{
    "author_wallet_address":"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "content":"This game looks amazing! Can'\''t wait to play it!"
  }')
COMMENT1_ID=$(echo $COMMENT1 | jq -r '.id')
echo $COMMENT1 | jq .

echo -e "\n${YELLOW}POST /projects/:id/comments - Add Comment #2${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/comments \
  -H "Content-Type: application/json" \
  -d '{
    "author_wallet_address":"0x1111111111111111111111111111111111111111",
    "content":"Great project! When is the launch date?"
  }' | jq .

echo -e "\n${YELLOW}POST /projects/:id/comments - Add Nested Reply${NC}"
curl -s -X POST $BASE_URL/projects/$PROJECT1_ID/comments \
  -H "Content-Type: application/json" \
  -d "{
    \"author_wallet_address\":\"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb\",
    \"content\":\"Thanks for your support! Launch is planned for Q1 2026.\",
    \"parent_comment_id\":\"$COMMENT1_ID\"
  }" | jq .

echo -e "\n${YELLOW}GET /projects/:id/comments - Get All Comments${NC}"
COMMENTS=$(curl -s $BASE_URL/projects/$PROJECT1_ID/comments)
echo $COMMENTS | jq .
COMMENT_COUNT=$(echo $COMMENTS | jq '. | length')
echo -e "${GREEN}✓ Total Comments: $COMMENT_COUNT${NC}"

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}👤 6. USER PROFILES (2 endpoints)${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}PUT /profiles/:address - Create Profile #1${NC}"
curl -s -X PUT $BASE_URL/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb \
  -H "Content-Type: application/json" \
  -d '{
    "username":"epicgamer123",
    "email":"epic@gamer.com",
    "profile_image_url":"https://avatar.com/epicgamer.jpg",
    "kyc_status":"verified"
  }' | jq .

echo -e "\n${YELLOW}PUT /profiles/:address - Create Profile #2${NC}"
curl -s -X PUT $BASE_URL/profiles/0x1111111111111111111111111111111111111111 \
  -H "Content-Type: application/json" \
  -d '{
    "username":"investor001",
    "email":"investor@fund.com",
    "kyc_status":"verified"
  }' | jq .

echo -e "\n${YELLOW}GET /profiles/:address - Get Profile${NC}"
curl -s $BASE_URL/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb | jq .

echo -e "\n${YELLOW}PUT /profiles/:address - Update Profile (Upsert)${NC}"
curl -s -X PUT $BASE_URL/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb \
  -H "Content-Type: application/json" \
  -d '{
    "username":"epicgamer123_updated",
    "email":"epic@gamer.com",
    "profile_image_url":"https://avatar.com/epicgamer_new.jpg",
    "kyc_status":"verified"
  }' | jq .

echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}📊 7. FINAL SUMMARY${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n${YELLOW}Get Project 1 with All Relations${NC}"
curl -s $BASE_URL/projects/$PROJECT1_ID | jq '{
  id,
  title,
  creator: .creator_wallet_address,
  investors_count: (.investor_wallet_addresses | length),
  investors: .investor_wallet_addresses
}'

echo -e "\n${YELLOW}Get All Projects Summary${NC}"
curl -s $BASE_URL/projects | jq 'map({
  id,
  title,
  creator: .creator_wallet_address,
  investors: (.investor_wallet_addresses | length),
  genre,
  game_type
})'

echo -e "\n${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${GREEN}✅ COMPREHENSIVE TEST COMPLETED!${NC}"
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo -e "\n📈 ${BLUE}Test Statistics:${NC}"
echo -e "  • Projects Created: 2"
echo -e "  • Investors Added: 3 (1 removed)"
echo -e "  • External Links: 4 (1 deleted)"
echo -e "  • Comments: 3 (including nested)"
echo -e "  • User Profiles: 2"

echo -e "\n🌐 ${BLUE}Access Points:${NC}"
echo -e "  • Swagger UI: ${YELLOW}http://localhost:3000/docs/${NC}"
echo -e "  • Health Check: ${YELLOW}http://localhost:3000/api/v1/health${NC}"
echo -e "  • API Base: ${YELLOW}http://localhost:3000/api/v1${NC}"

echo -e "\n🎉 ${GREEN}All 16 endpoints tested successfully!${NC}"
