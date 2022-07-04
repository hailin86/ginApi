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
		//c.Set("trace_id",uuid.New().String())
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method
		c.Next()
		cost := time.Since(start)

		if method != "OPTIONS" {
			common.LogInfo(c,"request_info:",
				"status:",c.Writer.Status(),
				";path:",path,
				";method:",method,
				";query:",query,
				";ip:",c.ClientIP(),
				";errors:",c.Errors.ByType(gin.ErrorTypePrivate).String(),
				";cost:",cost,
			)
		}

	}
}
