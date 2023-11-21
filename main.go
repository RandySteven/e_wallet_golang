package main

import (
	"e_wallet/apps"
	"e_wallet/infrastructure/persistences"
	"e_wallet/interfaces"
	"e_wallet/interfaces/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return
	}

	// appEnv := os.Getenv("APP_ENV")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	services, err := persistences.NewRepository(dbHost, dbName, dbUser, dbPass, dbPort)
	if err != nil {
		log.Println(err)
		return
	}

	defer services.Close()

	err = services.Automigrate()
	if err != nil {
		log.Println(err)
		return
	}

	gin.SetMode("debug")
	validate := validator.New()

	users := interfaces.NewUser(apps.NewUserApp(*services), validate)
	transfers := interfaces.NewTransferTransaction(apps.NewTransferApp(*services))
	topups := interfaces.NewTopup(apps.NewTopup(*services))
	r := gin.New()
	v1 := r.Group("/v1")
	routers.UserRouter(v1, users)
	routers.TransferRouter(v1, transfers)
	routers.TopupRouter(v1, topups)

	appPort := os.Getenv("APP_PORT")
	log.Fatal(r.Run(":" + appPort))
}
