package repository

import (
	"errors"

	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"gorm.io/gorm"
)

// ExternalLinkRepository menangani operasi database untuk external links
type ExternalLinkRepository struct {
	db *gorm.DB
}

// NewExternalLinkRepository membuat instance baru dari ExternalLinkRepository
func NewExternalLinkRepository(db *gorm.DB) *ExternalLinkRepository {
	return &ExternalLinkRepository{db: db}
}

// GetByProjectID mengambil semua external links untuk sebuah proyek
func (r *ExternalLinkRepository) GetByProjectID(projectID uint64) ([]model.ExternalLink, error) {
	var links []model.ExternalLink
	result := r.db.Where("project_id = ?", projectID).Order("created_at ASC").Find(&links)
	return links, result.Error
}

// GetByID mengambil external link berdasarkan ID
func (r *ExternalLinkRepository) GetByID(id uint64) (*model.ExternalLink, error) {
	var link model.ExternalLink
	result := r.db.First(&link, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &link, result.Error
}

// Create membuat external link baru
func (r *ExternalLinkRepository) Create(link *model.ExternalLink) error {
	return r.db.Create(link).Error
}

// Update memperbarui external link
func (r *ExternalLinkRepository) Update(link *model.ExternalLink) error {
	return r.db.Save(link).Error
}

// Delete menghapus external link
func (r *ExternalLinkRepository) Delete(id uint64) error {
	return r.db.Delete(&model.ExternalLink{}, "id = ?", id).Error
}

// DeleteByProjectID menghapus semua external links untuk sebuah proyek
func (r *ExternalLinkRepository) DeleteByProjectID(projectID uint64) error {
	return r.db.Where("project_id = ?", projectID).Delete(&model.ExternalLink{}).Error
}
