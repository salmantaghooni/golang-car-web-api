package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTesthHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserByID)
	r.GET("/user/username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.AccountByID)
	r.POST("/user/store", h.AddUser)
}
