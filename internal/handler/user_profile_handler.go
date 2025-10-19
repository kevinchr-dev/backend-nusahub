package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"github.com/kevinchr/web3-crowdfunding-api/internal/repository"
)

// UserProfileHandler menangani HTTP requests untuk user profiles
type UserProfileHandler struct {
	repo *repository.UserProfileRepository
}

// NewUserProfileHandler membuat instance baru dari UserProfileHandler
func NewUserProfileHandler(repo *repository.UserProfileRepository) *UserProfileHandler {
	return &UserProfileHandler{repo: repo}
}

// GetProfileByWalletAddress godoc
// @Summary      Get user profile
// @Description  Get user profile by wallet address
// @Tags         User Profiles
// @Accept       json
// @Produce      json
// @Param        walletAddress  path      string  true  "Ethereum Wallet Address (42 chars)"
// @Success      200            {object}  model.UserProfile
// @Failure      400            {object}  map[string]string
// @Failure      404            {object}  map[string]string
// @Failure      500            {object}  map[string]string
// @Router       /profiles/{walletAddress} [get]
func (h *UserProfileHandler) GetProfileByWalletAddress(c *fiber.Ctx) error {
	walletAddress := c.Params("walletAddress")

	if strings.TrimSpace(walletAddress) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wallet address is required",
		})
	}

	profile, err := h.repo.GetByWalletAddress(walletAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch profile",
		})
	}

	if profile == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Profile not found",
		})
	}

	return c.JSON(profile)
}

// UpsertProfile godoc
// @Summary      Create or update user profile
// @Description  Create a new profile or update existing one (Upsert operation)
// @Tags         User Profiles
// @Accept       json
// @Produce      json
// @Param        walletAddress  path      string             true  "Ethereum Wallet Address (42 chars)"
// @Param        profile        body      model.UserProfile  true  "User profile data"
// @Success      200            {object}  model.UserProfile
// @Failure      400            {object}  map[string]string
// @Failure      409            {object}  map[string]string
// @Failure      500            {object}  map[string]string
// @Router       /profiles/{walletAddress} [put]
func (h *UserProfileHandler) UpsertProfile(c *fiber.Ctx) error {
	walletAddress := c.Params("walletAddress")

	if strings.TrimSpace(walletAddress) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wallet address is required",
		})
	}

	var profile model.UserProfile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Set wallet address dari URL param
	profile.WalletAddress = walletAddress

	// Validasi field yang wajib diisi
	if strings.TrimSpace(profile.Username) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username is required",
		})
	}

	if strings.TrimSpace(profile.Email) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email is required",
		})
	}

	// Set default KYC status jika kosong
	if strings.TrimSpace(profile.KYCStatus) == "" {
		profile.KYCStatus = "unverified"
	}

	if err := h.repo.Upsert(&profile); err != nil {
		// Check jika error karena duplicate username atau email
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Username or email already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create or update profile",
		})
	}

	return c.JSON(profile)
}
