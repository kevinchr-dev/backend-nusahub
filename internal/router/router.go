package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kevinchr/web3-crowdfunding-api/internal/handler"
)

// SetupRoutes mengatur semua rute API
func SetupRoutes(app *fiber.App, projectHandler *handler.ProjectHandler, profileHandler *handler.UserProfileHandler, commentHandler *handler.CommentHandler, linkHandler *handler.ExternalLinkHandler) {
	// Swagger documentation endpoint
	app.Get("/docs/*", swagger.HandlerDefault)

	// API v1 group
	api := app.Group("/api/v1")

	// Routes untuk Projects
	projects := api.Group("/projects")
	projects.Get("/", projectHandler.GetAllProjects)
	projects.Get("/:id", projectHandler.GetProjectByID)
	projects.Post("/", projectHandler.CreateProject)
	projects.Patch("/:id", projectHandler.UpdateProject)

	// Routes untuk Investors (nested under projects)
	projects.Get("/:id/investors", projectHandler.GetInvestors)
	projects.Post("/:id/investors", projectHandler.AddInvestor)
	projects.Delete("/:id/investors/:walletAddress", projectHandler.RemoveInvestor)

	// Routes untuk Comments (nested under projects)
	projects.Get("/:id/comments", commentHandler.GetCommentsByProjectID)
	projects.Post("/:id/comments", commentHandler.CreateComment)

	// Routes untuk External Links (nested under projects)
	projects.Get("/:id/links", linkHandler.GetLinksByProjectID)
	projects.Post("/:id/links", linkHandler.CreateLink)
	projects.Put("/:id/links/:linkId", linkHandler.UpdateLink)
	projects.Delete("/:id/links/:linkId", linkHandler.DeleteLink)

	// Routes untuk User Profiles
	profiles := api.Group("/profiles")
	profiles.Get("/:walletAddress", profileHandler.GetProfileByWalletAddress)
	profiles.Put("/:walletAddress", profileHandler.UpsertProfile)

	// Health check endpoint
	// @Summary      Health check
	// @Description  Check API health status
	// @Tags         Health
	// @Accept       json
	// @Produce      json
	// @Success      200  {object}  map[string]string
	// @Router       /health [get]
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Web3 Crowdfunding API is running",
		})
	})
}
