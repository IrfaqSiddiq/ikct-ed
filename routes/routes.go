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
		student := api.Group("/student", controllers.AuthorizationOfRoles2PermissionsAPI)
		{
			student.GET("/list", controllers.GetStudentsList)
			student.POST("/add/csv", controllers.AddStudentsCSV)
			student.GET("/detail/:id", controllers.GetStudentDetail)
			student.POST("/add/sheet", controllers.AddStudentsFromSheet)
			student.PUT("/update/:id", controllers.UpdateStudentDetail)
			student.POST("/upload/img/:id", controllers.UploadImageofStudent)
			student.GET("/image/:id", controllers.GetImageData)
			student.DELETE("/delete/img/:id", controllers.DeleteImageOfStudent)
			student.POST("/logout", controllers.Logout)
			student.POST("/insert", controllers.AddStudentRecord)

		}

		admin := api.Group("/user", controllers.AuthorizationOfRoles2PermissionsAPI)
		{
			admin.DELETE("/delete/:id", controllers.DeleteAdminByID)
			admin.GET("/list", controllers.GetAdminList)
			admin.POST("/create", controllers.CreateUser)
			admin.GET("/detail", controllers.GetAdminDetails)
		}
		schools := api.Group("/school", controllers.AuthorizationOfRoles2PermissionsAPI)
		{
			schools.GET("/list", controllers.GetSchoolList)
			schools.POST("/add", controllers.AddSchool)
			schools.DELETE("/delete/:id", controllers.DeleteSchool)
		}
		api.POST("/login", controllers.Login)
		api.GET("/religion", controllers.GetReligions)

	}

	v1 := router.Group("/v1", controllers.AuthorizationOfRoles2PermissionsPage)
	{
		v1.GET("/student/list", controllers.StudentListPage)
		v1.GET("/student/detail/:id", controllers.StudentDetailPage)
		v1.GET("/student/update/:id", controllers.UpdateStudentTemplate)
		v1.GET("/student/add", controllers.InsertStudentPage)
		v1.GET("/school/list", controllers.SchoolPage)
		v1.GET("/admin/list", controllers.AdminListPage)

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
