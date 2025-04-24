package controllers

import (
	"net/http"
	"strconv"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"
	"beres/repository"

	"github.com/gin-gonic/gin"
)

// GetSectionData returns all sections
func GetSectionData(ctx *gin.Context) {
	var sections []models.Section
	if err := repository.Get(&sections); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch sections",
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Message: "Sections retrieved",
		Data:    sections,
	})
}

// GetSectionByID returns a single section by its ID
func GetSectionByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid section ID",
		})
		return
	}

	var section models.Section
	if err := database.DB.First(&section, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Response{
			Code:    http.StatusNotFound,
			Message: "Section not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Message: "Section retrieved",
		Data:    section,
	})
}

// CreateSection creates a new section
func CreateSection(ctx *gin.Context) {
	var section models.Section
	if err := ctx.ShouldBindJSON(&section); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Data:    err.Error(),
		})
		return
	}

	if err := repository.Save(&section); err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create section",
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, helpers.Response{
		Code:    http.StatusCreated,
		Message: "Section created",
		Data:    section,
	})
}

// UpdateSection updates an existing section by ID
func UpdateSection(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid section ID",
		})
		return
	}

	var input models.Section
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Data:    err.Error(),
		})
		return
	}

	var section models.Section
	if err := database.DB.First(&section, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Response{
			Code:    http.StatusNotFound,
			Message: "Section not found",
		})
		return
	}

	// Update fields
	section.Name = input.Name
	section.SectionType = input.SectionType
	section.DisplayOrder = input.DisplayOrder
	section.IsActive = input.IsActive
	section.Details = input.Details

	if err := database.DB.Save(&section).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update section",
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Message: "Section updated",
		Data:    section,
	})
}

// DeleteSection removes a section by ID
func DeleteSection(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid section ID",
		})
		return
	}

	var section models.Section
	if err := database.DB.First(&section, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Response{
			Code:    http.StatusNotFound,
			Message: "Section not found",
		})
		return
	}

	if err := database.DB.Delete(&section).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete section",
			Data:    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response{
		Code:    http.StatusOK,
		Message: "Section deleted",
	})
}
