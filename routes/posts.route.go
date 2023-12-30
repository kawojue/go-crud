package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-crud/controllers"
)

func PostRoute(route *gin.Engine) {
	postRoute := route.Group("/api/post")
	{
		postRoute.GET("/", controllers.ReadPosts)
		postRoute.POST("/", controllers.CreatePost)
		postRoute.GET("/:id", controllers.ReadPost)
		postRoute.PUT("/:id", controllers.UpdatePost)
		postRoute.DELETE("/:id", controllers.DeletePost)
	}
}
