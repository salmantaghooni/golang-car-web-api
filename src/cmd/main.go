package main

import (
	"log"

	"github.com/salmantaghooni/golang-car-web-api/src/api"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/data/cache"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
)

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
