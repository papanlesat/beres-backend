package routers

import (
	"beres/controllers"
	"beres/routers/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoute() *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	sections := router.Group("/sections")
	{
		sections.GET("", controllers.GetSectionData)     // GET    /sections
		sections.GET("/:id", controllers.GetSectionByID) // GET    /sections/:id
	}
	auth := router.Group("/")
	auth.Use(middleware.TokenAuth())
	{
		auth.POST("/logout", controllers.Logout)
		sections := auth.Group("/sections")
		{
			sections.POST("", controllers.CreateSection)       // POST   /sections
			sections.PUT("/:id", controllers.UpdateSection)    // PUT    /sections/:id
			sections.DELETE("/:id", controllers.DeleteSection) // DELETE /sections/:id
		}
	}

	RegisterRoutes(router) //routes register

	return router
}
