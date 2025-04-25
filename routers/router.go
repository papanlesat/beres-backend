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

	posts := router.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)        // GET    /posts      (list)
		posts.GET("/:id", controllers.GetPostByID) // GET    /posts/:id  (retrieve)
	}

	// Categories CRUD
	categories := router.Group("/categories")
	{
		categories.GET("", controllers.GetCategories)
		categories.GET("/:id", controllers.GetCategoryByID)
		categories.POST("", controllers.CreateCategory)
		categories.PUT("/:id", controllers.UpdateCategory)
		categories.DELETE("/:id", controllers.DeleteCategory)
	}

	// Tags CRUD
	tags := router.Group("/tags")
	{
		tags.GET("", controllers.GetTags)
		tags.GET("/:id", controllers.GetTagByID)
		tags.POST("", controllers.CreateTag)
		tags.PUT("/:id", controllers.UpdateTag)
		tags.DELETE("/:id", controllers.DeleteTag)
	}

	menus := router.Group("/menus")
	{
		menus.GET("", controllers.GetMenus)
		menus.GET("/:id", controllers.GetMenuByID)
		menus.POST("", controllers.CreateMenu)
		menus.PUT("/:id", controllers.UpdateMenu)
		menus.DELETE("/:id", controllers.DeleteMenu)
	}

	menuItems := router.Group("/menus/:menu_id/items")
	{
		menuItems.GET("", controllers.GetMenuItems)
	}

	items := router.Group("/items")
	{
		items.GET("/:id", controllers.GetMenuItemByID)
		items.POST("", controllers.CreateMenuItem)
		items.PUT("/:id", controllers.UpdateMenuItem)
		items.DELETE("/:id", controllers.DeleteMenuItem)
	}

	widgets := router.Group("/widgets")
	{
		widgets.GET("", controllers.GetWidgets)
		widgets.GET("/:id", controllers.GetWidgetByID)
		widgets.POST("", controllers.CreateWidget)
		widgets.PUT("/:id", controllers.UpdateWidget)
		widgets.DELETE("/:id", controllers.DeleteWidget)
	}

	settings := router.Group("/settings")
	{
		settings.GET("", controllers.GetSettings)
		settings.GET("/:id", controllers.GetSettingByID)
		settings.POST("", controllers.CreateSetting)
		settings.PUT("/:id", controllers.UpdateSetting)
		settings.DELETE("/:id", controllers.DeleteSetting)
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
		posts := auth.Group("/posts")
		{
			posts.POST("", controllers.CreatePost)       // POST   /posts      (create)
			posts.PUT("/:id", controllers.UpdatePost)    // PUT    /posts/:id  (update)
			posts.DELETE("/:id", controllers.DeletePost) // DELETE /posts/:id  (delete)
		}
	}

	RegisterRoutes(router) //routes register

	return router
}
