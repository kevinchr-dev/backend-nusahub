# Contributing Guide

Terima kasih atas minat Anda untuk berkontribusi pada Web3 Crowdfunding API!

## Development Setup

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Git
- Docker (optional)

### Setup Local Environment

1. **Fork dan Clone Repository**
```bash
git clone https://github.com/yourusername/hackathon.git
cd hackathon
```

2. **Install Dependencies**
```bash
go mod download
```

3. **Setup Database**
```bash
# Menggunakan Docker
make docker-db

# Atau manual PostgreSQL
createdb web3_crowdfunding
```

4. **Configure Environment**
```bash
cp .env.example .env
# Edit .env sesuai kebutuhan
```

5. **Run Application**
```bash
make run
```

## Code Style

### Go Formatting
Gunakan `gofmt` untuk formatting:
```bash
go fmt ./...
```

### Naming Conventions
- **Packages**: lowercase, single word
- **Files**: snake_case
- **Types**: PascalCase
- **Functions/Methods**: PascalCase (exported), camelCase (unexported)
- **Variables**: camelCase
- **Constants**: PascalCase atau UPPER_CASE

### Example
```go
// Good
type UserProfile struct {
    WalletAddress string
}

func (r *UserProfileRepository) GetByWalletAddress(addr string) (*UserProfile, error) {
    // implementation
}

// Bad
type user_profile struct {
    wallet_address string
}

func (r *UserProfileRepository) get_by_wallet_address(addr string) (*UserProfile, error) {
    // implementation
}
```

## Project Structure

Ikuti struktur yang sudah ada:

```
internal/
├── config/      # Configuration management
├── database/    # Database initialization
├── handler/     # HTTP handlers (controllers)
├── model/       # GORM models
├── repository/  # Database operations
└── router/      # Route definitions
```

### Adding New Feature

Untuk menambahkan feature baru (misalnya: Investments):

1. **Create Model** (`internal/model/model.go`)
```go
type Investment struct {
    ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    ProjectID     uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
    InvestorAddr  string    `gorm:"type:varchar(42);not null" json:"investor_address"`
    // ... other fields
}
```

2. **Create Repository** (`internal/repository/investment_repository.go`)
```go
type InvestmentRepository struct {
    db *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) *InvestmentRepository {
    return &InvestmentRepository{db: db}
}

// Add CRUD methods
```

3. **Create Handler** (`internal/handler/investment_handler.go`)
```go
type InvestmentHandler struct {
    repo *repository.InvestmentRepository
}

func NewInvestmentHandler(repo *repository.InvestmentRepository) *InvestmentHandler {
    return &InvestmentHandler{repo: repo}
}

// Add HTTP handlers
```

4. **Add Routes** (`internal/router/router.go`)
```go
investments := api.Group("/investments")
investments.Get("/", investmentHandler.GetAll)
investments.Post("/", investmentHandler.Create)
```

5. **Update Main** (`cmd/main/main.go`)
```go
investmentRepo := repository.NewInvestmentRepository(db)
investmentHandler := handler.NewInvestmentHandler(investmentRepo)
// Pass to router
```

## Commit Guidelines

### Commit Message Format
```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Maintenance tasks

### Examples
```bash
feat(projects): add filtering by genre

- Add query parameter for genre filtering
- Update repository to support filtering
- Add tests for new functionality

Closes #123
```

```bash
fix(comments): resolve nested comment validation

- Fix parent comment validation logic
- Add null check for parent_comment_id
- Update error messages

Fixes #456
```

## Testing

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/repository/...
```

### Writing Tests

Create test files alongside your code:
```
internal/
  repository/
    project_repository.go
    project_repository_test.go
```

Example test:
```go
func TestProjectRepository_Create(t *testing.T) {
    // Setup test database
    db := setupTestDB()
    repo := NewProjectRepository(db)
    
    // Create test data
    project := &model.Project{
        Title: "Test Project",
        CreatorWalletAddress: "0x123...",
    }
    
    // Test
    err := repo.Create(project)
    
    // Assert
    if err != nil {
        t.Errorf("Create failed: %v", err)
    }
    
    if project.ID == uuid.Nil {
        t.Error("ID should not be nil")
    }
}
```

## Pull Request Process

1. **Create a Branch**
```bash
git checkout -b feature/your-feature-name
```

2. **Make Changes**
- Write clean, documented code
- Follow style guidelines
- Add tests if applicable
- Update documentation

3. **Test Your Changes**
```bash
go test ./...
go fmt ./...
go vet ./...
```

4. **Commit**
```bash
git add .
git commit -m "feat(scope): description"
```

5. **Push**
```bash
git push origin feature/your-feature-name
```

6. **Create Pull Request**
- Provide clear description
- Reference related issues
- Wait for review

### PR Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] Tests added/updated
- [ ] All tests pass
- [ ] No merge conflicts

## Code Review

### As a Reviewer
- Be respectful and constructive
- Focus on code quality and maintainability
- Suggest improvements, don't demand
- Approve when ready

### As an Author
- Be open to feedback
- Explain your decisions
- Make requested changes
- Thank reviewers

## Documentation

Update documentation when:
- Adding new endpoints
- Changing API behavior
- Adding new features
- Modifying configuration

Files to update:
- `README.md` - Main documentation
- `API_DOCUMENTATION.md` - API specs
- Code comments - Inline documentation

## Common Issues

### Import Issues
```bash
go mod tidy
```

### Database Migrations
Auto-migration runs on startup. For manual migration:
```go
db.AutoMigrate(&model.NewModel{})
```

### Port Already in Use
```bash
lsof -i :3000
kill -9 <PID>
```

## Getting Help

- **Issues**: Open an issue on GitHub
- **Discussions**: Use GitHub Discussions
- **Email**: contact@example.com

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing! 🎉
