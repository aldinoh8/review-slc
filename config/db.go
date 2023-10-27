package config

import (
	"example/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=ftgo-sandbox port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.Customer{}, &entity.Product{})

	return db
}
