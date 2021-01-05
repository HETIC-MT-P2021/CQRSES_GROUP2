package database

import (
	"cqrses/models"

	"fmt"
)

// Migrate creates the tables and the database structure.
func Migrate() {
	if err := Gorm.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		models.User{},
	); err != nil {
		panic(fmt.Errorf("Migration error: %s", err.Error()))
	}
}
