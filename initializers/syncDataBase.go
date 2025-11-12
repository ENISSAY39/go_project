package initializers

import (
	"github.com/ENISSAY39/go_project/models"
)

func SyncDataBase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
