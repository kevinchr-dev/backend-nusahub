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
