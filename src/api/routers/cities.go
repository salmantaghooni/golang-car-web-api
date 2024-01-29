package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salmantaghooni/golang-car-web-api/src/api/handlers"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
)

func City(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCityHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
