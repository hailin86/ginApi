package middleware

import (
	"ginApi/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set(common.TraceKey,uuid.New().String())
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		common.LogInfo(c,"request_info:",
			"status:",c.Writer.Status(),
			";path:",path,
			";method:",c.Request.Method,
			";query:",query,
			";ip:",c.ClientIP(),
			";errors:",c.Errors.ByType(gin.ErrorTypePrivate).String(),
			";cost:",cost,
		)
	}
}
