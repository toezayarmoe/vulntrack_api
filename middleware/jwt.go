package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/toezayarmoe/vulntrack_api/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(401, gin.H{"error": "Missing Auth Token"})
			ctx.Abort()
			return
		}

		tokenStr := strings.Replace(auth, "Bearer ", "", 1)
		token, err := utils.ParseToken(tokenStr)

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "Invalid Token"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("is_admin", claims["is_admin"])
	}
}
