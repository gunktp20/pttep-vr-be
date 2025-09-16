package permission

import (
	"fmt"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/permissions"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/server/response"
	"strconv"
	"time"

	"pttep-vr-api/pkg/utils/pagination"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service *service
}

type service struct {
	permission *permissions.Service
}

func newHandler(permission *permissions.Service) *handler {
	return &handler{
		service: &service{
			permission: permission,
		},
	}
}

func (h *handler) Get(ctx *fiber.Ctx) error {
	var result ResultGet

	var paginate pagination.Pagination
	if err := ctx.QueryParser(&paginate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	data, total, err := h.service.permission.Get(ctx.Context(), paginate.Get())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}
	result.List = data
	result.Total = int(total)
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (h *handler) GetByID(ctx *fiber.Ctx) error {

	var id uint
	if ctx.Params("permission_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("permission_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	data, err := h.service.permission.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), data, nil))
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	var request Request

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	data, err := h.service.permission.Create(ctx.Context(), models.Permission{
		Key:         request.Key,
		Name:        request.Name,
		IsActive:    true,
		IsRemove:    false,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), data, nil))
}

func (h *handler) UpdateIsActive(ctx *fiber.Ctx) error {
	var request Request

	var id uint
	if ctx.Params("permission_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("permission_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	err := h.service.permission.UpdateIsActive(ctx.Context(), models.Permission{
		ID:       id,
		IsActive: request.IsActive,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (h *handler) Delete(ctx *fiber.Ctx) error {

	var id uint
	if ctx.Params("permission_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("permission_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	err := h.service.permission.Delete(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}
