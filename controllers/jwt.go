package controllers

import (
	"ikct-ed/models"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTAuthClaims Structure to store JWT
type JWTAuthClaims struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}

// GetTimeForCookies to get time of 30 min
func GetTimeForCookies() time.Time {
	return time.Now().Add((365 * 5) * 24 * time.Hour)
}

// CreateJWT will create JWT string and send
func CreateJWT(user models.User) (string, error) {
	expirationTime := GetTimeForCookies()

	claims := &JWTAuthClaims{
		Email:     user.Email,
		CreatedAt: time.Now(),
		Name:      user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Println("CreateJWT: failed with an error: ", err)
	}

	return tokenString, err
}
