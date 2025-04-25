package controllers

import (
	"net/http"
	"strconv"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"

	"github.com/gin-gonic/gin"
)

// DTO for binding Widget
type widgetInput struct {
	Type      string `json:"type" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Position  string `json:"position" binding:"required"`
	SortOrder int    `json:"sort_order"`
}

// GetWidgets returns all widgets
func GetWidgets(c *gin.Context) {
	var widgets []models.Widget
	if err := database.DB.Find(&widgets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch widgets", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Widgets retrieved", Data: widgets})
}

// GetWidgetByID returns a widget by its ID
func GetWidgetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid widget ID"})
		return
	}
	var widget models.Widget
	if err := database.DB.First(&widget, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Widget not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Widget retrieved", Data: widget})
}

// CreateWidget creates a new widget
func CreateWidget(c *gin.Context) {
	var input widgetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	widget := models.Widget{
		Type:      input.Type,
		Title:     input.Title,
		Content:   input.Content,
		Position:  input.Position,
		SortOrder: input.SortOrder,
	}
	if err := database.DB.Create(&widget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to create widget", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Widget created", Data: widget})
}

// UpdateWidget updates an existing widget by ID
func UpdateWidget(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid widget ID"})
		return
	}
	var input widgetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var widget models.Widget
	if err := database.DB.First(&widget, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Widget not found"})
		return
	}
	database.DB.Model(&widget).Updates(models.Widget{
		Type:      input.Type,
		Title:     input.Title,
		Content:   input.Content,
		Position:  input.Position,
		SortOrder: input.SortOrder,
	})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Widget updated", Data: widget})
}

// DeleteWidget deletes a widget by ID
func DeleteWidget(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid widget ID"})
		return
	}
	if err := database.DB.Delete(&models.Widget{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to delete widget", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Widget deleted"})
}
