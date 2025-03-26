package routes

import (
	"FinalCrossplatform/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/changepassword", controllers.ChangePassword)
	// r.POST("/auth/register", controllers.Register)
	return r
}
