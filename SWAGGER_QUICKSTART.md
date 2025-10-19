# 🚀 Quick Start - Swagger Documentation

## Langkah Cepat untuk Melihat API Documentation

### 1. Jalankan Server
```bash
# Pilih salah satu cara:

# Cara 1: Langsung run
go run cmd/main/main.go

# Cara 2: Build dulu, lalu run
make build
./bin/api

# Cara 3: Docker Compose (dengan PostgreSQL)
docker-compose up -d
```

### 2. Buka Swagger UI
Setelah server berjalan, buka browser dan akses:
```
http://localhost:3000/docs/
```

### 3. Test API Langsung dari Browser
- ✅ Klik endpoint yang ingin di-test
- ✅ Klik tombol **"Try it out"**
- ✅ Isi parameter/body sesuai contoh
- ✅ Klik **"Execute"**
- ✅ Lihat response langsung!

## 📸 Screenshot Workflow

### Tampilan Swagger UI
```
┌─────────────────────────────────────────────────┐
│ Web3 Crowdfunding API v1.0                      │
│                                                  │
│ [ Authorize ]                                    │
│                                                  │
│ ▼ Projects                                       │
│   GET  /api/v1/projects                         │
│   GET  /api/v1/projects/{id}                    │
│   POST /api/v1/projects                         │
│   PATCH /api/v1/projects/{id}                   │
│                                                  │
│ ▼ User Profiles                                  │
│   GET  /api/v1/profiles/{walletAddress}         │
│   PUT  /api/v1/profiles/{walletAddress}         │
│                                                  │
│ ▼ Comments                                       │
│   GET  /api/v1/projects/{id}/comments           │
│   POST /api/v1/projects/{id}/comments           │
│                                                  │
│ ▼ Models                                         │
│   Project                                        │
│   UserProfile                                    │
│   Comment                                        │
└─────────────────────────────────────────────────┘
```

## 🎯 Contoh Testing

### Test 1: Create Project
1. Scroll ke **POST /api/v1/projects**
2. Klik **"Try it out"**
3. Copy-paste JSON ini:
```json
{
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "Amazing Web3 Game",
  "description": "A revolutionary blockchain game",
  "developer_name": "Your Studio",
  "genre": "RPG",
  "game_type": "web3"
}
```
4. Klik **"Execute"**
5. Lihat response dengan UUID v7 yang ter-generate!

### Test 2: Get All Projects
1. Scroll ke **GET /api/v1/projects**
2. Klik **"Try it out"**
3. Klik **"Execute"**
4. Lihat list semua projects dalam format JSON

### Test 3: Create User Profile
1. Scroll ke **PUT /api/v1/profiles/{walletAddress}**
2. Klik **"Try it out"**
3. Isi `walletAddress` dengan: `0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb`
4. Copy-paste JSON ini:
```json
{
  "username": "gamer123",
  "email": "gamer@example.com",
  "kyc_status": "verified"
}
```
5. Klik **"Execute"**

## 📤 Share dengan Partner

### Export Swagger JSON
File Swagger tersedia di:
```
docs/swagger.json
```

Partner bisa:
- ✅ Import ke Postman
- ✅ Import ke Insomnia
- ✅ View di Swagger Editor
- ✅ Generate client code

### Kirim Link
Setelah deploy ke production:
```
https://your-domain.com/docs/
```

## 🔄 Update Documentation

Setelah menambah/mengubah endpoint:
```bash
# Regenerate Swagger docs
make swagger

# Atau manual
swag init -g cmd/main/main.go --output docs

# Restart server
make run
```

## 📱 Mobile Friendly

Swagger UI responsive dan bisa diakses dari:
- ✅ Desktop browser
- ✅ Mobile browser
- ✅ Tablet

## 🎓 Learn More

Dokumentasi lengkap tersedia di:
- `SWAGGER_GUIDE.md` - Complete Swagger guide
- `API_DOCUMENTATION.md` - Detailed API specs
- `README.md` - Project overview

---

**Selamat mencoba! 🎉**

Jika ada pertanyaan, buka `SWAGGER_GUIDE.md` untuk panduan lengkap.
