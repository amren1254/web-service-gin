package router

import (
	"github.com/amren1254/gin-docker/auth/login"
	"github.com/amren1254/gin-docker/auth/signup"
	"github.com/amren1254/gin-docker/controller"
	"github.com/amren1254/gin-docker/middleware"

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
