package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rellyson/go-pay/application/controllers"
)

var (
	hcController controllers.HealthCheckController = controllers.NewHealthCheckController()
)

func SetRoutes(app *fiber.App) {
	app.Get("/healthcheck", hcController.Check)
}
