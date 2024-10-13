package link

import (
	"context"

	"github.com/arvinpaundra/el-shrtn/internal/app/link"
	"github.com/arvinpaundra/el-shrtn/internal/dto/request"
	"github.com/arvinpaundra/el-shrtn/internal/dto/response"
	"github.com/arvinpaundra/el-shrtn/internal/factory"
	"github.com/arvinpaundra/el-shrtn/pkg/format"
	"github.com/arvinpaundra/el-shrtn/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	CreateLink(ctx context.Context, payload request.CreateLink) (response.CreatedLink, error)
	AccessLink(ctx context.Context, code string) (string, error)
}

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: link.NewService(f),
	}
}

func (h *Handler) HandlerCreateLink(c *fiber.Ctx) error {
	var payload request.CreateLink

	_ = c.BodyParser(&payload)

	errorValidation := validator.Validate(payload, validator.JSON)
	if errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest("invalid request body", errorValidation))
	}

	res, err := h.service.CreateLink(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(format.InternalServerError(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(format.SuccessCreated("success shortening url", res))
}

func (h *Handler) HandlerAccessLink(c *fiber.Ctx) error {
	code := c.Params("code")

	res, err := h.service.AccessLink(c.Context(), code)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Redirect(res)
}
