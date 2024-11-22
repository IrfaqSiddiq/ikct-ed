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

func AuthorizationOfRoles2PermissionsPage(c *gin.Context) {
	path := c.Request.URL.Path
	var name string
	switch {
	case strings.Contains(path, "/student/"):
		name = "student"
	case strings.Contains(path, "/school/"):
		name = "school"
	case strings.Contains(path, "/user/"):
		name = "user"
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
	fmt.Println("permission name", name)
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

func AuthorizationOfRoles2PermissionsAPI(c *gin.Context) {
	path := c.Request.URL.Path
	var name string
	switch {
	case strings.Contains(path, "/student/"):
		name = "student"
	case strings.Contains(path, "/school/"):
		name = "school"
	case strings.Contains(path, "/user/"):
		name = "user"
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "GetRoleIdByUSerID failed while geeting userID"})
	}
	permissionID, err := models.GetPermissionId(name)
	if err != nil {
		log.Println("Failed while getting the permission name", err)

	}
	rbac, err := models.AuthorizationOfRoles2Permission(roleId, permissionID)
	if err != nil {
		log.Println("failed while getting ROLE at controller", err)
	}
	fmt.Println("permissions", rbac)
	if method == "POST" && !rbac.Role2permission.Create {
		// NoAccess(c)

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  "you are not authorized",
		})
		return
	}

	if method == "GET" && !rbac.Role2permission.Read {
		// NoAccess(c)

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  "you are not authorized",
		})
		return

	}

	if method == "PUT" && !rbac.Role2permission.Update {
		// NoAccess(c)

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  "you are not authorized",
		})
		return
	}

	if method == "DELETE" && !rbac.Role2permission.Delete {
		// NoAccess(c)

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "fail",
			"error":  "you are not authorized",
		})
		return
	}
	c.Set("permissions", rbac)
	c.Next()
}

func NoAccess(c *gin.Context) {

	c.HTML(http.StatusOK, "authentication/no-access.html", gin.H{})
}
