package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kawojue/go-crud/DB"
	"github.com/kawojue/go-crud/models"
	"gorm.io/gorm"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	err := ctx.Bind(&body)
	if err != nil {
		log.Fatal("Error binding data.", err)
	}

	if body.Title == "" && body.Body == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "All fields cannot be blank",
		})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := db.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"post":   post,
	})
}

func ReadPosts(ctx *gin.Context) {
	var posts []*models.Post
	db.DB.Find(&posts)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"posts":  posts,
	})
}

func ReadPost(ctx *gin.Context) {
	var post *models.Post
	id := ctx.Param("id")

	err := db.DB.Where("id = ?", id).First(&post).Error

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Post not found.",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Something went wrong.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"post":   post,
	})
}
