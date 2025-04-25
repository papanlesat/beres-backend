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

type categoryInput struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
	ParentID    *uint  `json:"parent_id"`
}

type tagInput struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

// ----- Posts Handlers -----

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

// ----- Category Handlers -----

// GetCategories returns all categories (including children)
func GetCategories(c *gin.Context) {
	var categories []models.Category
	db := database.DB.Preload("Children")
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch categories", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Categories retrieved", Data: categories})
}

// GetCategoryByID returns a single category by ID
func GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid category ID"})
		return
	}
	var category models.Category
	if err := database.DB.Preload("Children").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Category not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Category retrieved", Data: category})
}

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	category := models.Category{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		ParentID:    input.ParentID,
	}
	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to create category", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Category created", Data: category})
}

// UpdateCategory updates an existing category
func UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid category ID"})
		return
	}
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Category not found"})
		return
	}
	database.DB.Model(&category).Updates(models.Category{
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		ParentID:    input.ParentID,
	})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Category updated", Data: category})
}

// DeleteCategory deletes a category by ID
func DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid category ID"})
		return
	}
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to delete category", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Category deleted"})
}

// ----- Tag Handlers -----

// GetTags returns all tags
func GetTags(c *gin.Context) {
	var tags []models.Tag
	if err := database.DB.Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch tags", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Tags retrieved", Data: tags})
}

// GetTagByID returns a single tag by ID
func GetTagByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid tag ID"})
		return
	}
	var tag models.Tag
	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Tag not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Tag retrieved", Data: tag})
}

// CreateTag creates a new tag
func CreateTag(c *gin.Context) {
	var input tagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	tag := models.Tag{Name: input.Name, Slug: input.Slug, Description: input.Description}
	if err := database.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to create tag", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Tag created", Data: tag})
}

// UpdateTag updates an existing tag
func UpdateTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid tag ID"})
		return
	}
	var input tagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var tag models.Tag
	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Tag not found"})
		return
	}
	database.DB.Model(&tag).Updates(models.Tag{Name: input.Name, Slug: input.Slug, Description: input.Description})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Tag updated", Data: tag})
}

// DeleteTag deletes a tag by ID
func DeleteTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid tag ID"})
		return
	}
	if err := database.DB.Delete(&models.Tag{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to delete tag", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Tag deleted"})
}
