package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type healthCheckController struct{}

type HealthCheckController interface {
	Check(c *fiber.Ctx) error
}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

func (*healthCheckController) Check(c *fiber.Ctx) error {
	r := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "Ok",
	}

	return c.Status(r.Status).JSON(r)
}
