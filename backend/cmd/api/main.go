package main

import (
	"fmt"

	"github.com/ferixthecat/bookingapp/config"
	"github.com/ferixthecat/bookingapp/db"
	"github.com/ferixthecat/bookingapp/handlers"
	"github.com/ferixthecat/bookingapp/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
