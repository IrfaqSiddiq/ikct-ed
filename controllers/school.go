package controllers

import (
	"encoding/json"
	"fmt"
	"ikct-ed/models"
	"ikct-ed/utility"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSchoolList(c *gin.Context) {
	name := c.Query("school")
	page := c.Query("page")
	limit := c.Query("limit")

	permissions, _ := c.Get("permissions")
	data, err := json.Marshal(permissions)
	if err != nil {
		log.Println("GetStudentsList: Failed to marshal permission")
	}

	var rbac models.RBAC

	err = json.Unmarshal(data, &rbac)
	if err != nil {
		log.Println("GetStudentsList: Failed to unmarshal permission")
	}

	var currentPage int64

	currentPage, err = strconv.ParseInt(page, 10, 64)
	if len(page) == 0 || err != nil {
		currentPage = 1
	}

	l, err := strconv.ParseInt(limit, 10, 64)
	if len(limit) == 0 || err != nil {
		log.Println("GetSchoolList: Failed to get limit of schools")
	}

	totalCount, schools, err := models.GetSchoolList(currentPage, name, l)
	if err != nil {
		log.Println("GetSchoolList: Failed to get school information with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get school info",
			"error":   err,
		})
		return
	}
	totalPages := ""
	if totalCount%10 == 0 {
		totalPages = fmt.Sprintf("%v", (totalCount / 10))
	} else {
		totalPages = fmt.Sprintf("%v", (totalCount/10)+1)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"message":     "successfully fetched school list",
		"schools":     schools,
		"total_page":  totalPages,
		"permissions": rbac,
	})
}

func AddSchool(c *gin.Context) {
	school := c.PostForm("school")
	if len(school) == 0 {
		log.Println("AddSchool: School is Mandatory field: ")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"error":   "School is missing",
			"message": "school is mandatory field",
		})
		return
	}

	err := models.AddSchool(school)
	if err != nil {
		log.Println("AddSchool: Failed to add school information in database with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"error":   err.Error(),
			"message": "failed to add school",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully added school",
	})
}

func SchoolPage(c *gin.Context) {
	hostURL := utility.GetHostURL()
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the student_detail.html template
		"schools.html",
		gin.H{
			"title":    "Test page",
			"host_url": hostURL,
		},
	)
}
