package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/voideus/mini-ecommerce/handler"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authoriztion header found",
			})
			return
		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := handler.ValidateToken(tokenString); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					fmt.Println("during authorization", claims["userID"])
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}

	}
}
