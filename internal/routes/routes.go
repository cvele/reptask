package routes

import (
	"github.com/cvele/reptask/internal/pack"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/packs", pack.GetPacksHandler)
	api.Post("/packs", pack.CreatePackHandler)
	api.Get("/packs/:id", pack.GetPackByIDHandler)
	api.Put("/packs/:id", pack.UpdatePackHandler)
	api.Delete("/packs/:id", pack.DeletePackHandler)

	api.Get("/calculate", pack.CalculatePacksHandler)
}
