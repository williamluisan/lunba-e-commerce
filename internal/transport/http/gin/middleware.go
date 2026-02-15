package gin

import (
	"net/http"
	"strings"

	"lunba-e-commerce/internal/transport/http/gin/handler"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	messages "lunba-e-commerce/internal/domain/messages"
)

func AuthJWTValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Error: &handler.APIError{
					Code: messages.MissingToken.Code,
					Message: messages.MissingToken.Message,
				},
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Error: &handler.APIError{
					Code: messages.InvalidTokenFormat.Code,
					Message: messages.InvalidTokenFormat.Message,
				},
			})
			return
		}

		tokenStr := parts[1]

		/* validate JWT */
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
					Success: false,
					Error: &handler.APIError{
						Code: messages.UnexpectedSigningMethod.Code,
						Message: messages.UnexpectedSigningMethod.Message,
					},
				})
				return nil, nil
			}
			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Error: &handler.APIError{
					Code: messages.UnexpectedSigningMethod.Code,
					Message: messages.UnexpectedSigningMethod.Message,
				},
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Error: &handler.APIError{
					Code: messages.InvalidTokenClaims.Code,
					Message: messages.InvalidTokenClaims.Message,
				},
			})
			return
		}
		/* // */

		userPubID, ok := claims["sub"].(string)
		if !ok || userPubID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Error: &handler.APIError{
					Code: messages.InvalidTokenClaims.Code,
					Message: messages.InvalidTokenClaims.Message,
				},
			})
			return
		}

		// attach user public id to context
		c.Set("authorization", "Bearer " + tokenStr)
		c.Set("user_public_id", userPubID)
		c.Next()
	}
}

func EmptyBodyRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			if c.Request.ContentLength == 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, handler.APIResponse{
					Success: false,
					Error: &handler.APIError{
						Code: "empty_body",
						Message: "Empty Body",
						Details: "Empty Body",
					},
				})
				return
			}
		}
	}
}