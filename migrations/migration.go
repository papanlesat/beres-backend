package migrations

import (
	"beres/infra/database"
	"beres/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{
		&models.User{},
		&models.Post{},
		&models.Category{},
		&models.Tag{},
		&models.Menu{},
		&models.MenuItem{},
		&models.Setting{},
		&models.Widget{},
		&models.PersonalAccessToken{},
		&models.Section{},
	}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
