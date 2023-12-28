package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kawojue/go-crud/DB"
	"github.com/kawojue/go-crud/models"
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
