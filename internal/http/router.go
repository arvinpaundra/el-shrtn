package http

import (
	"github.com/arvinpaundra/el-shrtn/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/arvinpaundra/el-shrtn/internal/factory"
	"github.com/arvinpaundra/el-shrtn/internal/http/link"
)

func NewHttpRouter(app *fiber.App, f *factory.Factory) {
	// setup app middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-type, Authorization",
	}))

	app.Use(recover.New())

	// index route
	app.Get("/", index)

	v1 := app.Group("/api/v1")

	// link routes
	link.NewHandler(f).RouterPublic(app.Group(""))
	link.NewHandler(f).RouterV1(v1.Group("/links"))
}

func index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name":    "el-shrtn",
		"version": util.LoadVersion(),
	})
}
