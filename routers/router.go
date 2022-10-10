package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naveenkakumanu/go-crud-gin/controllers"
)

// Router to Prepare all the routes
func Router() *gin.Engine {

	// Reading Parameters from commad line

	router := gin.Default()
	// Grouping Routes
	user := router.Group("v1/user")
	book := router.Group("/v1/books")
	// user Routes
	user.POST("/login", controllers.UserLogin)
	user.POST("/register", controllers.UserRegister)

	// Book Routes
	book.POST("/create", controllers.CreateBook)
	book.POST("/update", controllers.UpdateBook)
	book.POST("/:id", controllers.FindBook)
	book.POST("/archive/:id", controllers.ArchiveBook)
	return router

}
