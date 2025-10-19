package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kevinchr/web3-crowdfunding-api/internal/config"
	"github.com/kevinchr/web3-crowdfunding-api/internal/database"
	"github.com/kevinchr/web3-crowdfunding-api/internal/handler"
	"github.com/kevinchr/web3-crowdfunding-api/internal/repository"
	"github.com/kevinchr/web3-crowdfunding-api/internal/router"

	_ "github.com/kevinchr/web3-crowdfunding-api/docs" // Import generated docs
)

// @title           Web3 Crowdfunding API
// @version         1.0
// @description     REST API service untuk platform crowdfunding Web3. API ini berfungsi sebagai lapisan data off-chain yang berkomunikasi dengan frontend. State krusial dan transaksi finansial ditangani oleh smart contract.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@web3crowdfunding.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:3000
// @BasePath  /api/v1

// @schemes http https

// @tag.name Projects
// @tag.description Endpoints untuk mengelola proyek crowdfunding

// @tag.name User Profiles
// @tag.description Endpoints untuk mengelola profil pengguna

// @tag.name Comments
// @tag.description Endpoints untuk mengelola komentar proyek

// @tag.name Health
// @tag.description Health check endpoint

func main() {
	// Load konfigurasi
	cfg := config.LoadConfig()

	// Inisialisasi database
	if err := database.InitDatabase(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db := database.GetDB()

	// Inisialisasi repositories
	projectRepo := repository.NewProjectRepository(db)
	profileRepo := repository.NewUserProfileRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	linkRepo := repository.NewExternalLinkRepository(db)

	// Inisialisasi handlers
	projectHandler := handler.NewProjectHandler(projectRepo)
	profileHandler := handler.NewUserProfileHandler(profileRepo)
	commentHandler := handler.NewCommentHandler(commentRepo, projectRepo)
	linkHandler := handler.NewExternalLinkHandler(linkRepo, projectRepo)

	// Inisialisasi Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Web3 Crowdfunding API v1.0",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New()) // Recover from panics
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup routes
	router.SetupRoutes(app, projectHandler, profileHandler, commentHandler, linkHandler)

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Web3 Crowdfunding API",
			"version": "1.0",
			"endpoints": fiber.Map{
				"health":   "/api/v1/health",
				"projects": "/api/v1/projects",
				"profiles": "/api/v1/profiles",
			},
		})
	})

	// Start server
	port := cfg.ServerPort
	log.Printf("Server starting on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
