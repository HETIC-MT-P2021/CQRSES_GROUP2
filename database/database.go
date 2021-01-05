package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gorm public declaration
var Gorm *gorm.DB

// GetSQLConfig return config variables of the database.
func GetSQLConfig() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
}

// Connect connect to the database.
func Connect() {
	var err error

	dsn := GetSQLConfig()

	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal error database connexion: %s", err.Error()))
	}

	sqlDB, err := Gorm.DB()
	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("Fatal error database connexion: %s", err.Error()))
	}
}
