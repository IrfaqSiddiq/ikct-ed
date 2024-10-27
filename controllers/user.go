package controllers

import (
	"errors"
	"fmt"
	"ikct-ed/models"
	"ikct-ed/services"
	"ikct-ed/utility"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	//create a variable of type User to receive data from the form
	var newUser models.User
	if c.ShouldBind(&newUser) != nil {
		c.JSON(http.StatusBadRequest, "Error trying get form data!")
		return
	}
	CheckUser, err := models.CheckUser(strings.ToLower(newUser.Email))
	if CheckUser {
		fmt.Println("User already exists with this email")
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "error": err, "message": "User already exists with this email !!!"})
	} else {
		//Password encryption
		passwordEncripted, err := PasswordEncrypter(newUser.Password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "failed", "error": err, "message": "Use another password !!!"})
		} else {
			userStorage := models.UserStorage{}
			err = userStorage.User.New(newUser, passwordEncripted)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"status": "failed", "error": err, "message": "User not created !!!"})
			} else {
				//return to the user list view
				c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User Created successfully!!!"})
			}
		}
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	email = strings.ToLower(email)
	password := c.PostForm("password")

	if len(email) == 0 {
		log.Println("Login Failed: Field Email is missing. Email is mandatory")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errors.New("email is missing"),
		})
		return
	}

	if len(password) == 0 {
		log.Println("Login Failed: Field password is missing. password is mandatory")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errors.New("password is missing"),
		})
		return
	}

	user, err := models.GetAdminDetails(email)
	if err != nil {
		log.Println("Login Failed: while fetching user details with error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Email entered is incorrect",
			"error":   "invalid email",
		})
		return
	}

	passwordMatched := utility.VerifyPassword(user.Password, password)
	if !passwordMatched {
		log.Println("Login Failed: Invalid password. Please enter correct password ")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed while verifying password",
			"error":   "invalid password",
		})
		return
	}

	token := CreateSession(c, &user)

	err = user.UserAuth.StoreJwtSessionInDB()
	if err != nil {
		log.Println("Login Failed: to save token in database with error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to store session details",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "successfully logged in",
		"user_id": user.ID,
	})

}

func Logout(c *gin.Context) {
	tokenString, err := c.Cookie("tokenString")
	if err != nil {
		log.Println("[UN-AUTHORIZED] TokenAuthMiddleware: Failed to fetch student details by token with error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err})
		return
	}
	log.Println("Token ", tokenString)
	log.Println("header", c.Request.Header)

	if len(tokenString) == 0 {
		log.Println("[UN-AUTHORIZED] Logout: No token provided.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "No Token provided",
		})
		return
	}

	// using token fetch student's information.
	_, err = models.GetUserProfileByToken(tokenString)
	if err != nil {
		log.Println("[UN-AUTHORIZED] TokenAuthMiddleware: Failed to fetch student details by token with error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "message": "Failed to fetch student", "error": err})
		return
	}
	models.ExpireSession(tokenString)
	services.RemoveCookies(c, "tokenString")
}

// PasswordEncrypter Encrypts the password and returns a text string
func PasswordEncrypter(password string) (string, error) {
	//Password encryption
	hash, err := (bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost))
	if err != nil {
		msgError := fmt.Errorf("PasswordEncrypter:bcrypt.GenerateFromPassword: Error: %v", err.Error())
		return "", msgError
	}
	//Hash value to string
	password = string(hash)
	//return password encripted
	return password, nil
}

// CreateSession setting cookies as per jwt
func CreateSession(c *gin.Context, user *models.User) string {
	var SessionID string

	SessionID, _ = CreateJWT(*user)
	user.UserAuth = models.UserAuth{
		UserID:   int64(user.ID),
		JWTToken: SessionID,
	}
	maxAge := 10000

	c.SetCookie("tokenString", SessionID, maxAge, "/", os.Getenv("DOMAIN"), false, true)
	fmt.Println("******cookie set")
	return SessionID
}
