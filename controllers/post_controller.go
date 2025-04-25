package controllers

import (
	"net/http"
	"strconv"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"

	"github.com/gin-gonic/gin"
)

// DTOs for binding
type postInput struct {
	Title         string `json:"title" binding:"required"`
	Slug          string `json:"slug" binding:"required"`
	Content       string `json:"content" binding:"required"`
	Excerpt       string `json:"excerpt"`
	AuthorID      uint   `json:"author_id" binding:"required"`
	Status        string `json:"status"` // draft, publish, trash
	FeaturedImage string `json:"featured_image"`
	CategoryIDs   []uint `json:"category_ids"` // many-to-many links
	TagIDs        []uint `json:"tag_ids"`
}

// GetPosts lists all posts with related data
func GetPosts(c *gin.Context) {
	var posts []models.Post
	db := database.DB.
		Preload("Author").
		Preload("Categories").
		Preload("Tags")
	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch posts", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Posts retrieved", Data: posts})
}

// GetPostByID fetches one post by its ID
func GetPostByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid post ID"})
		return
	}
	var post models.Post
	db := database.DB.
		Preload("Author").
		Preload("Categories").
		Preload("Tags")
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Post not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Post retrieved", Data: post})
}

// CreatePost creates a new post and its associations
func CreatePost(c *gin.Context) {
	var input postInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	post := models.Post{
		Title:         input.Title,
		Slug:          input.Slug,
		Content:       input.Content,
		Excerpt:       input.Excerpt,
		AuthorID:      input.AuthorID,
		Status:        input.Status,
		FeaturedImage: input.FeaturedImage,
	}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Creation failed", Data: err.Error()})
		return
	}
	// Load associations by IDs
	var cats []models.Category
	database.DB.Find(&cats, input.CategoryIDs)
	database.DB.Model(&post).Association("Categories").Replace(cats) // many2many join

	var tags []models.Tag
	database.DB.Find(&tags, input.TagIDs)
	database.DB.Model(&post).Association("Tags").Replace(tags)

	// Return full post
	database.DB.
		Preload("Author").
		Preload("Categories").
		Preload("Tags").
		First(&post, post.ID)

	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Post created", Data: post})
}

// UpdatePost updates a post and its many-to-many links
func UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid post ID"})
		return
	}
	var input postInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Post not found"})
		return
	}
	// Update fields
	updates := map[string]interface{}{
		"Title":         input.Title,
		"Slug":          input.Slug,
		"Content":       input.Content,
		"Excerpt":       input.Excerpt,
		"Status":        input.Status,
		"FeaturedImage": input.FeaturedImage,
	}
	database.DB.Model(&post).Updates(updates) // update columns

	// Replace associations
	var cats []models.Category
	database.DB.Find(&cats, input.CategoryIDs)
	database.DB.Model(&post).Association("Categories").Replace(cats)

	var tags []models.Tag
	database.DB.Find(&tags, input.TagIDs)
	database.DB.Model(&post).Association("Tags").Replace(tags)

	// Return updated post
	database.DB.
		Preload("Author").
		Preload("Categories").
		Preload("Tags").
		First(&post, post.ID)

	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Post updated", Data: post})
}

// DeletePost deletes a post by its ID
func DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid post ID"})
		return
	}
	if err := database.DB.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Deletion failed", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Post deleted"})
}
