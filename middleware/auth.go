package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mcharolabs/go-crud/helpers"
	"github.com/mcharolabs/go-crud/response"
)

// Authenticate validates JWT and sets userClaims in context
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Send[any](c.Writer, http.StatusUnauthorized, false, "Authorization header required", nil, c.FullPath())
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helpers.ValidateToken(token)
		if err != nil {
			response.Send[any](c.Writer, http.StatusUnauthorized, false, "Invalid or expired token", nil, c.FullPath())
			c.Abort()
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

// AuthorizeRole checks if user has one of allowed roles
func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsRaw, exists := c.Get("userClaims")
		if !exists {
			response.Send[any](c.Writer, http.StatusForbidden, false, "User claims not found", nil, c.FullPath())
			c.Abort()
			return
		}

		claims, ok := claimsRaw.(helpers.SignedDetails)
		if !ok {
			response.Send[any](c.Writer, http.StatusForbidden, false, "Invalid user claims", nil, c.FullPath())
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if claims.UserType == role {
				c.Next()
				return
			}
		}

		response.Send[any](c.Writer, http.StatusForbidden, false, "Access denied: insufficient permissions", nil, c.FullPath())
		c.Abort()
	}
}
