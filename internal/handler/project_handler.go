package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"github.com/kevinchr/web3-crowdfunding-api/internal/repository"
)

// ProjectHandler menangani HTTP requests untuk projects
type ProjectHandler struct {
	repo *repository.ProjectRepository
}

// NewProjectHandler membuat instance baru dari ProjectHandler
func NewProjectHandler(repo *repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{repo: repo}
}

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
	projects, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch projects",
		})
	}

	return c.JSON(projects)
}

// GetProjectByID godoc
// @Summary      Get project by ID
// @Description  Get detailed information about a specific project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Project ID (UUID v7)"
// @Success      200  {object}  model.Project
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /projects/{id} [get]
func (h *ProjectHandler) GetProjectByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	project, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch project",
		})
	}

	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	return c.JSON(project)
}

// CreateProject godoc
// @Summary      Create new project
// @Description  Create a new crowdfunding project entry. UUID v7 will be auto-generated
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body      model.Project  true  "Project data"
// @Success      201      {object}  model.Project
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /projects [post]
func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	var project model.Project

	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validasi field yang wajib diisi
	if strings.TrimSpace(project.CreatorWalletAddress) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "creator_wallet_address is required",
		})
	}

	if strings.TrimSpace(project.Title) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title is required",
		})
	}

	// Generate UUID v7 baru (akan di-handle oleh BeforeCreate hook juga)
	project.ID = uuid.Must(uuid.NewV7())

	if err := h.repo.Create(&project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create project",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}

// UpdateProject godoc
// @Summary      Update project
// @Description  Partially update project information
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        id       path      string             true  "Project ID (UUID v7)"
// @Param        project  body      map[string]interface{}  true  "Fields to update"
// @Success      200      {object}  model.Project
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /projects/{id} [patch]
func (h *ProjectHandler) UpdateProject(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Hapus field yang tidak boleh diupdate
	delete(updates, "id")
	delete(updates, "created_at")

	project, err := h.repo.UpdatePartial(id, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update project",
		})
	}

	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	return c.JSON(project)
}

// AddInvestor godoc
// @Summary      Add investor to project
// @Description  Add an investor's wallet address to a project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Project ID (UUID v7)"
// @Param        body body      map[string]string  true  "Investor wallet address"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      409  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /projects/{id}/investors [post]
func (h *ProjectHandler) AddInvestor(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	var body struct {
		WalletAddress string `json:"wallet_address"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if strings.TrimSpace(body.WalletAddress) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wallet_address is required",
		})
	}

	// Validate wallet address format (basic check for Ethereum address)
	if len(body.WalletAddress) != 42 || !strings.HasPrefix(body.WalletAddress, "0x") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid wallet address format",
		})
	}

	err = h.repo.AddInvestor(id, body.WalletAddress)
	if err != nil {
		if err.Error() == "project not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Project not found",
			})
		}
		if err.Error() == "investor already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Investor already exists in this project",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add investor",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Investor added successfully",
	})
}

// RemoveInvestor godoc
// @Summary      Remove investor from project
// @Description  Remove an investor's wallet address from a project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        id             path      string  true  "Project ID (UUID v7)"
// @Param        walletAddress  path      string  true  "Investor wallet address"
// @Success      200            {object}  map[string]string
// @Failure      400            {object}  map[string]string
// @Failure      404            {object}  map[string]string
// @Failure      500            {object}  map[string]string
// @Router       /projects/{id}/investors/{walletAddress} [delete]
func (h *ProjectHandler) RemoveInvestor(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	walletAddress := c.Params("walletAddress")
	if strings.TrimSpace(walletAddress) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wallet_address is required",
		})
	}

	err = h.repo.RemoveInvestor(id, walletAddress)
	if err != nil {
		if err.Error() == "project not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Project not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to remove investor",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Investor removed successfully",
	})
}

// GetInvestors godoc
// @Summary      Get project investors
// @Description  Get all investor wallet addresses for a project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Project ID (UUID v7)"
// @Success      200  {object}  map[string][]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /projects/{id}/investors [get]
func (h *ProjectHandler) GetInvestors(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID format",
		})
	}

	investors, err := h.repo.GetInvestors(id)
	if err != nil {
		if err.Error() == "project not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Project not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch investors",
		})
	}

	// Return empty array if no investors
	if investors == nil {
		investors = []string{}
	}

	return c.JSON(fiber.Map{
		"investors": investors,
	})
}
