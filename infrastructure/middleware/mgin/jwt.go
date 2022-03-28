package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"cc/infrastructure/utils/identity"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
			c.Abort()
			return
		}

		// 验证 token 的有效性
		if err := identity.ValidTokenError(auth, []byte(viper.GetString("app.jwt_secret"))); err != nil {
			//if ve, ok := err.(*errors.Error); ok {
			//	// token 过期
			//	if ve.ErrCode == ccErr.ErrExpiredToken {
			//		// 续签？
			//	}
			//}
			c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
			c.Abort()
			return
		}

		// 解析 token
		claims, err := identity.ParseToken(auth, []byte(viper.GetString("app.jwt_secret")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "parse token error"})
			c.Abort()
			return
		}

		if claims.UserID <= 0 || !claims.Authrized {
			c.JSON(http.StatusNotFound, gin.H{"msg": "user not found"})
			c.Abort()
			return
		}
		c.Next()
	}
}
