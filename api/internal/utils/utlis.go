package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

func ValidateClerkJWT(tokenStr string) (*jwt.Token, error) {

	// TODO: Don't hardcode secrets
	const clerkJWKURL string = "secret"

	keySet, err := jwk.Fetch(context.Background(), clerkJWKURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWK: %v", err)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"]
		if kid == nil {
			return nil, fmt.Errorf("missing kid in header")
		}
		key, ok := keySet.LookupKeyID(kid.(string))
		if !ok {
			return nil, fmt.Errorf("key not found")
		}
		var rawKey interface{}
		if err := key.Raw(&rawKey); err != nil {
			return nil, fmt.Errorf("failed to get raw key: %v", err)
		}
		return rawKey, nil
	})

	return token, err
}

func GetJWTUserID(tokenStr string) (*jwt.Token, string, error) {
	const clerkJWKURL string = "secret"

	keySet, err := jwk.Fetch(context.Background(), clerkJWKURL)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch JWK: %v", err)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"]
		if kid == nil {
			return nil, fmt.Errorf("missing kid in header")
		}
		key, ok := keySet.LookupKeyID(kid.(string))
		if !ok {
			return nil, fmt.Errorf("key not found")
		}
		var rawKey interface{}
		if err := key.Raw(&rawKey); err != nil {
			return nil, fmt.Errorf("failed to get raw key: %v", err)
		}
		return rawKey, nil
	})

	if err != nil {
		return nil, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, "", fmt.Errorf("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, "", fmt.Errorf("sub claim not found or invalid")
	}

	return token, sub, nil
}

func GetUserIdByJWT(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return "", fmt.Errorf("missing token")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return "", fmt.Errorf("invalid token")
	}
	sub, ok := claims["sub"].(string)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
		return "", fmt.Errorf("invalid token payload")
	}

	return sub, nil
}
