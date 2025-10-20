package database

import (
	"log"

	"github.com/kevinchr/web3-crowdfunding-api/internal/config"
	"github.com/kevinchr/web3-crowdfunding-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase menginisialisasi koneksi database dan melakukan migrasi
func InitDatabase(cfg *config.Config) error {
	var err error

	// Konfigurasi GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Buka koneksi ke database
	DB, err = gorm.Open(postgres.Open(cfg.GetDSN()), gormConfig)
	if err != nil {
		return err
	}

	log.Println("Database connection established")

	// Auto migrate semua model jika diizinkan oleh konfigurasi
	if cfg.MigrateOnStart {
		err = DB.AutoMigrate(
			&model.Project{},
			&model.UserProfile{},
			&model.Comment{},
			&model.ExternalLink{},
		)
		if err != nil {
			return err
		}

		log.Println("Database migration completed")
	} else {
		log.Println("Skipping auto-migration (MIGRATE_ON_START is false)")
	}

	return nil
}

// GetDB mengembalikan instance database
func GetDB() *gorm.DB {
	return DB
}
