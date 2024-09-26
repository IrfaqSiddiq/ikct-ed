package controllers

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"ikct-ed/models"
	"io"
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
	studentsInfo, err := models.GetStudentsList(currentPage)
	if err != nil {
		log.Println("GetStudentsList: Failed to get students details with error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed while fetching students list",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"students_info": studentsInfo,
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
	// err = os.Remove(filePath)
	// if err != nil {
	// 	log.Println("AddStudentsCSV: Failed to delete the temporary file with error: ", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  "fail",
	// 		"message": "failed to delete temporary file",
	// 		"error":   err,
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully inserted records",
	})
}

func StudentListPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the login.html template
		"student_list.html",
		gin.H{
			"title": "Student List page",
		},
	)
}

func StudentDetailPage(c *gin.Context) {
	// Get the dynamic id from the path
	studentID := c.Param("id")

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the student_detail.html template
		"student_detail.html",
		gin.H{
			"title":      "Student Detail page",
			"student_id": studentID, // Pass the student ID to the template
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
