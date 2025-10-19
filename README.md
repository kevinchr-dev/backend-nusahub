# Web3 Crowdfunding API

REST API service untuk platform crowdfunding Web3 yang dibangun dengan Golang, Fiber, dan PostgreSQL.

## 🚀 Fitur

- ✅ CRUD operations untuk Projects
- ✅ **Investor Tracking** - Track investor wallet addresses dengan PostgreSQL arrays
- ✅ Manajemen User Profiles (Upsert)
- ✅ Sistem Comments dengan support untuk nested comments
- ✅ Manajemen External Links untuk setiap project (social media, website, dll)
- ✅ Auto-migration database dengan GORM
- ✅ UUID v7 untuk primary keys (time-ordered, better DB performance)
- ✅ CORS enabled untuk kemudahan pengembangan
- ✅ Request logging
- ✅ Error handling yang konsisten
- ✅ Validasi input
- ✅ Interactive Swagger/OpenAPI Documentation

## 📋 Prasyarat

- Go 1.21 atau lebih tinggi
- PostgreSQL 12 atau lebih tinggi

## 🛠️ Instalasi

1. Clone repository:
```bash
git clone <repository-url>
cd hackathon
```

2. Install dependencies:
```bash
go mod download
```

3. Setup database PostgreSQL dan buat database baru:
```sql
CREATE DATABASE web3_crowdfunding;
```

4. Copy file `.env` dan sesuaikan konfigurasi database:
```bash
# File .env sudah ada, sesuaikan dengan konfigurasi PostgreSQL Anda
```

5. Jalankan aplikasi:
```bash
go run cmd/main/main.go
```

Server akan berjalan di `http://localhost:3000`

## � API Documentation

API ini dilengkapi dengan **Swagger/OpenAPI documentation** yang interaktif!

**Akses Swagger UI:**
```
http://localhost:3000/docs/
```

**Fitur Swagger Documentation:**
- ✅ Interactive API testing
- ✅ Complete request/response examples
- ✅ Model schemas
- ✅ Try it out functionality
- ✅ Easy sharing dengan partner

**Generate ulang Swagger docs:**
```bash
make swagger
# atau
swag init -g cmd/main/main.go --output docs
```

## �📚 API Endpoints

### Projects

#### GET /api/v1/projects
Mendapatkan semua proyek (termasuk investor_wallet_addresses array).

**Response:**
```json
[
  {
    "id": "uuid",
    "creator_wallet_address": "0x...",
    "title": "My Game Project",
    "description": "Description here",
    "cover_image_url": "https://...",
    "developer_name": "Developer Name",
    "genre": "RPG",
    "game_type": "web3",
    "created_at": "2025-10-15T10:00:00Z",
    "updated_at": "2025-10-15T10:00:00Z"
  }
]
```

#### GET /api/v1/projects/:id
Mendapatkan detail proyek berdasarkan ID.

**Response:**
- `200 OK`: Detail proyek
- `404 Not Found`: Proyek tidak ditemukan

#### POST /api/v1/projects
Membuat proyek baru.

**Request Body:**
```json
{
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "My Awesome Game",
  "description": "This is an amazing Web3 game",
  "cover_image_url": "https://example.com/image.jpg",
  "developer_name": "John Doe",
  "genre": "Action",
  "game_type": "web3"
}
```

**Response:**
- `201 Created`: Proyek berhasil dibuat (dengan UUID yang dihasilkan)
- `400 Bad Request`: Data tidak valid

#### PATCH /api/v1/projects/:id
Memperbarui proyek.

**Request Body:**
```json
{
  "title": "Updated Title",
  "description": "Updated description"
}
```

**Response:**
- `200 OK`: Proyek berhasil diperbarui
- `404 Not Found`: Proyek tidak ditemukan

### Investors

#### GET /api/v1/projects/:id/investors
Mendapatkan semua investor wallet addresses untuk sebuah proyek.

**Response:**
```json
{
  "investors": [
    "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "0x1234567890abcdef1234567890abcdef12345678"
  ]
}
```

#### POST /api/v1/projects/:id/investors
Menambahkan investor wallet address ke proyek.

**Request Body:**
```json
{
  "wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
}
```

**Response:**
- `200 OK`: Investor berhasil ditambahkan
- `404 Not Found`: Proyek tidak ditemukan
- `409 Conflict`: Investor sudah ada
- `400 Bad Request`: Format wallet address tidak valid

#### DELETE /api/v1/projects/:id/investors/:walletAddress
Menghapus investor dari proyek.

**Response:**
- `200 OK`: Investor berhasil dihapus
- `404 Not Found`: Proyek tidak ditemukan

### User Profiles

#### GET /api/v1/profiles/:walletAddress
Mendapatkan profil berdasarkan wallet address.

**Response:**
- `200 OK`: Detail profil
- `404 Not Found`: Profil tidak ditemukan

#### PUT /api/v1/profiles/:walletAddress
Membuat atau memperbarui profil (Upsert).

**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "profile_image_url": "https://example.com/avatar.jpg",
  "kyc_status": "verified"
}
```

**Response:**
- `200 OK`: Profil berhasil dibuat/diperbarui
- `400 Bad Request`: Data tidak valid
- `409 Conflict`: Username atau email sudah digunakan

### Comments

#### GET /api/v1/projects/:id/comments
Mendapatkan semua komentar untuk sebuah proyek.

**Response:**
```json
[
  {
    "id": "uuid",
    "project_id": "uuid",
    "author_wallet_address": "0x...",
    "parent_comment_id": null,
    "content": "Great project!",
    "created_at": "2025-10-15T10:00:00Z",
    "updated_at": "2025-10-15T10:00:00Z"
  }
]
```

#### POST /api/v1/projects/:id/comments
Menambahkan komentar baru.

**Request Body:**
```json
{
  "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "content": "This looks amazing!",
  "parent_comment_id": "uuid-optional"
}
```

### External Links

#### GET /api/v1/projects/:id/links
Mendapatkan semua external links untuk sebuah proyek (social media, website, dll).

**Response:**
```json
[
  {
    "id": "uuid",
    "project_id": "uuid",
    "name": "Instagram",
    "url": "https://instagram.com/mygame",
    "created_at": "2025-10-15T10:00:00Z",
    "updated_at": "2025-10-15T10:00:00Z"
  },
  {
    "id": "uuid",
    "project_id": "uuid",
    "name": "Twitter",
    "url": "https://twitter.com/mygame",
    "created_at": "2025-10-15T10:00:00Z",
    "updated_at": "2025-10-15T10:00:00Z"
  }
]
```

#### POST /api/v1/projects/:id/links
Menambahkan external link baru ke proyek.

**Request Body:**
```json
{
  "name": "Instagram",
  "url": "https://instagram.com/mygame"
}
```

**Response:**
- `201 Created`: Link berhasil ditambahkan
- `404 Not Found`: Proyek tidak ditemukan
- `400 Bad Request`: Data tidak valid

#### PUT /api/v1/projects/:id/links/:linkId
Memperbarui external link.

**Request Body:**
```json
{
  "name": "Instagram Official",
  "url": "https://instagram.com/mygame_official"
}
```

**Response:**
- `200 OK`: Link berhasil diperbarui
- `404 Not Found`: Link tidak ditemukan
- `400 Bad Request`: Data tidak valid

#### DELETE /api/v1/projects/:id/links/:linkId
Menghapus external link.

**Response:**
- `200 OK`: Link berhasil dihapus
- `404 Not Found`: Link tidak ditemukan

### Health Check

#### GET /api/v1/health
Mengecek status API.

**Response:**
```json
{
  "status": "ok",
  "message": "Web3 Crowdfunding API is running"
}
```

## 🗂️ Struktur Proyek

```
/hackathon
├── cmd/
│   └── main/
│       └── main.go              # Entry point aplikasi
├── internal/
│   ├── config/
│   │   └── config.go            # Konfigurasi aplikasi
│   ├── database/
│   │   └── database.go          # Inisialisasi database
│   ├── handler/
│   │   ├── project_handler.go   # Project handlers
│   │   ├── user_profile_handler.go  # Profile handlers
│   │   ├── comment_handler.go   # Comment handlers
│   │   └── external_link_handler.go  # External link handlers
│   ├── model/
│   │   └── model.go             # GORM models (4 tables)
│   ├── repository/
│   │   ├── project_repository.go     # Project repository
│   │   ├── user_profile_repository.go  # Profile repository
│   │   ├── comment_repository.go     # Comment repository
│   │   └── external_link_repository.go  # External link repository
│   └── router/
│       └── router.go            # Route definitions
├── docs/
│   ├── docs.go                  # Generated Swagger docs
│   ├── swagger.json             # OpenAPI JSON spec
│   └── swagger.yaml             # OpenAPI YAML spec
├── .env                         # Environment variables
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## 🔧 Konfigurasi

Edit file `.env` untuk mengubah konfigurasi:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=web3_crowdfunding
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=3000
```

## 🗄️ Database Schema

### Table: projects
- `id` (UUID, Primary Key)
- `creator_wallet_address` (VARCHAR(42))
- `title` (VARCHAR(255))
- `description` (TEXT)
- `cover_image_url` (VARCHAR(255))
- `developer_name` (VARCHAR(100))
- `genre` (VARCHAR(50))
- `game_type` (VARCHAR(10))
- `investor_wallet_addresses` (TEXT[], Array of wallet addresses)
- `created_at` (TIMESTAMPTZ)
- `updated_at` (TIMESTAMPTZ)

### Table: user_profiles
- `wallet_address` (VARCHAR(42), Primary Key)
- `username` (VARCHAR(50), Unique)
- `email` (VARCHAR(255), Unique)
- `profile_image_url` (VARCHAR(255))
- `kyc_status` (VARCHAR(20), Default: 'unverified')
- `created_at` (TIMESTAMPTZ)
- `updated_at` (TIMESTAMPTZ)

### Table: comments
- `id` (UUID, Primary Key)
- `project_id` (UUID, Foreign Key)
- `author_wallet_address` (VARCHAR(42))
- `parent_comment_id` (UUID, Foreign Key, Nullable)
- `content` (TEXT)
- `created_at` (TIMESTAMPTZ)
- `updated_at` (TIMESTAMPTZ)

### Table: external_links
- `id` (UUID, Primary Key)
- `project_id` (UUID, Foreign Key)
- `name` (VARCHAR(50)) - e.g., "Instagram", "Twitter", "Website"
- `url` (VARCHAR(500)) - The actual link
- `created_at` (TIMESTAMPTZ)
- `updated_at` (TIMESTAMPTZ)

## 🧪 Testing dengan cURL

### Membuat Project Baru
```bash
curl -X POST http://localhost:3000/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "title": "Epic Web3 Game",
    "description": "An amazing blockchain game",
    "developer_name": "GameDev Studio",
    "genre": "RPG",
    "game_type": "web3"
  }'
```

### Mendapatkan Semua Projects
```bash
curl http://localhost:3000/api/v1/projects
```

### Membuat User Profile
```bash
curl -X PUT http://localhost:3000/api/v1/profiles/0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb \
  -H "Content-Type: application/json" \
  -d '{
    "username": "gamer123",
    "email": "gamer@example.com",
    "kyc_status": "verified"
  }'
```

### Menambahkan Comment
```bash
curl -X POST http://localhost:3000/api/v1/projects/{project-id}/comments \
  -H "Content-Type: application/json" \
  -d '{
    "author_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "content": "This project looks promising!"
  }'
```

### Menambahkan External Link
```bash
curl -X POST http://localhost:3000/api/v1/projects/{project-id}/links \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Instagram",
    "url": "https://instagram.com/mygame"
  }'
```

### Mendapatkan External Links untuk Project
```bash
curl http://localhost:3000/api/v1/projects/{project-id}/links
```

## 📝 Catatan Pengembangan

- API ini tidak memerlukan autentikasi karena state krusial ditangani oleh smart contract
- UUID digunakan untuk memastikan konsistensi antara database off-chain dan smart contract on-chain
- Middleware CORS dikonfigurasi untuk menerima request dari semua origin (untuk development)
- Semua timestamp menggunakan TIMESTAMPTZ untuk timezone awareness
- Auto-migration akan otomatis membuat tabel saat aplikasi pertama kali dijalankan

## 🚀 Deployment

Untuk production:
1. Update CORS configuration untuk membatasi origins
2. Set `DB_SSLMODE=require` untuk koneksi database yang aman
3. Gunakan environment variables untuk konfigurasi sensitif
4. Implementasikan rate limiting
5. Setup monitoring dan logging yang proper

## 📄 License

MIT License

## 👨‍💻 Author

Backend Engineering Team
