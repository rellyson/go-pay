package controllers_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rellyson/go-pay/application/controllers"
	"github.com/stretchr/testify/assert"
)

func TestCheckResponse(t *testing.T) {
	mockedApp := fiber.New()
	SUTController := controllers.NewHealthCheckController()

	mockedApp.Get("/healthcheck", SUTController.Check)

	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res, _ := mockedApp.Test(req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Check the response body is what we expect.
	expectedBody := `{"status": 200, "message": "Ok"}`
	resBody, _ := io.ReadAll(res.Body)

	assert.JSONEq(t, expectedBody, string(resBody))
}
