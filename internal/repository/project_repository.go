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
	result := r.db.Preload("Links").Find(&projects)
	return projects, result.Error
}

// GetByID mengambil proyek berdasarkan ID
func (r *ProjectRepository) GetByID(id uint64) (*model.Project, error) {
	var project model.Project
	result := r.db.Preload("Links").First(&project, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &project, result.Error
}

// Create membuat proyek baru
func (r *ProjectRepository) Create(project *model.Project) error {
	// create project and its links in a transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(project).Error; err != nil {
			return err
		}
		// If links provided, ensure ProjectID is set and create them
		if len(project.Links) > 0 {
			for i := range project.Links {
				// ensure DB will assign the ID (avoid inserting explicit id causing duplicates)
				project.Links[i].ID = 0
				project.Links[i].ProjectID = project.ID
			}
			if err := tx.Create(&project.Links).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// Update memperbarui proyek
func (r *ProjectRepository) Update(project *model.Project) error {
	// replace project fields and replace links in a transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(project).Error; err != nil {
			return err
		}
		// Replace links: delete existing and insert new ones if provided
		if err := tx.Where("project_id = ?", project.ID).Delete(&model.ExternalLink{}).Error; err != nil {
			return err
		}
		if len(project.Links) > 0 {
			for i := range project.Links {
				// ensure DB will assign the ID for new links
				project.Links[i].ID = 0
				project.Links[i].ProjectID = project.ID
			}
			if err := tx.Create(&project.Links).Error; err != nil {
				return err
			}
		}
		return nil
	})
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
