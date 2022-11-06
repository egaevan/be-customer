package cmd

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"golang-demo/internal/config"
	"log"
	"net"
	"os"
)

var (
	cfg = config.Server()
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", config.Server().Port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	route := gin.New()
	route.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health", "/metrics"),
		gin.Recovery(),
	)

	route.Use(cors.AllowAll())

	err = route.Run(os.Getenv("APP_PORT"))
	if err != nil {
		log.Println()
		os.Exit(0)
	}

	log.Println("Serving gRPC on " + config.Server().Port)
	log.Fatal(lis)
}
