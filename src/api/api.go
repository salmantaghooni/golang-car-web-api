package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/salmantaghooni/golang-car-web-api/src/api/middlewares"
	"github.com/salmantaghooni/golang-car-web-api/src/api/routers"
	"github.com/salmantaghooni/golang-car-web-api/src/api/validations"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/docs"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterMiddlewares(r, cfg)
	RegisterValidator()
	RegisterRoutes(r)
	RegisterSwagger(r, cfg)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterMiddlewares(r *gin.Engine, cfg *config.Config) {
	r.Use(middlewares.DefaultStracturedLogger(cfg))
	r.Use(middlewares.LimitByRequest())
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery())
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
}
func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
