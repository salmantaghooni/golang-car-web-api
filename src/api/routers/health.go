package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.POST("/:id", handler.HealthPostById)
}
