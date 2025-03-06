package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/cvele/reptask/internal/config"
	"github.com/cvele/reptask/internal/db"
	"github.com/cvele/reptask/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadConfig()
	db.InitDB(config.Config.SQLiteDB)
	defer db.CloseDB()

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	routes.SetupRoutes(app)

	go func() {
		if err := app.Listen(":" + config.Config.Port); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped successfully.")
}
