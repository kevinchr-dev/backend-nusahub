# Swagger Documentation Guide

## 🎯 Akses Swagger UI

Setelah server berjalan, buka browser dan akses:

```
http://localhost:3000/docs/
```

Atau untuk production:
```
https://your-domain.com/docs/
```

## 📋 Fitur Swagger UI

### 1. **Interactive API Testing**
- Klik endpoint yang ingin di-test
- Klik tombol "Try it out"
- Isi parameter/body sesuai kebutuhan
- Klik "Execute"
- Lihat response langsung

### 2. **Model Schemas**
- Scroll ke bawah untuk melihat "Schemas"
- Lihat struktur lengkap untuk:
  - Project
  - UserProfile
  - Comment

### 3. **Request/Response Examples**
- Setiap endpoint menampilkan contoh request body
- Menampilkan semua kemungkinan response codes
- Contoh response untuk success dan error cases

## 🔧 Generate Swagger Documentation

### Method 1: Menggunakan Makefile
```bash
make swagger
```

### Method 2: Manual Command
```bash
swag init -g cmd/main/main.go --output docs
```

### Method 3: Dengan Watch Mode (Auto-regenerate)
```bash
swag init -g cmd/main/main.go --output docs --watch
```

## 📝 Update Swagger Annotations

Saat menambahkan endpoint baru atau mengubah existing endpoints, tambahkan/update Swagger annotations:

### Contoh Annotation untuk Handler

```go
// GetAllProjects godoc
// @Summary      Get all projects
// @Description  Retrieve list of all crowdfunding projects
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Project
// @Failure      500  {object}  map[string]string
// @Router       /projects [get]
func (h *ProjectHandler) GetAllProjects(c *fiber.Ctx) error {
    // implementation
}
```

### Annotation Tags Explanation

| Tag | Description | Example |
|-----|-------------|---------|
| `@Summary` | Short description | `Get all projects` |
| `@Description` | Detailed description | `Retrieve list of all crowdfunding projects` |
| `@Tags` | Group endpoints | `Projects` |
| `@Accept` | Request content type | `json` |
| `@Produce` | Response content type | `json` |
| `@Param` | Parameter definition | `id path string true "Project ID"` |
| `@Success` | Success response | `200 {object} model.Project` |
| `@Failure` | Error response | `404 {object} map[string]string` |
| `@Router` | Route path and method | `/projects [get]` |

### Parameter Types

```go
// Path parameter
// @Param id path string true "Project ID"

// Query parameter
// @Param page query int false "Page number"

// Body parameter
// @Param project body model.Project true "Project data"

// Header parameter
// @Param Authorization header string false "Bearer token"
```

## 🎨 Customize Swagger Configuration

Edit file `cmd/main/main.go` untuk mengubah konfigurasi Swagger:

```go
// @title           Web3 Crowdfunding API
// @version         1.0
// @description     Your API description

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:3000
// @BasePath  /api/v1
```

## 📤 Export Swagger Documentation

### JSON Format
File tersedia di: `docs/swagger.json`

### YAML Format
File tersedia di: `docs/swagger.yaml`

### Sharing dengan Team
Kirim file `swagger.json` atau `swagger.yaml` ke:
- Frontend developers
- QA team
- Partner integration
- API consumers

Mereka bisa import ke:
- Postman (import swagger.json)
- Insomnia
- Swagger Editor (https://editor.swagger.io/)
- API documentation tools

## 🚀 Best Practices

### 1. Selalu Update Docs
Setiap kali mengubah endpoint, jangan lupa:
```bash
make swagger
```

### 2. Use Descriptive Tags
Grouping yang baik memudahkan navigasi:
```go
// @Tags Projects
// @Tags User Profiles
// @Tags Comments
```

### 3. Provide Examples
Tambahkan contoh di description:
```go
// @Description  Create project. Example: {"title": "My Game", "creator_wallet_address": "0x123..."}
```

### 4. Document All Status Codes
```go
// @Success 200 {object} model.Project
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
```

### 5. Keep Models Updated
Pastikan model di `internal/model/model.go` memiliki json tags yang jelas:
```go
type Project struct {
  ID    uint64    `json:"id" example:"1420070400000000"`
    Title string    `json:"title" example:"My Awesome Game"`
}
```

## 🔍 Troubleshooting

### Swagger UI tidak muncul
1. Pastikan server running
2. Cek endpoint `/docs/` atau `/docs/index.html`
3. Cek browser console untuk errors

### Documentation tidak update
```bash
# Clean dan regenerate
make clean
make swagger
make build
```

### Import errors
```bash
# Pastikan docs package di-import di main.go
import _ "github.com/kevinchr/web3-crowdfunding-api/docs"
```

### Go module issues
```bash
go mod tidy
go mod download
```

## 📱 Testing dengan Swagger UI

### Test Create Project
1. Buka `/docs/`
2. Cari section "Projects"
3. Klik `POST /api/v1/projects`
4. Klik "Try it out"
5. Edit request body:
```json
{
  "creator_wallet_address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "title": "Test from Swagger",
  "description": "Testing API"
}
```
6. Klik "Execute"
7. Lihat response dengan UUID yang di-generate

### Test Get All Projects
1. Klik `GET /api/v1/projects`
2. Klik "Try it out"
3. Klik "Execute"
4. Lihat list semua projects

## 🌐 Production Deployment

Untuk production, update `@host` di main.go:

```go
// Development
// @host localhost:3000

// Production
// @host api.yourdomain.com
```

Kemudian regenerate:
```bash
make swagger
```

## 📚 Additional Resources

- [Swagger/OpenAPI Specification](https://swagger.io/specification/)
- [swag GitHub](https://github.com/swaggo/swag)
- [Fiber Swagger](https://github.com/gofiber/swagger)

---

**Happy Testing! 🎉**
