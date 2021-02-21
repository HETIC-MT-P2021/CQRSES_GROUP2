package controllers

import (
	"net/http"

	db "cqrses/app/database"
	"cqrses/app/helpers"
	"cqrses/app/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
)

// Login verifies the identifiers and connects the user by creating a token.
func Login(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Register error", err.Error()))
	}

	if err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.EmailFormat),
		validation.Field(&user.Password, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	email := user.Email
	password := user.Password

	if err := db.Gorm.Where("email = ?", email).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Connexion error", err.Error()))
	}

	if match := helpers.CheckPasswordHash(password, user.Password); !match {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Connexion error", "Bad password"))
	}

	t, rt, err := helpers.GenerateTokenPair(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Connexion error", err.Error()))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "User connected", map[string]string{
		"refresh_token": rt,
		"token":         t,
	}))
}

// Register create a new user in the database
func Register(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Register error", err.Error()))
	}

	if err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.EmailFormat),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.LastName, validation.Required),
		validation.Field(&user.FirstName, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Validation error", err.Error()))
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Register error", err.Error()))
	}

	user.Password = hashedPassword

	if err := db.Gorm.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Register error", err.Error()))
	}

	return c.JSON(http.StatusCreated, SetResponse(http.StatusCreated, "User registered", user.ID))
}

// RefreshToken refresh token
func RefreshToken(c echo.Context) error {
	refreshToken := c.FormValue("refresh_token")
	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "JWT Refresh error", "Missing refresh token"))
	}

	t, rt, err := helpers.RefreshJWTToken(refreshToken, models.User{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "JWT Refresh error", err.Error()))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "JWT refreshed", map[string]string{
		"refresh_token": rt,
		"token":         t,
	}))
}
