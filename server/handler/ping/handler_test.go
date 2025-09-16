package ping

import (
	"net/http/httptest"
	"testing"

	"pttep-vr-api/pkg/services/ping"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Ping(t *testing.T) {
	service := &ping.Service{}
	handler := newHandler(service)

	app := fiber.New()
	app.Get("/ping", handler.Ping)

	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}