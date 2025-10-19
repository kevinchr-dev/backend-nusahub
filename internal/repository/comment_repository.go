package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"gorm.io/gorm"
)

// CommentRepository menangani operasi database untuk comments
type CommentRepository struct {
	db *gorm.DB
}

// NewCommentRepository membuat instance baru dari CommentRepository
func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// GetByProjectID mengambil semua komentar untuk sebuah proyek
func (r *CommentRepository) GetByProjectID(projectID uuid.UUID) ([]model.Comment, error) {
	var comments []model.Comment
	result := r.db.Where("project_id = ?", projectID).Order("created_at DESC").Find(&comments)
	return comments, result.Error
}

// GetByID mengambil komentar berdasarkan ID
func (r *CommentRepository) GetByID(id uuid.UUID) (*model.Comment, error) {
	var comment model.Comment
	result := r.db.First(&comment, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &comment, result.Error
}

// Create membuat komentar baru
func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

// Update memperbarui komentar
func (r *CommentRepository) Update(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

// Delete menghapus komentar
func (r *CommentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Comment{}, "id = ?", id).Error
}

// GetReplies mengambil semua balasan untuk sebuah komentar
func (r *CommentRepository) GetReplies(parentID uuid.UUID) ([]model.Comment, error) {
	var comments []model.Comment
	result := r.db.Where("parent_comment_id = ?", parentID).Order("created_at ASC").Find(&comments)
	return comments, result.Error
}
