package handlers

import (
	"test-backend-pari/items"

	"github.com/gofiber/fiber/v2"
)

func MapItem(routes fiber.Router, h items.Handler) {
	routes.Get("/items", h.Index)
	routes.Get("/items/:id", h.Show)
	routes.Post("/items", h.Store)
	routes.Put("/items/:id", h.Update)
	routes.Delete("/items/:id", h.Delete)
}