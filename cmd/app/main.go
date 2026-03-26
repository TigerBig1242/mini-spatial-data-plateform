package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/tigerbig/spatial-data-plateform/internal/config"
	"github.com/tigerbig/spatial-data-plateform/internal/domain/entities"
	"github.com/tigerbig/spatial-data-plateform/internal/infrastructure/database"
)

func main() {
	databaseConfig := config.LoadConfig()

	db, err := database.ConnectDatabase(databaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entities.Test{}, &entities.SpatialFeature{})
	if err != nil {
		log.Fatalf("Database Migration failed: %v", err)
	}

	app := fiber.New()

	app.Get("/hello", func(c fiber.Ctx) error {
		if err != nil {
			return c.Status(500).SendString("Internal server error")
		}
		return c.SendString("Hello World at main.go")
	})
	app.Listen(":8085")
}
