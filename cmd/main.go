package cmd

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-demo/internal/config"
	"gorm.io/gorm/logger"
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

	err = route.Run(os.Getenv("APP_PORT"))
	if err != nil {
		logger.Error("Can't Run Soil", zap.Error(err))
		os.Exit(0)
	}
}
