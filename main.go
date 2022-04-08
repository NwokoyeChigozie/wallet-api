package main

import (
	"log"

	"github.com/NwokoyeChigozie/quik_task/internal/config"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository/storage/mysql"
	"github.com/NwokoyeChigozie/quik_task/pkg/repository/storage/redis"
	"github.com/NwokoyeChigozie/quik_task/pkg/router"
	"github.com/go-playground/validator/v10"
)

func init() {
	config.Setup()
	mysql.ConnectToDB()
	redis.SetupRedis()
}

func main() {
	//Load config
	getConfig := config.GetConfig()
	validatorRef := validator.New()
	r := router.Setup(validatorRef)

	log.Printf("Server is starting at 127.0.0.1:%s", getConfig.Server.Port)
	log.Fatal(r.Run(":" + getConfig.Server.Port))

}
