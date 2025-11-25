package config

import (
	"log"
	"os"

	"github.com/Baihaki123/customer-service/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// contoh pakai ENV atau static DSN
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		// ubah sesuai user/password/database
		dsn = "root:password@tcp(127.0.0.1:3306)/customer_db?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate (schema dari domain)
	if err := db.AutoMigrate(&domain.Customer{}, &domain.FamilyList{}, &domain.Nationality{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	return db
}
