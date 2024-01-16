package main

import (
	// gin-swagger middleware

	"github.com/salmantaghooni/golang-car-web-api/src/api"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/data/cache"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()
	if err := db.InitDb(cfg); err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()
	api.InitServer(cfg)

}
