package controllers

import (
	"errors"
	"net/http"
	"product/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validation
type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Err message
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown Error"
}

func GetPosts(c *gin.Context) {
	// Get Data
	var posts []models.Post
	models.DB.Find(&posts)

	// return
	c.JSON(200, gin.H{
		"success": true,
		"message": "Data berhasil ditambahkan",
		"data":    posts,
	})
}

func StorePost(c *gin.Context) {
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	// create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	models.DB.Create(&post)

	c.JSON(201, gin.H{
		"success": true,
		"message": "Berhasil ditambahkan",
		"data":    post,
	})
}

func DetailPost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ada"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail product dari Id" + c.Param("id"),
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ada"})
		return
	}
	// Form validation
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	// update
	models.DB.Model(&post).Updates(input)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//delete post
	models.DB.Delete(&post)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}
