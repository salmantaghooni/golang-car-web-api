package api

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/api/routers"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(":5005")

}
