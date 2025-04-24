package controllers

import (
	"net/http"
	"time"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Name: input.Name, Email: input.Email, PasswordHash: string(hash)}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Registration failed", Data: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, helpers.Response{Code: http.StatusCreated, Message: "User registered", Data: user})
}

func Login(c *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
		TokenName string `json:"token_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.Response{Code: http.StatusBadRequest, Message: "Invalid input", Data: err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, helpers.Response{Code: http.StatusUnauthorized, Message: "Invalid credentials"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, helpers.Response{Code: http.StatusUnauthorized, Message: "Invalid credentials"})
		return
	}
	// Generate a random token (e.g. 40-char string)
	rawToken := helpers.GenerateRandomString(40)
	hash := helpers.HashToken(rawToken)
	now := time.Now()
	token := models.PersonalAccessToken{
		UserID:     user.ID,
		Name:       input.TokenName,
		TokenHash:  hash,
		Abilities:  "*",
		LastUsedAt: &now,
		ExpiresAt:  ptrTime(time.Now().Add(24 * time.Hour)),
	}
	if err := database.DB.Create(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.Response{Code: http.StatusInternalServerError, Message: "Token creation failed", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Login successful", Data: gin.H{"token": rawToken}})
}

func Logout(c *gin.Context) {
	t := c.GetString("token_hash") // set by middleware
	database.DB.Where("token_hash = ?", t).Delete(&models.PersonalAccessToken{})
	c.JSON(http.StatusOK, helpers.Response{Code: http.StatusOK, Message: "Logged out"})
}

func ptrTime(t time.Time) *time.Time { return &t }
