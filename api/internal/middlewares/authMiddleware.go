package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/scallyt/neko/internal/utils"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := utils.ValidateClerkJWT(tokenStr)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.Next()
}
