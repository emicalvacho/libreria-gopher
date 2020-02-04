package router

import (
	"api_gopher_library/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {

	// CRUD users
	router.POST("/create_user", controllers.CreatingUser)
	router.GET("/get_user", controllers.GettingUser)
	router.GET("/get_all_users", controllers.GettingUsers)

}

func Run() {
	router.Run(":8080")
}
