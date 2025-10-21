package repository

import (
	"errors"

	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"gorm.io/gorm"
)

// ProjectRepository menangani operasi database untuk projects
type ProjectRepository struct {
	db *gorm.DB
}

// NewProjectRepository membuat instance baru dari ProjectRepository
func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// GetAll mengambil semua proyek
func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project
	result := r.db.Find(&projects)
	return projects, result.Error
}

// GetByID mengambil proyek berdasarkan ID
func (r *ProjectRepository) GetByID(id uint64) (*model.Project, error) {
	var project model.Project
	result := r.db.First(&project, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &project, result.Error
}

// Create membuat proyek baru
func (r *ProjectRepository) Create(project *model.Project) error {
	return r.db.Create(project).Error
}

// Update memperbarui proyek
func (r *ProjectRepository) Update(project *model.Project) error {
	return r.db.Save(project).Error
}

// UpdatePartial memperbarui sebagian field proyek
func (r *ProjectRepository) UpdatePartial(id uint64, updates map[string]interface{}) (*model.Project, error) {
	var project model.Project
	result := r.db.First(&project, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	if err := r.db.Model(&project).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

// Delete menghapus proyek
func (r *ProjectRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Project{}, "id = ?", id).Error
}

// AddInvestor menambahkan wallet address investor ke project
func (r *ProjectRepository) AddInvestor(projectID uint64, walletAddress string) error {
	var project model.Project
	result := r.db.First(&project, "id = ?", projectID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("project not found")
	}
	if result.Error != nil {
		return result.Error
	}

	// Check if investor already exists
	for _, addr := range project.InvestorWalletAddresses {
		if addr == walletAddress {
			return errors.New("investor already exists")
		}
	}

	// Add investor using PostgreSQL array_append
	return r.db.Model(&project).Update("investor_wallet_addresses",
		gorm.Expr("array_append(investor_wallet_addresses, ?)", walletAddress)).Error
}

// RemoveInvestor menghapus wallet address investor dari project
func (r *ProjectRepository) RemoveInvestor(projectID uint64, walletAddress string) error {
	var project model.Project
	result := r.db.First(&project, "id = ?", projectID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("project not found")
	}
	if result.Error != nil {
		return result.Error
	}

	// Remove investor using PostgreSQL array_remove
	return r.db.Model(&project).Update("investor_wallet_addresses",
		gorm.Expr("array_remove(investor_wallet_addresses, ?)", walletAddress)).Error
}

// GetInvestors mengambil semua investor wallet addresses untuk project
func (r *ProjectRepository) GetInvestors(projectID uint64) ([]string, error) {
	var project model.Project
	result := r.db.First(&project, "id = ?", projectID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("project not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return project.InvestorWalletAddresses, nil
}
