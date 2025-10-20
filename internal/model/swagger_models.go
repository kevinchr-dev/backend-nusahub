package model

import "time"

// ProjectSwagger is a swagger-friendly representation of Project used only
// for documentation generation. It mirrors model.Project but uses primitive
// slice types that swag can parse (e.g., []string for investor addresses).
type ProjectSwagger struct {
	ID                      string    `json:"id" example:"0199fb01-ae3c-7c26-b70a-8f221585ccb4"`
	CreatorWalletAddress    string    `json:"creator_wallet_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Title                   string    `json:"title" example:"Epic RPG Game"`
	Description             string    `json:"description" example:"An amazing blockchain RPG game"`
	CoverImageURL           string    `json:"cover_image_url" example:"https://example.com/cover.jpg"`
	DeveloperName           string    `json:"developer_name" example:"Epic Games Studio"`
	Genre                   string    `json:"genre" example:"RPG"`
	GameType                string    `json:"game_type" example:"web3"`
	InvestorWalletAddresses []string  `json:"investor_wallet_addresses" example:"[\"0xabc...\"]"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

// Generic API responses used for documentation
type ErrorResponse struct {
	Error string `json:"error" example:"Resource not found"`
}

type GenericMessage struct {
	Message string `json:"message" example:"Investor added successfully"`
}

// Request/Response models
type AddInvestorRequest struct {
	WalletAddress string `json:"wallet_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

type InvestorsResponse struct {
	Investors []string `json:"investors" example:"[\"0x111...\", \"0x222...\"]"`
}

// ProjectPatch is used for PATCH /projects/{id} - all fields optional
type ProjectPatch struct {
	Title         *string `json:"title,omitempty" example:"Updated Game Title"`
	Description   *string `json:"description,omitempty" example:"Updated description"`
	CoverImageURL *string `json:"cover_image_url,omitempty" example:"https://example.com/new-image.jpg"`
	DeveloperName *string `json:"developer_name,omitempty" example:"New Studio Name"`
	Genre         *string `json:"genre,omitempty" example:"Action RPG"`
	GameType      *string `json:"game_type,omitempty" example:"web3"`
}

// ProjectCreate represents fields required to create a project (request body)
type ProjectCreate struct {
	CreatorWalletAddress string `json:"creator_wallet_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Title                string `json:"title" example:"My Awesome Game"`
	Description          string `json:"description,omitempty" example:"This is an amazing Web3 game"`
	CoverImageURL        string `json:"cover_image_url,omitempty" example:"https://example.com/image.jpg"`
	DeveloperName        string `json:"developer_name,omitempty" example:"GameDev Studios"`
	Genre                string `json:"genre,omitempty" example:"RPG"`
	GameType             string `json:"game_type,omitempty" example:"web3"`
}

// ExternalLinkCreate represents fields required to create/update an external link
type ExternalLinkCreate struct {
	Name string `json:"name" example:"Instagram"`
	URL  string `json:"url" example:"https://instagram.com/mygame"`
}

// CommentCreate represents fields required to create a comment
type CommentCreate struct {
	AuthorWalletAddress string  `json:"author_wallet_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Content             string  `json:"content" example:"This project looks amazing!"`
	ParentCommentID     *string `json:"parent_comment_id,omitempty" example:"0199fb01-b44b-7a26-9625-3be2baf2d905"`
}
