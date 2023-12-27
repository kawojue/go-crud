package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-crud/controllers"
)

func PostRoute(route *gin.Engine) {
	postRoute := route.Group("/api/post")
	{
		postRoute.POST("/", controllers.CreatePost)
	}
}
