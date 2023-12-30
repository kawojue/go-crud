package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

	title := strings.TrimSpace(body.Title)
	body_ := strings.TrimSpace(body.Body)

	if title == "" && body_ == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "All fields cannot be blank",
		})
		return
	}

	post := models.Post{
		Title: title,
		Body:  body_,
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
	var posts []models.Post
	db.DB.Find(&posts)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"posts":  posts,
	})
}

func ReadPost(ctx *gin.Context) {
	var post models.Post
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

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var (
		body struct {
			Title string
			Body  string
		}
		err  error
		post models.Post
	)

	if err = ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body.",
		})
		return
	}

	title := strings.TrimSpace(body.Title)
	body_ := strings.TrimSpace(body.Body)

	if title == "" && body_ == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "All fields cannot be empty.",
		})
		return
	}

	if err = db.DB.First(&post, id).Error; err == gorm.ErrRecordNotFound || err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Post not found.",
		})
		return
	}

	if err = db.DB.Model(&post).Updates(models.Post{Title: title, Body: body_}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Error updating post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"post":   post,
	})

}

func DeletePost(ctx *gin.Context) {
	var (
		id   string = ctx.Param("id")
		post models.Post
		err  error
	)

	if err = db.DB.First(&post, id).Error; err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Post not found.",
		})
		return
	}

	if err = db.DB.Where("id = ?", id).Delete(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error deleting post.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Post with ID: %v has been deleted", id),
	})
}
