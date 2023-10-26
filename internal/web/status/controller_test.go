package status

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStatus(t *testing.T) {
	// setup
	app := fiber.New()
	ctrl := &Controller{}

	// handler
	app.Get("/api/v1/status/", ctrl.Status)

	// request + response
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/status/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
