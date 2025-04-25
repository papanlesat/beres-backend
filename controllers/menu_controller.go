package controllers

import (
	"net/http"
	"strconv"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"

	"github.com/gin-gonic/gin"
)

// DTOs
type menuInput struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type menuItemInput struct {
	MenuID   uint   `json:"menu_id" binding:"required"`
	ParentID *uint  `json:"parent_id"`
	Title    string `json:"title" binding:"required"`
	URL      string `json:"url" binding:"required"`
	Order    int    `json:"order"`
	Class    string `json:"class"`
	Target   string `json:"target"`
}

// ----- Menu Handlers -----

// GetMenus lists all menus with nested items
func GetMenus(c *gin.Context) {
	var menus []models.Menu
	db := database.DB.
		Preload("Items").
		Preload("Items.Children")
	if err := db.Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch menus", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menus retrieved", Data: menus})
}

// GetMenuByID returns a single menu by ID
func GetMenuByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid menu ID"})
		return
	}
	var menu models.Menu
	db := database.DB.
		Preload("Items").
		Preload("Items.Children")
	if err := db.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Menu not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu retrieved", Data: menu})
}

// CreateMenu creates a new menu
func CreateMenu(c *gin.Context) {
	var input menuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	menu := models.Menu{Name: input.Name, Location: input.Location}
	if err := database.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to create menu", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Menu created", Data: menu})
}

// UpdateMenu updates an existing menu
func UpdateMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid menu ID"})
		return
	}
	var input menuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var menu models.Menu
	if err := database.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Menu not found"})
		return
	}
	database.DB.Model(&menu).Updates(models.Menu{Name: input.Name, Location: input.Location})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu updated", Data: menu})
}

// DeleteMenu deletes a menu by ID
func DeleteMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid menu ID"})
		return
	}
	if err := database.DB.Delete(&models.Menu{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to delete menu", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu deleted"})
}

// ----- MenuItem Handlers -----

// GetMenuItems lists items for a specific menu, including nested children
func GetMenuItems(c *gin.Context) {
	menuID, err := strconv.ParseUint(c.Param("menu_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid menu ID"})
		return
	}
	var items []models.MenuItem
	db := database.DB.Where("menu_id = ?", menuID).
		Preload("Children")
	if err := db.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to fetch menu items", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu items retrieved", Data: items})
}

// GetMenuItemByID returns one menu item by ID
func GetMenuItemByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid item ID"})
		return
	}
	var item models.MenuItem
	if err := database.DB.Preload("Children").First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Menu item not found"})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu item retrieved", Data: item})
}

// CreateMenuItem creates a new menu item
func CreateMenuItem(c *gin.Context) {
	var input menuItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	item := models.MenuItem{
		MenuID:   input.MenuID,
		ParentID: input.ParentID,
		Title:    input.Title,
		URL:      input.URL,
		Order:    input.Order,
		Class:    input.Class,
		Target:   input.Target,
	}
	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to create menu item", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "Menu item created", Data: item})
}

// UpdateMenuItem updates an existing menu item
func UpdateMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid item ID"})
		return
	}
	var input menuItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var item models.MenuItem
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.Response{Code: http.StatusNotFound, Message: "Menu item not found"})
		return
	}
	database.DB.Model(&item).Updates(models.MenuItem{
		MenuID:   input.MenuID,
		ParentID: input.ParentID,
		Title:    input.Title,
		URL:      input.URL,
		Order:    input.Order,
		Class:    input.Class,
		Target:   input.Target,
	})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu item updated", Data: item})
}

// DeleteMenuItem deletes a menu item by ID
func DeleteMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid item ID"})
		return
	}
	if err := database.DB.Delete(&models.MenuItem{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Failed to delete menu item", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Menu item deleted"})
}
