package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/config"
)

func Cors(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Cors.AllowOrigins)
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Max-Age", "21600")
		c.Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

	}
}
