package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/api/handlers"
	"github.com/salmantaghooni/golang-car-web-api/src/api/middlewares"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
}
