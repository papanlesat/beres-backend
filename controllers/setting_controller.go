package controllers

import (
	"net/http"
	"strconv"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"

	"github.com/gin-gonic/gin"
)

// DTO for binding Setting
type settingInput struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// GetSettings returns all settings
func GetSettings(c *gin.Context) {
	var settings []models.Setting
	if err := database.DB.Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch settings", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Settings retrieved", Data: settings})
}

// GetSettingByID returns a setting by its ID
func GetSettingByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid setting ID"})
		return
	}
	var setting models.Setting
	if err := database.DB.First(&setting, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Setting not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Setting retrieved", Data: setting})
}

// CreateSetting creates a new setting
func CreateSetting(c *gin.Context) {
	var input settingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	setting := models.Setting{Key: input.Key, Value: input.Value}
	if err := database.DB.Create(&setting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to create setting", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Setting created", Data: setting})
}

// UpdateSetting updates an existing setting by ID
func UpdateSetting(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid setting ID"})
		return
	}
	var input settingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var setting models.Setting
	if err := database.DB.First(&setting, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Setting not found"})
		return
	}
	database.DB.Model(&setting).Updates(models.Setting{Key: input.Key, Value: input.Value})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Setting updated", Data: setting})
}

// DeleteSetting deletes a setting by ID
func DeleteSetting(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid setting ID"})
		return
	}
	if err := database.DB.Delete(&models.Setting{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to delete setting", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Setting deleted"})
}
