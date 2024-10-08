package services

import (
	"os"

	"github.com/gin-gonic/gin"
)

// RemoveCookies used to remove cookies
func RemoveCookies(c *gin.Context, cookiesName string) {
	domain := os.Getenv("DOMAIN")
	c.SetCookie(cookiesName, "", -1, "/", domain, false, true)
}
