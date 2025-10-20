package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"github.com/kevinchr/web3-crowdfunding-api/internal/repository"
)

// ExternalLinkHandler menangani HTTP requests untuk external links
type ExternalLinkHandler struct {
	repo        *repository.ExternalLinkRepository
	projectRepo *repository.ProjectRepository
}

// NewExternalLinkHandler membuat instance baru dari ExternalLinkHandler
func NewExternalLinkHandler(repo *repository.ExternalLinkRepository, projectRepo *repository.ProjectRepository) *ExternalLinkHandler {
	return &ExternalLinkHandler{
		repo:        repo,
		projectRepo: projectRepo,
	}
}

// GetLinksByProjectID godoc
// @Summary      Get project external links
// @Description  Get all external links for a specific project
// @Tags         External Links
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Project ID (UUID v7)"
// @Success      200  {array}   model.ExternalLink
// @Failure      400  {object}  model.ErrorResponse
// @Failure      404  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /projects/{id}/links [get]
func (h *ExternalLinkHandler) GetLinksByProjectID(c *fiber.Ctx) error {
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

	links, err := h.repo.GetByProjectID(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch external links",
		})
	}

	return c.JSON(links)
}

// CreateLink godoc
// @Summary      Create external link
// @Description  Add a new external link to a project (e.g., Instagram, Twitter, Website)
// @Tags         External Links
// @Accept       json
// @Produce      json
// @Param        id    path      string              true  "Project ID (UUID v7)"
// @Param        link  body      model.ExternalLinkCreate  true  "External link data"
// @Success      201   {object}  model.ExternalLink
// @Failure      400   {object}  model.ErrorResponse
// @Failure      404   {object}  model.ErrorResponse
// @Failure      500   {object}  model.ErrorResponse
// @Router       /projects/{id}/links [post]
func (h *ExternalLinkHandler) CreateLink(c *fiber.Ctx) error {
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

	var link model.ExternalLink
	if err := c.BodyParser(&link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Set project ID dari URL param
	link.ProjectID = projectID

	// Validasi field yang wajib diisi
	if strings.TrimSpace(link.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name is required",
		})
	}

	if strings.TrimSpace(link.URL) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "url is required",
		})
	}

	// Generate UUID v7 baru
	link.ID = uuid.Must(uuid.NewV7())

	if err := h.repo.Create(&link); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create external link",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(link)
}

// UpdateLink godoc
// @Summary      Update external link
// @Description  Update an external link
// @Tags         External Links
// @Accept       json
// @Produce      json
// @Param        id      path      string              true  "Project ID (UUID v7)"
// @Param        linkId  path      string              true  "Link ID (UUID v7)"
// @Param        link    body      model.ExternalLinkCreate  true  "External link data"
// @Success      200     {object}  model.ExternalLink
// @Failure      400     {object}  model.ErrorResponse
// @Failure      404     {object}  model.ErrorResponse
// @Failure      500     {object}  model.ErrorResponse
// @Router       /projects/{id}/links/{linkId} [put]
func (h *ExternalLinkHandler) UpdateLink(c *fiber.Ctx) error {
	linkIDParam := c.Params("linkId")
	linkID, err := uuid.Parse(linkIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link ID format",
		})
	}

	existingLink, err := h.repo.GetByID(linkID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch link",
		})
	}

	if existingLink == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "External link not found",
		})
	}

	var updates model.ExternalLink
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update only allowed fields
	if strings.TrimSpace(updates.Name) != "" {
		existingLink.Name = updates.Name
	}
	if strings.TrimSpace(updates.URL) != "" {
		existingLink.URL = updates.URL
	}

	if err := h.repo.Update(existingLink); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update external link",
		})
	}

	return c.JSON(existingLink)
}

// DeleteLink godoc
// @Summary      Delete external link
// @Description  Delete an external link from a project
// @Tags         External Links
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Project ID (UUID v7)"
// @Param        linkId  path      string  true  "Link ID (UUID v7)"
// @Success      200     {object}  model.GenericMessage
// @Failure      400     {object}  model.ErrorResponse
// @Failure      404     {object}  model.ErrorResponse
// @Failure      500     {object}  model.ErrorResponse
// @Router       /projects/{id}/links/{linkId} [delete]
func (h *ExternalLinkHandler) DeleteLink(c *fiber.Ctx) error {
	linkIDParam := c.Params("linkId")
	linkID, err := uuid.Parse(linkIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid link ID format",
		})
	}

	existingLink, err := h.repo.GetByID(linkID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch link",
		})
	}

	if existingLink == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "External link not found",
		})
	}

	if err := h.repo.Delete(linkID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete external link",
		})
	}

	return c.JSON(fiber.Map{
		"message": "External link deleted successfully",
	})
}
