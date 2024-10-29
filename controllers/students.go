package controllers

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"ikct-ed/models"
	"ikct-ed/utility"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetStudentsList(c *gin.Context) {
	pageNo := c.Query("page")
	currentPage, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil || currentPage == 0 {
		log.Println("GetStudentsList: Failed to covert  pageNo into integer")
		currentPage = 1
	}
	filter := models.FilterParameters{}

	filter.SearchText = c.Query("search")
	religion := c.Query("religion")
	schools := c.Query("school")
	assistance := c.Query("assistance")

	if len(religion) > 0 {
		filter.Religion = strings.Split(religion, ",")
	}

	if len(schools) > 0 {
		filter.Schools = strings.Split(schools, ",")
	}

	if len(assistance) > 0 {
		filter.Assistance = strings.Split(assistance, ",")
	}

	studentsInfo, totalCount, err := models.GetStudentsList(currentPage, filter)
	if err != nil {
		log.Println("GetStudentsList: Failed to get students details with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed while fetching students list",
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
		"status":        "success",
		"students_info": studentsInfo,
		"total_page":    totalPages,
	})
}

func AddStudentsCSV(c *gin.Context) {
	// Parse multipart form to get the file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("AddStudentsCSV: failed to read file with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to read file",
			"error":   err,
		})
		return
	}
	defer file.Close()

	// Read file content into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file) // Copy the file into buffer
	if err != nil {
		log.Println("AddStudentsCSV: failed to copy file with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to copy file content",
			"error":   err,
		})
		return
	}

	// Parse CSV
	reader := csv.NewReader(bytes.NewReader(buf.Bytes()))
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("AddStudentsCSV: invalid csv format with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CSV format"})
		return
	}

	// Regular expression to match numeric values with commas
	re := regexp.MustCompile(`(\d+),(\d+)`)

	// Clean up commas in numeric values
	for i, record := range records {
		for j, field := range record {
			// Replace commas in numeric values
			if re.MatchString(field) {
				records[i][j] = strings.ReplaceAll(field, ",", "")
			}
		}
	}

	// Define a path to temporarily save the uploaded file
	filePath := "./uploads/" + header.Filename

	// Save the file on the server
	out, err := os.Create(filePath)
	if err != nil {
		log.Println("AddStudentsCSV: failed to create temporary file on server with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to create temporary file",
			"error":   err,
		})
		return
	}
	defer out.Close()

	// Write the cleaned CSV back to the file
	writer := csv.NewWriter(out)
	err = writer.WriteAll(records)
	if err != nil {
		log.Println("AddStudentsCSV: failed to write cleaned CSV content with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to write cleaned CSV content",
			"error":   err,
		})
		return
	}
	writer.Flush()

	// Reopen the file for database insertion
	file, err = os.Open(filePath)
	if err != nil {
		log.Println("AddStudentsCSV: failed to open file with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to open file for DB insertion",
			"error":   err,
		})
		return
	}
	defer file.Close()

	// // Create a new scanner to read the file
	// scanner := bufio.NewScanner(file)

	// // Loop through the file line by line
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text()) // Print each line
	// }

	// // Check for errors in scanning
	// if err := scanner.Err(); err != nil {
	// 	fmt.Println("Error reading file:", err)
	// }

	// Call a function to insert the CSV into the database using \copy
	err = models.InsertCSVIntoDB(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "failed to insert csv records",
			"error":   err,
		})
		return
	}
	// COMMENT OUT TO MAKE SURE FILE EXISTS ON SERVER
	err = os.Remove(filePath)
	if err != nil {
		log.Println("AddStudentsCSV: Failed to delete the temporary file with error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "failed to delete temporary file",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully inserted records",
	})
}

func StudentListPage(c *gin.Context) {

	hostURL := utility.GetHostURL()
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the login.html template
		"student_list.html",
		gin.H{
			"title":    "Student List page",
			"host_url": hostURL,
		},
	)
}

func StudentDetailPage(c *gin.Context) {
	// Get the dynamic id from the path
	studentID := c.Params.ByName("id")
	hostURL := utility.GetHostURL()
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the student_detail.html template
		"student_detail.html",
		gin.H{
			"title":      "Student Detail page",
			"student_id": studentID,
			"host_url":   hostURL,
		},
	)
}

func UpdateStudentTemplate(c *gin.Context) {
	// Get the dynamic id from the path
	studentID := c.Params.ByName("id")
	hostURL := utility.GetHostURL()
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the student_detail.html template
		"student_update.html",
		gin.H{
			"title":      "Student Detail page",
			"student_id": studentID,
			"host_url":   hostURL,
		},
	)
}

func GetStudentDetail(c *gin.Context) {
	studentID, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	if err != nil || studentID == 0 {
		log.Println("GetStudentDetail: Failed to get student id ")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get student id",
		})
		return
	}
	studentInfo, err := models.GetStudentDetail(studentID)
	if err != nil {
		log.Println("GetStudentDetail: Failed to get student details with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get student details",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"student_info": studentInfo,
	})
}

func UploadImageofStudent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println("UploadImageofStudent: Failed to get student id with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get student id",
			"error":   err,
		})
		return
	}
	profilePic, fileHandler, err := c.Request.FormFile("profile_pic")
	if err != nil {
		log.Println("UploadImageofStudent: Failed to get profile image with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get profile picture",
			"error":   err,
		})
		return
	}

	defer profilePic.Close()

	// Read the image file into a byte array
	imageData, err := ioutil.ReadAll(profilePic)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read the image: %v", err)
		return
	}

	err = models.UploadImageofStudent(imageData, id)
	if err != nil {
		log.Println("UploadImageofStudent: Failed to upload profile image with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to upload profile picture",
			"error":   err,
		})
		return
	}
	fmt.Printf("Uploaded file: %s, Size: %d bytes\n", fileHandler.Filename, len(imageData))
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully uploaded image",
	})
}

func DeleteImageOfStudent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println("DeleteImageOfStudent: Failed to get student id with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get student id",
			"error":   err,
		})
		return
	}

	err = models.UploadImageofStudent([]byte{}, id)
	if err != nil {
		log.Println("DeleteImageOfStudent: Failed to delete profile image with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to delete profile picture",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully deleted image",
	})
}

func UpdateStudentDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println("UpdateStudentDetail: Failed to get student id with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get student id",
			"error":   err,
		})
		return
	}
	var studentInfo models.StudentsFinancialInfo
	err = c.ShouldBindJSON(&studentInfo)

	if err != nil {
		log.Println("UpdateStudentDetail: Failed to bind json in studentInfo struct with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to bind json",
			"error":   err,
		})
		return
	}

	studentInfo.Id = id
	err = models.UpdateStudentDetail(studentInfo)

	if err != nil {
		log.Println("UpdateStudentDetail: Failed to update student detail with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Failed to update student detail",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "succesfully updated student details",
	})
}

func AddStudentsFromSheet(c *gin.Context) {

	const (
		googleSheetsSpreadsheetID = "17Mh4qPgvaTab2rx6r-aZZn3OkitAwY2wwF7HONF7iIM"
		googleSheetsRange         = "A1:Z" // Read all columns and rows
	)

	ctx := context.Background()
	credsFile := "./credentials.json" // Path to your Google Sheets API credentials
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credsFile))
	if err != nil {
		log.Println("Unable to retrieve Sheets client: ", err)
	}

	values := readGoogleSheetsValues(srv, googleSheetsSpreadsheetID, googleSheetsRange)
	fmt.Println("************values********", values)
}

func readGoogleSheetsValues(sheetsService *sheets.Service, spreadsheetID string, rangeName string) [][]interface{} {
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, rangeName).Do()
	if err != nil {
		log.Println("readGoogleSheetsValues: Unable to retrieve data from sheet: ", err)
	}
	return resp.Values
}

func GetImageData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Println("GetImageData: Failed to get student id with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get student id",
			"error":   err,
		})
		return
	}

	imageData, err := models.GetImageData(id)
	if err != nil {
		log.Println("GetImageData: Failed to get image with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get image",
			"error":   err,
		})
		return
	}
	c.Data(http.StatusOK, "image/jpeg", imageData)

}

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

func GetReligions(c *gin.Context) {
	religion, err := models.GetReligions()
	if err != nil {
		log.Println("GetReligions: Failed to get religion information with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed to get religion info",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"message":  "successfully fetched religion list",
		"religion": religion,
	})
}
