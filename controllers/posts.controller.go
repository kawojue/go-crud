package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kawojue/go-crud/DB"
	"github.com/kawojue/go-crud/models"
)

func CreatePost(ctx *gin.Context) {
	post := models.Post{Title: "New Post", Body: "Hello, World!"}

	result := db.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post created.",
	})
}
