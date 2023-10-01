package database

import (
	"fmt"
	"log"

	"github.com/Kozzen890/assignment2-016/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	host     = "localhost"
	user     = "postgres"
	password = "postgres890"
	port     = 5432
	dbname   = "apps_order"
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error:", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDatabase() *gorm.DB {
	return db
}