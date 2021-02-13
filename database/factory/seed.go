package factory

import (
	"cqrses/database"
	"cqrses/helpers"
	"cqrses/models"
	"fmt"
)

// Seed seed the database
func Seed() {
	UserSeeder("Doe", "John", "admin@hetic.net", "root", true)
	UserSeeder("Impsum", "Laurène", "lorem@hetic.net", "root", false)
}

// UserSeeder seed users table.
func UserSeeder(lastName string, firstName string, email string, password string, admin bool) {
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		panic(fmt.Errorf("seeding error: %s", err.Error()))
	}

	database.Gorm.
		Where(models.User{Email: email}).
		FirstOrCreate(&models.User{
			LastName:  lastName,
			FirstName: firstName,
			Email:     email,
			Password:  hashedPassword,
			Admin:     admin,
		})
}
