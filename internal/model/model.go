package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/kevinchr/web3-crowdfunding-api/internal/idgen"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Project merepresentasikan tabel projects
type Project struct {
	ID                      uint64      `gorm:"type:bigint;primaryKey" json:"id"`
	CreatorWalletAddress    string      `gorm:"type:varchar(42);not null" json:"creator_wallet_address"`
	Title                   string      `gorm:"type:varchar(255);not null" json:"title"`
	Description             string      `gorm:"type:text" json:"description"`
	CoverImageURL           string      `gorm:"type:varchar(255)" json:"cover_image_url"`
	DeveloperName           string      `gorm:"type:varchar(100)" json:"developer_name"`
	Genre                   string      `gorm:"type:varchar(50)" json:"genre"`
	GameType                string      `gorm:"type:varchar(10)" json:"game_type"`
	InvestorWalletAddresses StringArray `gorm:"type:text[]" json:"investor_wallet_addresses" swaggertype:"[]string"` // Array of investor wallet addresses
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
}

// StringArray is a thin alias over pq.StringArray that implements
// sql.Scanner, driver.Valuer and JSON marshal/unmarshal so it works
// with GORM, pq and also is recognized by swag when annotated with
// `swaggertype:"[]string"`.
type StringArray pq.StringArray

// Value implements driver.Valuer
func (a StringArray) Value() (driver.Value, error) {
	return pq.StringArray(a).Value()
}

// Scan implements sql.Scanner
func (a *StringArray) Scan(src interface{}) error {
	var tmp pq.StringArray
	if err := tmp.Scan(src); err != nil {
		return err
	}
	*a = StringArray(tmp)
	return nil
}

// MarshalJSON converts to JSON array
func (a StringArray) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(a))
}

// UnmarshalJSON converts from JSON array
func (a *StringArray) UnmarshalJSON(b []byte) error {
	var s []string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*a = StringArray(s)
	return nil
}

// BeforeCreate hook untuk generate numeric timestamped ID sebelum insert
func (p *Project) BeforeCreate(tx *gorm.DB) error {
	if p.ID == 0 {
		p.ID = idgen.Generate()
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
	ID                  uint64    `gorm:"type:bigint;primaryKey" json:"id"`
	ProjectID           uint64    `gorm:"type:bigint;not null;index" json:"project_id"`
	AuthorWalletAddress string    `gorm:"type:varchar(42);not null" json:"author_wallet_address"`
	ParentCommentID     *uint64   `gorm:"type:bigint;index" json:"parent_comment_id"`
	Content             string    `gorm:"type:text;not null" json:"content"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	// Relasi
	Project       Project  `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"-"`
	ParentComment *Comment `gorm:"foreignKey:ParentCommentID" json:"-"`
}

// BeforeCreate hook untuk generate numeric timestamped ID sebelum insert
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if c.ID == 0 {
		c.ID = idgen.Generate()
	}
	return nil
}

// ExternalLink merepresentasikan tabel external_links
type ExternalLink struct {
	ID        uint64    `gorm:"type:bigint;primaryKey" json:"id"`
	ProjectID uint64    `gorm:"type:bigint;not null;index" json:"project_id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"` // e.g., "Instagram", "Twitter", "Website"
	URL       string    `gorm:"type:varchar(500);not null" json:"url"` // The actual URL
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"-"`
}

// BeforeCreate hook untuk generate numeric timestamped ID sebelum insert
func (e *ExternalLink) BeforeCreate(tx *gorm.DB) error {
	if e.ID == 0 {
		e.ID = idgen.Generate()
	}
	return nil
}
