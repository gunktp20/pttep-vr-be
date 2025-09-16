package ping

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"pttep-vr-api/pkg/services/ping"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/server/response"
	"time"
)

type handler struct {
	service *service
}

type service struct {
	ping *ping.Service
}

func newHandler(ping *ping.Service) *handler {
	return &handler{
		service: &service{
			ping: ping,
		},
	}
}

func (h *handler) Ping(ctx *fiber.Ctx) error {
	msg, err := h.service.ping.Ping()
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("ping_error"), nil, err))
	}

	result := Result{Message: msg}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, err))
}

func (h *handler) PingPanic(ctx *fiber.Ctx) error {

	msg, err := h.service.ping.Ping()
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("ping_error"), nil, err))
	}

	panic("ping panic")

	result := Result{Message: msg}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, err))
}

func (h *handler) PingSleep(ctx *fiber.Ctx) error {

	msg, err := h.service.ping.Ping()
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("ping_error"), nil, err))
	}

	fmt.Println("Ping")
	time.Sleep(20 * time.Second)

	result := Result{Message: msg}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, err))
}
