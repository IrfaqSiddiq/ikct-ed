package controllers

import (
	"fmt"
	"ikct-ed/models"
	"ikct-ed/utility"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func ValidatePageJWT(c *gin.Context) {

	tokenString, _ := c.Cookie("tokenString")
	fmt.Println("*******tokenString", tokenString)
	_, err := models.GetUserProfileByToken(tokenString)
	if err != nil {
		log.Println("ValidatePageJWT : Failed to get user by user id with error : ", err)
		redirectURL := utility.GetHostURL()
		c.Redirect(http.StatusMovedPermanently, redirectURL)
		c.Abort()
		return
	}
	c.Next()

}

func ValidateAPIJWT(c *gin.Context) {
	tokenString, err := c.Cookie("tokenString")
	if err != nil {
		log.Println("ValidateAPIJWT: failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to get token.",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	Valid, err := VerifyJWT(tokenString)
	log.Println("tokenString :	", Valid)
	if !Valid || err != nil {
		log.Println("ValidateAPIJWT: failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to verify token.",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	Valid, err = models.CheckSession(tokenString)
	if err != nil {
		log.Println("ValidateAPIJWT: failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to check session of this token.",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	if !Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to get active session.",
		})
		c.Abort()
		return
	}
	_, err = models.GetUserProfileByToken(tokenString)
	if err != nil {
		log.Println("ValidateAPIJWT : Failed to get user by user id with error : ", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to get user",
			"error":   err.Error(),
		})
		c.Abort()
	}
	c.Next()
}

// VerifyJWT will verify JWT key
func VerifyJWT(tokenString string) (bool, error) {
	claims := &JWTAuthClaims{}
	var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Println("error at ParseWithClaims : ", err.Error())
		return false, err
	}

	return tkn.Valid, err
}
