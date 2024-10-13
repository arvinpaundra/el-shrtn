package link

import "github.com/gofiber/fiber/v2"

func (h *Handler) RouterPublic(app fiber.Router) {
	app.Get("/:code", h.HandlerAccessLink)
}

func (h *Handler) RouterV1(app fiber.Router) {
	app.Post("", h.HandlerCreateLink)
}
