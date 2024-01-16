package main

import (
	"log"
	// gin-swagger middleware

	"github.com/salmantaghooni/golang-car-web-api/src/api"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/data/cache"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()
	if err := db.InitDb(cfg); err != nil {
		panic(err)
	}
	defer db.CloseDb()
	api.InitServer(cfg)

}
