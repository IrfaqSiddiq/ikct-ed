package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"ikct-ed/models"
	"ikct-ed/utility"

	"github.com/gin-gonic/gin"
)

//RBAC means ROLE BASED ACCESS CONTROL

func AuthorizationOfRoles2Permissions(c *gin.Context) {
	path := c.Request.URL.Path
	var name string
	switch {
	case strings.Contains(path, "/blog/"):
		name = "blog"
	case strings.Contains(path, "/job/"):
		name = "job"
	case strings.Contains(path, "/sitemap/"):
		name = "sitemap"
	case strings.Contains(path, "/user/"):
		name = "users"
	case strings.Contains(path, "/synthetic-job"):
		name = "synthetic_job"
	case strings.Contains(path, "/admin/"):
		name = "permissionNrole"
	case strings.Contains(path, "/bounty-company/"):
		name = "bounty_company"
	case strings.Contains(path, "/payout/") || strings.Contains(path, "/get-user-earnings"):
		name = "payout"
	case strings.Contains(path, "/headhunter/"):
		name = "associate"
	case strings.Contains(path, "/pipeline"):
		name = "pipeline"
	case strings.Contains(path, "/website"):
		name = "websites"
	case strings.Contains(path, "/seo-attribute"):
		name = "seo"
	case strings.Contains(path, "/supported-countries/"):
		name = "supportedcountry"
	case strings.Contains(path, "/news/"):
		name = "news"
	case strings.Contains(path, "/company/"):
		name = "company"
	default:
		log.Println("AuthorizationOfRoles2Permission: Couldn't find permission name from controller side")
	}

	method := c.Request.Method

	tokenString, _ := c.Cookie("tokenString")
	fmt.Println("*******tokenString", tokenString)
	user, err := models.GetUserProfileByToken(tokenString)
	if err != nil {
		log.Println("ValidatePageJWT : Failed to get user by user id with error : ", err)
		redirectURL := utility.GetHostURL()
		c.Redirect(http.StatusMovedPermanently, redirectURL)
		c.Abort()
		return
	}

	roleId, err := models.GetRoleIdByUserId(user.ID)
	if err != nil {
		log.Println("failed while getting roleID", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "GetRoleIdByUSerID failed while geeting userID"})
	}
	permissionID, err := models.GetPermissionId(name)
	if err != nil {
		log.Println("Failed while getting the permission name", err)

	}
	rbac, err := models.AuthorizationOfRoles2Permission(roleId, permissionID)
	if err != nil {
		log.Println("failed while getting ROLE at controller", err)
	}
	if method == "POST" && !rbac.Role2permission.Create {
		NoAccess(c)
		c.Abort()
		return
	}

	if method == "GET" && !rbac.Role2permission.Read {
		NoAccess(c)
		c.Abort()
		return

	}

	if method == "PUT" && !rbac.Role2permission.Update {
		NoAccess(c)
		c.Abort()
		return
	}

	if method == "DELETE" && !rbac.Role2permission.Delete {
		NoAccess(c)
		c.Abort()
		return
	}
	c.Next()
}

func NoAccess(c *gin.Context) {

	c.HTML(http.StatusOK, "authentication/no-access.html", gin.H{})
}
