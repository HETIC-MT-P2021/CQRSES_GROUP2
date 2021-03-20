package models

import (
	"time"

	"gorm.io/gorm"
)

// User defines the structure of the user entity
type User struct {
	gorm.Model        // Contain ID primary Key and Timestamps
	LastName   string `form:"last_name" gorm:"not null; type:varchar(55)"`
	FirstName  string `form:"first_name" gorm:"not null; type:varchar(55)"`
	Email      string `form:"email" gorm:"unique; not null; type:varchar(55)"`
	Password   string `form:"password" gorm:"not null; type:varchar(255)"`
	Admin      bool   `form:"admin" gorm:"default:0; type:boolean"`
}

// UserGet defines the structure of the user API response entity
type UserGet struct {
	ID        uint      `json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Admin     bool      `json:"admin"`
}
