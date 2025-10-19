package repository

import (
	"errors"

	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserProfileRepository menangani operasi database untuk user profiles
type UserProfileRepository struct {
	db *gorm.DB
}

// NewUserProfileRepository membuat instance baru dari UserProfileRepository
func NewUserProfileRepository(db *gorm.DB) *UserProfileRepository {
	return &UserProfileRepository{db: db}
}

// GetByWalletAddress mengambil profil berdasarkan wallet address
func (r *UserProfileRepository) GetByWalletAddress(walletAddress string) (*model.UserProfile, error) {
	var profile model.UserProfile
	result := r.db.First(&profile, "wallet_address = ?", walletAddress)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &profile, result.Error
}

// Create membuat profil baru
func (r *UserProfileRepository) Create(profile *model.UserProfile) error {
	return r.db.Create(profile).Error
}

// Update memperbarui profil
func (r *UserProfileRepository) Update(profile *model.UserProfile) error {
	return r.db.Save(profile).Error
}

// Upsert membuat atau memperbarui profil (create or update)
func (r *UserProfileRepository) Upsert(profile *model.UserProfile) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "wallet_address"}},
		DoUpdates: clause.AssignmentColumns([]string{"username", "email", "profile_image_url", "kyc_status", "updated_at"}),
	}).Create(profile).Error
}

// Delete menghapus profil
func (r *UserProfileRepository) Delete(walletAddress string) error {
	return r.db.Delete(&model.UserProfile{}, "wallet_address = ?", walletAddress).Error
}
