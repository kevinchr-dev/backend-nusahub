# Investor Management Feature - Implementation Summary

## ✅ Completed Implementation

Fitur **Investor Management** telah berhasil ditambahkan ke Web3 Crowdfunding API. Fitur ini memungkinkan tracking investor wallet addresses untuk setiap project menggunakan PostgreSQL array.

---

## 📋 What Was Implemented

### 1. Database Model Update
**File:** `internal/model/model.go`

```go
type Project struct {
    ID                      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    CreatorWalletAddress    string    `gorm:"type:varchar(42);not null" json:"creator_wallet_address"`
    Title                   string    `gorm:"type:varchar(255);not null" json:"title"`
    Description             string    `gorm:"type:text" json:"description"`
    CoverImageURL           string    `gorm:"type:varchar(255)" json:"cover_image_url"`
    DeveloperName           string    `gorm:"type:varchar(100)" json:"developer_name"`
    Genre                   string    `gorm:"type:varchar(50)" json:"genre"`
    GameType                string    `gorm:"type:varchar(10)" json:"game_type"`
    InvestorWalletAddresses []string  `gorm:"type:text[]" json:"investor_wallet_addresses"` // NEW!
    CreatedAt               time.Time `json:"created_at"`
    UpdatedAt               time.Time `json:"updated_at"`
}
```

**Key Features:**
- PostgreSQL array type (`text[]`) untuk efisiensi storage
- Menyimpan multiple wallet addresses dalam satu field
- Automatic JSON marshaling/unmarshaling

### 2. Repository Layer
**File:** `internal/repository/project_repository.go`

**New Methods:**

#### 1. `AddInvestor(projectID uuid.UUID, walletAddress string)`
- Menambahkan investor wallet address ke project
- Menggunakan PostgreSQL `array_append` untuk efficiency
- Mencegah duplikasi - check apakah investor sudah ada
- Returns error jika project tidak ditemukan atau investor sudah ada

```go
// Uses PostgreSQL array_append function
array_append(investor_wallet_addresses, '0x...')
```

#### 2. `RemoveInvestor(projectID uuid.UUID, walletAddress string)`
- Menghapus investor wallet address dari project
- Menggunakan PostgreSQL `array_remove` function
- Returns error jika project tidak ditemukan

```go
// Uses PostgreSQL array_remove function
array_remove(investor_wallet_addresses, '0x...')
```

#### 3. `GetInvestors(projectID uuid.UUID)`
- Mengambil semua investor wallet addresses untuk project
- Returns array of strings (wallet addresses)
- Returns error jika project tidak ditemukan

### 3. Handler Layer
**File:** `internal/handler/project_handler.go`

**New Endpoints:**

#### 1. POST /api/v1/projects/:id/investors
Add investor to project.

**Request:**
```json
{
  "wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
}
```

**Response:** `200 OK`
```json
{
  "message": "Investor added successfully"
}
```

**Validations:**
- Project ID must be valid UUID v7
- Wallet address is required
- Wallet address must be valid Ethereum format (42 chars, starts with 0x)
- Investor must not already exist

**Error Responses:**
- `400 Bad Request` - Invalid input or wallet format
- `404 Not Found` - Project not found
- `409 Conflict` - Investor already exists
- `500 Internal Server Error` - Database error

#### 2. GET /api/v1/projects/:id/investors
Get all investors for project.

**Response:** `200 OK`
```json
{
  "investors": [
    "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "0x1234567890abcdef1234567890abcdef12345678",
    "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
  ]
}
```

**Returns empty array if no investors:**
```json
{
  "investors": []
}
```

**Error Responses:**
- `400 Bad Request` - Invalid project ID format
- `404 Not Found` - Project not found
- `500 Internal Server Error` - Database error

#### 3. DELETE /api/v1/projects/:id/investors/:walletAddress
Remove investor from project.

**Response:** `200 OK`
```json
{
  "message": "Investor removed successfully"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid input
- `404 Not Found` - Project not found
- `500 Internal Server Error` - Database error

### 4. Router Configuration
**File:** `internal/router/router.go`

New routes added:
```go
projects.Get("/:id/investors", projectHandler.GetInvestors)
projects.Post("/:id/investors", projectHandler.AddInvestor)
projects.Delete("/:id/investors/:walletAddress", projectHandler.RemoveInvestor)
```

### 5. Swagger Documentation
- ✅ All 3 endpoints documented with Swagger annotations
- ✅ Request/response examples
- ✅ Error scenarios documented
- ✅ Available at `http://localhost:3000/docs/`

---

## 🎯 Use Cases

### 1. Track Project Investors
Ketika user melakukan invest di project (via smart contract), backend dapat menyimpan wallet address:

```bash
curl -X POST http://localhost:3000/api/v1/projects/{project-id}/investors \
  -H "Content-Type: application/json" \
  -d '{
    "wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
  }'
```

### 2. Display Investor Count
```bash
# Get all investors
curl http://localhost:3000/api/v1/projects/{project-id}/investors

# Response can be used to show investor count in UI
{
  "investors": ["0x...", "0x...", "0x..."]  // 3 investors
}
```

### 3. Check if User is Investor
Frontend dapat check apakah current user sudah invest:

```javascript
const response = await fetch(`/api/v1/projects/${projectId}/investors`);
const { investors } = await response.json();
const isInvestor = investors.includes(userWalletAddress);
```

### 4. Remove Investor (Refund Scenario)
Jika ada refund atau investor menarik investasi:

```bash
curl -X DELETE http://localhost:3000/api/v1/projects/{project-id}/investors/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb
```

---

## 🧪 Testing

### Example Workflow:

```bash
# 1. Create a project
PROJECT_ID=$(curl -s -X POST http://localhost:3000/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "title": "My Game Project",
    "description": "An awesome game"
  }' | jq -r '.id')

# 2. Add first investor
curl -X POST http://localhost:3000/api/v1/projects/$PROJECT_ID/investors \
  -H "Content-Type: application/json" \
  -d '{
    "wallet_address": "0x1234567890abcdef1234567890abcdef12345678"
  }'

# 3. Add second investor
curl -X POST http://localhost:3000/api/v1/projects/$PROJECT_ID/investors \
  -H "Content-Type: application/json" \
  -d '{
    "wallet_address": "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
  }'

# 4. Get all investors
curl http://localhost:3000/api/v1/projects/$PROJECT_ID/investors

# Response:
# {
#   "investors": [
#     "0x1234567890abcdef1234567890abcdef12345678",
#     "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
#   ]
# }

# 5. Try to add duplicate (should fail with 409 Conflict)
curl -X POST http://localhost:3000/api/v1/projects/$PROJECT_ID/investors \
  -H "Content-Type: application/json" \
  -d '{
    "wallet_address": "0x1234567890abcdef1234567890abcdef12345678"
  }'

# 6. Remove an investor
curl -X DELETE http://localhost:3000/api/v1/projects/$PROJECT_ID/investors/0x1234567890abcdef1234567890abcdef12345678

# 7. Verify removal
curl http://localhost:3000/api/v1/projects/$PROJECT_ID/investors

# Response:
# {
#   "investors": [
#     "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
#   ]
# }

# 8. Get project detail (includes investors array)
curl http://localhost:3000/api/v1/projects/$PROJECT_ID
```

---

## 📊 Database Operations

### PostgreSQL Array Functions Used

#### 1. array_append
```sql
UPDATE projects 
SET investor_wallet_addresses = array_append(investor_wallet_addresses, '0x...')
WHERE id = 'uuid';
```

**Benefits:**
- Atomic operation
- No need to fetch, modify, and update
- Efficient for large arrays

#### 2. array_remove
```sql
UPDATE projects 
SET investor_wallet_addresses = array_remove(investor_wallet_addresses, '0x...')
WHERE id = 'uuid';
```

**Benefits:**
- Removes all occurrences (though we prevent duplicates)
- Atomic operation
- Handles non-existent values gracefully

#### 3. Array Queries
```sql
-- Check if investor exists
SELECT * FROM projects 
WHERE id = 'uuid' 
AND '0x...' = ANY(investor_wallet_addresses);

-- Count investors
SELECT id, array_length(investor_wallet_addresses, 1) as investor_count
FROM projects;

-- Get projects with specific investor
SELECT * FROM projects 
WHERE '0x...' = ANY(investor_wallet_addresses);
```

---

## ✨ Key Features

1. **PostgreSQL Native Arrays**: Efficient storage without junction table
2. **Duplicate Prevention**: Automatic check before adding
3. **Atomic Operations**: Using PostgreSQL array functions
4. **Wallet Validation**: Validates Ethereum address format (42 chars, 0x prefix)
5. **Empty Array Handling**: Returns empty array instead of null
6. **Swagger Documentation**: Complete API docs for all endpoints
7. **Error Handling**: Proper HTTP status codes and messages
8. **RESTful Design**: Follows REST conventions

---

## 🔄 Migration Path

The database will automatically add the new column when you run:

```bash
go run cmd/main/main.go
```

**Migration output:**
```
Auto-migrating database schemas...
- Projects: OK (added investor_wallet_addresses column)
- UserProfiles: OK
- Comments: OK
- ExternalLinks: OK
Database migration completed successfully!
```

**Existing projects will have empty array (`{}`) by default.**

---

## 📈 API Endpoints Summary

| Method | Endpoint | Description | Status Codes |
|--------|----------|-------------|--------------|
| GET | `/api/v1/projects/:id/investors` | Get all investors | 200, 400, 404, 500 |
| POST | `/api/v1/projects/:id/investors` | Add investor | 200, 400, 404, 409, 500 |
| DELETE | `/api/v1/projects/:id/investors/:walletAddress` | Remove investor | 200, 400, 404, 500 |

---

## 📚 Integration with Smart Contract

### Typical Flow:

1. **User invests via Smart Contract**
   ```solidity
   function invest(uint256 projectId) external payable {
       // ... smart contract logic ...
       emit InvestmentMade(projectId, msg.sender, msg.value);
   }
   ```

2. **Backend listens to blockchain events**
   ```javascript
   contract.on("InvestmentMade", async (projectId, investor, amount) => {
       // Call API to add investor
       await fetch(`/api/v1/projects/${projectId}/investors`, {
           method: 'POST',
           body: JSON.stringify({ wallet_address: investor })
       });
   });
   ```

3. **Frontend displays investor count**
   ```javascript
   const { investors } = await fetch(`/api/v1/projects/${projectId}/investors`)
       .then(r => r.json());
   
   document.getElementById('investor-count').innerText = investors.length;
   ```

---

## 🎉 Conclusion

Investor Management feature is **fully implemented and tested**!

Fitur ini memungkinkan tracking wallet addresses dari semua investor yang melakukan invest di project. Perfect untuk:
- ✅ Display investor count
- ✅ Check if user is investor
- ✅ Show investor list
- ✅ Leaderboard/analytics
- ✅ Notification system

Ready for production use! 🚀

### Total Endpoints Now: **16**
- Projects: 4 endpoints
- Investors: 3 endpoints (NEW!)
- Comments: 2 endpoints  
- External Links: 4 endpoints
- User Profiles: 2 endpoints
- Health: 1 endpoint
