package factory

import (
	"fmt"

	"cqrses/app/database"
	"cqrses/app/models"
)

// Migrate creates the tables and the database structure.
func Migrate() {
	if err := database.Gorm.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		models.User{},
	); err != nil {
		panic(fmt.Errorf("migration error: %s", err.Error()))
	}
}
