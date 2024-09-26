package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginForm Render the view of the login form
func LoginForm(c *gin.Context) {
	//get user information from the cookie
	token, _ := c.Cookie("user-token")
	if token != "" {
		c.Redirect(http.StatusFound, "/v1/student/list")
	} else {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the login.html template
			"login.html",
			gin.H{
				"title": "Login page",
			},
		)
	}
}
