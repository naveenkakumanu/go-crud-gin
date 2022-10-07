package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	// Reading Parameters from commad line

	router := gin.Default()
	router.POST("/user/login", userLogin)
	router.POST("/user/register", userRegister)

	return router

}

func userLogin(c *gin.Context) {

}

func userRegister(c *gin.Context) {

}
