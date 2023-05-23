package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/kumarvikramshahi/idealTutor_backend/routes"
	"github.com/kumarvikramshahi/idealTutor_backend/utils/configs"
)

func main(){
	app:=fiber.New()
	
	// DatabaseConnections
	configs.DatabaseConnection()

	// for logging incoming requests
	app.Use(logger.New())
	
	routes.Profiles(app)
	
	app.Listen("127.0.0.1:3000")
}