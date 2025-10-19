.PHONY: help run build clean test dev migrate swagger

help: ## Menampilkan bantuan
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Menjalankan aplikasi
	go run cmd/main/main.go

build: ## Build aplikasi
	go build -o bin/api cmd/main/main.go

swagger: ## Generate Swagger documentation
	swag init -g cmd/main/main.go --output docs
	@echo "Swagger docs generated! View at http://localhost:3000/docs/"

dev: ## Menjalankan aplikasi dengan auto-reload (requires air)
	air

clean: ## Membersihkan binary
	rm -rf bin/
	rm -rf docs/

test: ## Menjalankan tests
	go test -v ./...

deps: ## Install dependencies
	go mod download
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

migrate: ## Menjalankan migrasi database (auto-migrate saat aplikasi start)
	@echo "Migration akan berjalan otomatis saat aplikasi start"

docker-db: ## Menjalankan PostgreSQL dengan Docker
	docker run --name web3-crowdfunding-db \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=web3_crowdfunding \
		-p 5432:5432 \
		-d postgres:15-alpine

docker-db-stop: ## Stop PostgreSQL container
	docker stop web3-crowdfunding-db
	docker rm web3-crowdfunding-db
