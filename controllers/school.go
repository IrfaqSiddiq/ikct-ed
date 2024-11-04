package controllers

import (
	"ikct-ed/models"
	"ikct-ed/utility"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSchoolList(c *gin.Context) {
	schools, err := models.GetSchoolList()
	if err != nil {
		log.Println("GetSchoolList: Failed to get school information with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get school info",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully fetched school list",
		"schools": schools,
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
