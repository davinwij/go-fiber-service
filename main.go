package main

import (
	"go-fiber-tutor/configs/database"
	"go-fiber-tutor/pkg/routes"
	"go-fiber-tutor/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	database.Init()

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
