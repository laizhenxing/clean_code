package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	uberLimite "go.uber.org/ratelimit"
)


// fillInterval 时间长度总共可以有 cap 个请求
func RateLimit(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	limiter := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) < 1 {
			c.JSON(http.StatusOK, "Rate limit.")
			c.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}
		c.Next()
	}
}

func Rate(rate int) gin.HandlerFunc {
	limiter := uberLimite.New(rate)

	return func(c *gin.Context) {
		limiter.Take()
		c.Next()
	}
}