package version

import (
	"net/http"
	"pttep-vr-api/pkg/config"

	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	config *config.Config
}

func newHandler(config *config.Config) *handler {
	return &handler{
		config: config,
	}
}
func (o *handler) GetVersion(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"name":     o.config.App.Name,
		"version":  o.config.App.Version,
		"state":    o.config.App.State,
		"timezone": o.config.App.Timezone,
		"time":     time.Now().Format(time.RFC3339),
	})
}
