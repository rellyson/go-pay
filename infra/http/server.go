package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func CreateHttpServer() *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Use(requestid.New(requestid.Config{
		Header: "X-Request-ID",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	SetRoutes(app)

	return app
}
