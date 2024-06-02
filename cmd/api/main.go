package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sebaperi/go-skeleton/cmd/api/handlers/player"
	"github.com/sebaperi/go-skeleton/internal/repositories/mongo"
	playerMongo "github.com/sebaperi/go-skeleton/internal/repositories/mongo/player"
	playerService "github.com/sebaperi/go-skeleton/internal/services/player"
	"log"
	"os"
)

func main() {
	// it is a good practice to return log.Fatal only on the main function
	// because log.Fatal will be thrown only on startup
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	playerRepo := playerMongo.Repository{
		Client: client,
	}

	playerSrv := playerService.Service{
		Repo: playerRepo,
	}

	playerHandler := player.Handler{
		PlayerService: playerSrv,
	}

	ginEngine.POST("players", playerHandler.CreatePlayer)

	log.Fatalln(ginEngine.Run(":8001"))
}
