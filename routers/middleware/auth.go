package middleware

import (
	"net/http"
	"strings"
	"time"

	"beres/helpers"
	"beres/infra/database"
	"beres/models"

	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.Response{Code: http.StatusUnauthorized, Message: "Missing token"})
			return
		}
		raw := strings.TrimPrefix(auth, "Bearer ")
		hash := helpers.HashToken(raw)

		var token models.PersonalAccessToken
		if err := database.DB.Where("token_hash = ?", hash).First(&token).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.Response{Code: http.StatusUnauthorized, Message: "Invalid token"})
			return
		}
		if token.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.Response{Code: http.StatusUnauthorized, Message: "Token expired"})
			return
		}
		// update last used
		database.DB.Model(&token).Update("last_used_at", time.Now())

		var user models.User
		database.DB.First(&user, token.UserID)
		c.Set("current_user", user)
		c.Set("token_hash", hash)
		c.Next()
	}
}
