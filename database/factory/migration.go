package factory

import (
	"cqrses/database"
	"cqrses/models"

	"fmt"
)

// Migrate creates the tables and the database structure.
func Migrate() {
	if err := database.Gorm.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		models.User{},
	); err != nil {
		panic(fmt.Errorf("Migration error: %s", err.Error()))
	}
}
