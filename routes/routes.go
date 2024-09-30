package routes

import (
	"ikct-ed/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.RouterGroup) {
	api := router.Group("/api")
	{
		students := api.Group("/students")
		{
			students.GET("/list", controllers.GetStudentsList)
			students.POST("/add/csv", controllers.AddStudentsCSV)
			students.GET("/detail/:id", controllers.GetStudentDetail)
			students.POST("/add/sheet", controllers.AddStudentsFromSheet)
			students.PUT("/update/:id", controllers.UpdateStudentDetail)
			students.POST("/upload/img/:id", controllers.UploadImageofStudent)
			students.GET("/image/:id", controllers.GetImageData)
		}

		admin := api.Group("/user")
		{
			admin.POST("/create", controllers.CreateUser)
			admin.POST("/login", controllers.Login)
		}
	}

	v1 := router.Group("/v1", controllers.ValidatePageJWT)
	{
		v1.GET("/student/list", controllers.StudentListPage)
		v1.GET("/student/detail/:id", controllers.StudentDetailPage)
		v1.GET("/student/update/:id", controllers.UpdateStudentTemplate)
	}

}

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.Use(corsMiddleware())
	gin.SetMode(os.Getenv("GIN_MODE"))
	router.LoadHTMLGlob("./templates/**/*")
	router.Static("/css", "./css")
	router.Static("/static/img", "./static/img")
	router.Static("/js", "./js")

	router.GET("/", controllers.LoginForm)
	// Add all current URls
	AddRoutes(&router.RouterGroup)

	return router
}

// corsMiddleware sets the necessary headers for CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,token")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
