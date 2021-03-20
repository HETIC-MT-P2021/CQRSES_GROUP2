package helpers

import (
	"fmt"
	"time"

	"cqrses/app/models"

	"github.com/dgrijalva/jwt-go"
)

const expTime = time.Hour * 24

// GenerateTokenPair action a new JWT token for users.
func GenerateTokenPair(user models.User) (t string, rt string, err error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["admin"] = user.Admin
	claims["name"] = user.LastName + user.FirstName
	claims["exp"] = time.Now().Add(expTime).Unix()

	// Generate encoded token and send it as response.
	t, _ = token.SignedString([]byte("secret"))

	// Create token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	// Set Rt claims
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(expTime).Unix()

	// Generate encoded token and send it as response.
	rt, err = refreshToken.SignedString([]byte("secret"))

	return
}

// RefreshJWTToken refresh the JWT token for users.
func RefreshJWTToken(refreshToken string, user models.User) (t string, rt string, err error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int(claims["sub"].(float64)) == 1 {
			t, rt, err = GenerateTokenPair(user)
		}
	}

	return
}
