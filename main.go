package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/kawojue/go-crud/DB"
	"github.com/kawojue/go-crud/routes"
	"github.com/kawojue/go-crud/utils"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	utils.LoadEnv(nil)
	db.ConnectDB()
}

func main() {
	router := gin.Default()
	PORT := utils.GetEnv("PORT", "2002")

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
	})

	routes.PostRoute(router)

	router.Run(fmt.Sprintf(":%s", PORT))
}
