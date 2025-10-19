package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Project merepresentasikan tabel projects
type Project struct {
	ID                      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatorWalletAddress    string    `gorm:"type:varchar(42);not null" json:"creator_wallet_address"`
	Title                   string    `gorm:"type:varchar(255);not null" json:"title"`
	Description             string    `gorm:"type:text" json:"description"`
	CoverImageURL           string    `gorm:"type:varchar(255)" json:"cover_image_url"`
	DeveloperName           string    `gorm:"type:varchar(100)" json:"developer_name"`
	Genre                   string    `gorm:"type:varchar(50)" json:"genre"`
	GameType                string    `gorm:"type:varchar(10)" json:"game_type"`
	InvestorWalletAddresses pq.StringArray `gorm:"type:text[]" json:"investor_wallet_addresses"` // Array of investor wallet addresses
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

// BeforeCreate hook untuk generate UUID v7 sebelum insert
func (p *Project) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// UserProfile merepresentasikan tabel user_profiles
type UserProfile struct {
	WalletAddress   string    `gorm:"type:varchar(42);primaryKey" json:"wallet_address"`
	Username        string    `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email           string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	ProfileImageURL string    `gorm:"type:varchar(255)" json:"profile_image_url"`
	KYCStatus       string    `gorm:"type:varchar(20);default:'unverified'" json:"kyc_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Comment merepresentasikan tabel comments
type Comment struct {
	ID                  uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	ProjectID           uuid.UUID  `gorm:"type:uuid;not null;index" json:"project_id"`
	AuthorWalletAddress string     `gorm:"type:varchar(42);not null" json:"author_wallet_address"`
	ParentCommentID     *uuid.UUID `gorm:"type:uuid;index" json:"parent_comment_id"`
	Content             string     `gorm:"type:text;not null" json:"content"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`

	// Relasi
	Project       Project  `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"-"`
	ParentComment *Comment `gorm:"foreignKey:ParentCommentID" json:"-"`
}

// BeforeCreate hook untuk generate UUID v7 sebelum insert
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// ExternalLink merepresentasikan tabel external_links
type ExternalLink struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ProjectID uuid.UUID `gorm:"type:uuid;not null;index" json:"project_id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"` // e.g., "Instagram", "Twitter", "Website"
	URL       string    `gorm:"type:varchar(500);not null" json:"url"` // The actual URL
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"-"`
}

// BeforeCreate hook untuk generate UUID v7 sebelum insert
func (e *ExternalLink) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
