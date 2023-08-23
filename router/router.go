package router

import (
	"web-service-gin/auth/login"
	"web-service-gin/auth/signup"
	"web-service-gin/controller"
	"web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	//initializing database here
	// db := database.NewDatabaseRepository(&sql.DB{}).InitDatabaseConnection()

	//unauthorized access
	router.POST("/signup", signup.Signup)
	router.POST("/login", login.Login)

	//authorized access to apis
	public := router.Group("/api")
	public.Use(middleware.JwtAuthMiddleware())
	public.GET("/albums", controller.GetAlbum)
	public.GET("/albums/:id", controller.GetAlbumsById)
	public.POST("/albums", controller.PostAlbum)

	return router
}
