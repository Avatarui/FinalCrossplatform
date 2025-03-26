package routes

import (
	"FinalCrossplatform/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/auth/login", controllers.Login)
	return r
}
