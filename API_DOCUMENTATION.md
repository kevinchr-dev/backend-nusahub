# API Documentation - Web3 Crowdfunding Platform

## Base URL
```
http://localhost:3000/api/v1
```

## Response Format

### Success Response
```json
{
  "id": "uuid",
  "field": "value",
  ...
}
```

### Error Response
```json
{
  "error": "Error message description"
}
```

## HTTP Status Codes

| Code | Description |
|------|-------------|
| 200  | OK - Request successful |
| 201  | Created - Resource created successfully |
| 400  | Bad Request - Invalid input |
| 404  | Not Found - Resource not found |
| 409  | Conflict - Resource already exists |
| 500  | Internal Server Error |

---

## Endpoints

### 1. Health Check

Check API status.

**Endpoint:** `GET /health`

**Response:**
```json
{
  "status": "ok",
  "message": "Web3 Crowdfunding API is running"
}
```

---

## Projects

### 2. Get All Projects

Retrieve list of all projects.

**Endpoint:** `GET /projects`

**Response:** `200 OK`
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "title": "Epic Web3 RPG Game",
    "description": "An immersive blockchain-based RPG",
    "cover_image_url": "https://example.com/image.jpg",
    "developer_name": "GameDev Studios",
    "genre": "RPG",
    "game_type": "web3",
    "created_at": "2025-10-15T10:00:00Z",
    "updated_at": "2025-10-15T10:00:00Z"
  }
]
```

---

### 3. Get Project by ID

Get detailed information about a specific project.

**Endpoint:** `GET /projects/{id}`

**URL Parameters:**
- `id` (UUID, required) - Project ID

**Response:** `200 OK`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "Epic Web3 RPG Game",
  "description": "An immersive blockchain-based RPG with NFT items",
  "cover_image_url": "https://example.com/image.jpg",
  "developer_name": "GameDev Studios",
  "genre": "RPG",
  "game_type": "web3",
  "created_at": "2025-10-15T10:00:00Z",
  "updated_at": "2025-10-15T10:00:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Project not found"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "Invalid project ID format"
}
```

---

### 4. Create Project

Create a new project entry.

**Endpoint:** `POST /projects`

**Request Body:**
```json
{
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "Epic Web3 RPG Game",
  "description": "An immersive blockchain-based RPG",
  "cover_image_url": "https://example.com/image.jpg",
  "developer_name": "GameDev Studios",
  "genre": "RPG",
  "game_type": "web3"
}
```

**Required Fields:**
- `creator_wallet_address` (string, VARCHAR(42))
- `title` (string, VARCHAR(255))

**Optional Fields:**
- `description` (string, TEXT)
- `cover_image_url` (string, VARCHAR(255))
- `developer_name` (string, VARCHAR(100))
- `genre` (string, VARCHAR(50))
- `game_type` (string, VARCHAR(10)) - e.g., "web2" or "web3"

**Response:** `201 Created`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "Epic Web3 RPG Game",
  "description": "An immersive blockchain-based RPG",
  "cover_image_url": "https://example.com/image.jpg",
  "developer_name": "GameDev Studios",
  "genre": "RPG",
  "game_type": "web3",
  "created_at": "2025-10-15T10:00:00Z",
  "updated_at": "2025-10-15T10:00:00Z"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "creator_wallet_address is required"
}
```

---

### 5. Update Project

Partially update project information.

**Endpoint:** `PATCH /projects/{id}`

**URL Parameters:**
- `id` (UUID, required) - Project ID

**Request Body:** (All fields optional)
```json
{
  "title": "Updated Game Title",
  "description": "Updated description",
  "cover_image_url": "https://example.com/new-image.jpg",
  "developer_name": "New Studio Name",
  "genre": "Action RPG",
  "game_type": "web3"
}
```

**Response:** `200 OK`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "Updated Game Title",
  "description": "Updated description",
  "cover_image_url": "https://example.com/new-image.jpg",
  "developer_name": "New Studio Name",
  "genre": "Action RPG",
  "game_type": "web3",
  "created_at": "2025-10-15T10:00:00Z",
  "updated_at": "2025-10-15T11:00:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Project not found"
}
```

---

## User Profiles

### 6. Get User Profile

Get user profile by wallet address.

**Endpoint:** `GET /profiles/{walletAddress}`

**URL Parameters:**
- `walletAddress` (string, required) - Ethereum wallet address (42 characters)

**Response:** `200 OK`
```json
{
  "wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "username": "epicgamer123",
  "email": "gamer@example.com",
  "profile_image_url": "https://example.com/avatars/gamer.jpg",
  "kyc_status": "verified",
  "created_at": "2025-10-15T10:00:00Z",
  "updated_at": "2025-10-15T10:00:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Profile not found"
}
```

---

### 7. Create or Update User Profile (Upsert)

Create a new profile or update existing one.

**Endpoint:** `PUT /profiles/{walletAddress}`

**URL Parameters:**
- `walletAddress` (string, required) - Ethereum wallet address

**Request Body:**
```json
{
  "username": "epicgamer123",
  "email": "gamer@example.com",
  "profile_image_url": "https://example.com/avatars/gamer.jpg",
  "kyc_status": "verified"
}
```

**Required Fields:**
- `username` (string, VARCHAR(50), unique)
- `email` (string, VARCHAR(255), unique)

**Optional Fields:**
- `profile_image_url` (string, VARCHAR(255))
- `kyc_status` (string, VARCHAR(20)) - Values: "unverified", "pending", "verified"

**Response:** `200 OK`
```json
{
  "wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "username": "epicgamer123",
  "email": "gamer@example.com",
  "profile_image_url": "https://example.com/avatars/gamer.jpg",
  "kyc_status": "verified",
  "created_at": "2025-10-15T10:00:00Z",
  "updated_at": "2025-10-15T10:00:00Z"
}
```

**Error Response:** `409 Conflict`
```json
{
  "error": "Username or email already exists"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "username is required"
}
```

---

## Comments

### 8. Get Project Comments

Get all comments for a specific project.

**Endpoint:** `GET /projects/{id}/comments`

**URL Parameters:**
- `id` (UUID, required) - Project ID

**Response:** `200 OK`
```json
[
  {
    "id": "660e8400-e29b-41d4-a716-446655440001",
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "parent_comment_id": null,
    "content": "This project looks amazing!",
    "created_at": "2025-10-15T10:30:00Z",
    "updated_at": "2025-10-15T10:30:00Z"
  },
  {
    "id": "770e8400-e29b-41d4-a716-446655440002",
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "parent_comment_id": "660e8400-e29b-41d4-a716-446655440001",
    "content": "Thank you for your support!",
    "created_at": "2025-10-15T10:35:00Z",
    "updated_at": "2025-10-15T10:35:00Z"
  }
]
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Project not found"
}
```

---

### 9. Create Comment

Add a new comment to a project.

**Endpoint:** `POST /projects/{id}/comments`

**URL Parameters:**
- `id` (UUID, required) - Project ID

**Request Body:**
```json
{
  "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "content": "This project looks amazing! Can't wait to play it.",
  "parent_comment_id": "660e8400-e29b-41d4-a716-446655440001"
}
```

**Required Fields:**
- `author_wallet_address` (string, VARCHAR(42))
- `content` (string, TEXT)

**Optional Fields:**
- `parent_comment_id` (UUID, nullable) - For reply comments

**Response:** `201 Created`
```json
{
  "id": "880e8400-e29b-41d4-a716-446655440003",
  "project_id": "550e8400-e29b-41d4-a716-446655440000",
  "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "parent_comment_id": "660e8400-e29b-41d4-a716-446655440001",
  "content": "This project looks amazing! Can't wait to play it.",
  "created_at": "2025-10-15T10:40:00Z",
  "updated_at": "2025-10-15T10:40:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Project not found"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "Parent comment not found"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "Parent comment does not belong to this project"
}
```

---

### 10. Get Project External Links

Get all external links (social media, website, etc.) for a specific project.

**Endpoint:** `GET /projects/{id}/links`

**URL Parameters:**
- `id` (UUID, required) - Project ID

**Response:** `200 OK`
```json
[
  {
    "id": "990e8400-e29b-41d4-a716-446655440004",
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Instagram",
    "url": "https://instagram.com/mygame",
    "created_at": "2025-10-15T10:50:00Z",
    "updated_at": "2025-10-15T10:50:00Z"
  },
  {
    "id": "aa0e8400-e29b-41d4-a716-446655440005",
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Twitter",
    "url": "https://twitter.com/mygame",
    "created_at": "2025-10-15T10:51:00Z",
    "updated_at": "2025-10-15T10:51:00Z"
  },
  {
    "id": "bb0e8400-e29b-41d4-a716-446655440006",
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Website",
    "url": "https://mygame.com",
    "created_at": "2025-10-15T10:52:00Z",
    "updated_at": "2025-10-15T10:52:00Z"
  }
]
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Project not found"
}
```

---

### 11. Create External Link

Add a new external link to a project (e.g., social media, website).

**Endpoint:** `POST /projects/{id}/links`

**URL Parameters:**
- `id` (UUID, required) - Project ID

**Request Body:**
```json
{
  "name": "Instagram",
  "url": "https://instagram.com/mygame"
}
```

**Required Fields:**
- `name` (string, VARCHAR(50)) - Link name/type (e.g., "Instagram", "Twitter", "Website")
- `url` (string, VARCHAR(500)) - The actual URL

**Response:** `201 Created`
```json
{
  "id": "990e8400-e29b-41d4-a716-446655440004",
  "project_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Instagram",
  "url": "https://instagram.com/mygame",
  "created_at": "2025-10-15T10:50:00Z",
  "updated_at": "2025-10-15T10:50:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "Project not found"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "Name and URL are required"
}
```

---

### 12. Update External Link

Update an existing external link for a project.

**Endpoint:** `PUT /projects/{id}/links/{linkId}`

**URL Parameters:**
- `id` (UUID, required) - Project ID
- `linkId` (UUID, required) - External Link ID

**Request Body:**
```json
{
  "name": "Instagram Official",
  "url": "https://instagram.com/mygame_official"
}
```

**Required Fields:**
- `name` (string, VARCHAR(50))
- `url` (string, VARCHAR(500))

**Response:** `200 OK`
```json
{
  "id": "990e8400-e29b-41d4-a716-446655440004",
  "project_id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Instagram Official",
  "url": "https://instagram.com/mygame_official",
  "created_at": "2025-10-15T10:50:00Z",
  "updated_at": "2025-10-15T11:00:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "External link not found"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "Name and URL are required"
}
```

---

### 13. Delete External Link

Delete a specific external link from a project.

**Endpoint:** `DELETE /projects/{id}/links/{linkId}`

**URL Parameters:**
- `id` (UUID, required) - Project ID
- `linkId` (UUID, required) - External Link ID

**Response:** `200 OK`
```json
{
  "message": "External link deleted successfully"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "External link not found"
}
```

---

## Data Models

### Project Model
```go
type Project struct {
    ID                   UUID      // Primary Key
    CreatorWalletAddress string    // VARCHAR(42)
    Title                string    // VARCHAR(255)
    Description          string    // TEXT
    CoverImageURL        string    // VARCHAR(255)
    DeveloperName        string    // VARCHAR(100)
    Genre                string    // VARCHAR(50)
    GameType             string    // VARCHAR(10)
    CreatedAt            timestamp // Auto-managed
    UpdatedAt            timestamp // Auto-managed
}
```

### UserProfile Model
```go
type UserProfile struct {
    WalletAddress   string    // VARCHAR(42), Primary Key
    Username        string    // VARCHAR(50), Unique
    Email           string    // VARCHAR(255), Unique
    ProfileImageURL string    // VARCHAR(255)
    KYCStatus       string    // VARCHAR(20), Default: 'unverified'
    CreatedAt       timestamp // Auto-managed
    UpdatedAt       timestamp // Auto-managed
}
```

### Comment Model
```go
type Comment struct {
    ID                  UUID      // Primary Key
    ProjectID           UUID      // Foreign Key -> projects.id
    AuthorWalletAddress string    // VARCHAR(42)
    ParentCommentID     UUID      // Foreign Key -> comments.id (nullable)
    Content             string    // TEXT
    CreatedAt           timestamp // Auto-managed
    UpdatedAt           timestamp // Auto-managed
}
```

### ExternalLink Model
```go
type ExternalLink struct {
    ID        UUID      // Primary Key
    ProjectID UUID      // Foreign Key -> projects.id
    Name      string    // VARCHAR(50) - e.g., "Instagram", "Twitter", "Website"
    URL       string    // VARCHAR(500) - The actual link
    CreatedAt timestamp // Auto-managed
    UpdatedAt timestamp // Auto-managed
}
```

    UpdatedAt           timestamp // Auto-managed
}
```

---

## Notes

1. **UUIDs**: All IDs are **UUID v7** format (time-ordered, sortable, more efficient for database indexing)
2. **Timestamps**: All timestamps are in RFC3339 format with timezone
3. **Wallet Addresses**: Ethereum addresses are 42 characters (0x + 40 hex chars)
4. **No Authentication**: This API doesn't require authentication as critical state is managed on-chain
5. **CORS**: Enabled for all origins in development mode
6. **Validation**: Basic validation is performed on required fields
7. **Nested Comments**: Use `parent_comment_id` to create threaded discussions

---

## Example Usage

### Creating a Complete Project Workflow

```bash
# 1. Create a project
PROJECT_ID=$(curl -s -X POST http://localhost:3000/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "title": "My Game",
    "description": "A great game"
  }' | jq -r '.id')

# 2. Create user profile
curl -X PUT http://localhost:3000/api/v1/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb \
  -H "Content-Type: application/json" \
  -d '{
    "username": "gamer1",
    "email": "gamer@example.com"
  }'

# 3. Add external links to the project
curl -X POST http://localhost:3000/api/v1/projects/$PROJECT_ID/links \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Instagram",
    "url": "https://instagram.com/mygame"
  }'

curl -X POST http://localhost:3000/api/v1/projects/$PROJECT_ID/links \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Twitter",
    "url": "https://twitter.com/mygame"
  }'

# 4. Add a comment
curl -X POST http://localhost:3000/api/v1/projects/$PROJECT_ID/comments \
  -H "Content-Type: application/json" \
  -d '{
    "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "content": "Great project!"
  }'

# 5. Get all project data (includes external links)
curl http://localhost:3000/api/v1/projects/$PROJECT_ID

# 6. Get all external links for a project
curl http://localhost:3000/api/v1/projects/$PROJECT_ID/links
```
