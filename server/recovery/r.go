package recovery

import (
	"fmt"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/server/response"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func New(ctx *fiber.Ctx) error {
	defer func(ctx *fiber.Ctx) {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fiber.ErrInternalServerError
			}
			stack := make([]byte, 4<<10)
			length := runtime.Stack(stack, false)

			fmt.Printf("%s\n %s\n", err, stack[:length])

			_ = ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get(""), nil, err))
		}
	}(ctx)
	return ctx.Next()
}
