package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kumarvikramshahi/ideal_tutor_backend_golang/utils/configs"
)

func main(){
	app:=fiber.New()
	
	// DatabaseConnection
	configs.DatabaseConnection()
	log.Println("hello main")
	// for logging incoming requests
	app.Use(logger.New())
	
	// routes.PractisePaper(app)
	
	app.Listen("127.0.0.1:3000")
}