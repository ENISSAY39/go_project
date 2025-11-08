package intializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ENISSAY39/go_project/models"
)

var DB *gorm.DB

func ConnectToDB() {
	// Database connection logic goes here
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
	log.Println("Connected to database successfully")
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
