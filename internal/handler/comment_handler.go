package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"github.com/kevinchr/web3-crowdfunding-api/internal/repository"
)

// CommentHandler menangani HTTP requests untuk comments
type CommentHandler struct {
	repo        *repository.CommentRepository
	projectRepo *repository.ProjectRepository
}

// NewCommentHandler membuat instance baru dari CommentHandler
func NewCommentHandler(repo *repository.CommentRepository, projectRepo *repository.ProjectRepository) *CommentHandler {
	return &CommentHandler{
		repo:        repo,
		projectRepo: projectRepo,
	}
}

// GetCommentsByProjectID godoc
// @Summary      Get project comments
// @Description  Get all comments for a specific project
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Project ID (UUID v7)"
// @Success      200  {array}   model.Comment
// @Failure      400  {object}  model.ErrorResponse
// @Failure      404  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /projects/{id}/comments [get]
func (h *CommentHandler) GetCommentsByProjectID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	projectID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	// Cek apakah proyek ada
	project, err := h.projectRepo.GetByID(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to verify project",
		})
	}

	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	comments, err := h.repo.GetByProjectID(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch comments",
		})
	}

	return c.JSON(comments)
}

// CreateComment godoc
// @Summary      Create comment
// @Description  Add a new comment to a project. Supports nested comments via parent_comment_id
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param        id       path      string         true  "Project ID (UUID v7)"
// @Param        comment  body      model.CommentCreate  true  "Comment data"
// @Success      201      {object}  model.Comment
// @Failure      400      {object}  model.ErrorResponse
// @Failure      404      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /projects/{id}/comments [post]
func (h *CommentHandler) CreateComment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	projectID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	// Cek apakah proyek ada
	project, err := h.projectRepo.GetByID(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to verify project",
		})
	}

	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	var comment model.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Set project ID dari URL param
	comment.ProjectID = projectID

	// Validasi field yang wajib diisi
	if strings.TrimSpace(comment.AuthorWalletAddress) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "author_wallet_address is required",
		})
	}

	if strings.TrimSpace(comment.Content) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "content is required",
		})
	}

	// Jika ada parent comment, validasi bahwa parent comment ada
	if comment.ParentCommentID != nil {
		parentComment, err := h.repo.GetByID(*comment.ParentCommentID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to verify parent comment",
			})
		}

		if parentComment == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Parent comment not found",
			})
		}

		// Pastikan parent comment juga untuk proyek yang sama
		if parentComment.ProjectID != projectID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Parent comment does not belong to this project",
			})
		}
	}

	// Generate UUID v7 baru
	comment.ID = uuid.Must(uuid.NewV7())

	if err := h.repo.Create(&comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create comment",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}
