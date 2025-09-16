package role

import (
	"fmt"
	"pttep-vr-api/pkg/services/roles"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/server/response"
	"strconv"
	"time"

	models "pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service *service
}

type service struct {
	role *roles.Service
}

func newHandler(role *roles.Service) *handler {
	return &handler{
		service: &service{
			role: role,
		},
	}
}

func (o *handler) Get(ctx *fiber.Ctx) error {
	var result ResultGet

	var paginate pagination.Pagination
	if err := ctx.QueryParser(&paginate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	data, total, err := o.service.role.Get(ctx.Context(), paginate.Get())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}
	result.List = data
	result.Total = int(total)
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) GetAndPermission(ctx *fiber.Ctx) error {
	var result ResultGetAndPermission

	var paginate pagination.Pagination
	if err := ctx.QueryParser(&paginate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	data, total, err := o.service.role.GetAndPermission(ctx.Context(), paginate.Get())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}
	result.List = data
	result.Total = int(total)
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) GetByID(ctx *fiber.Ctx) error {

	var id uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	data, err := o.service.role.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), data, nil))
}

func (o *handler) GetByIDAndPermission(ctx *fiber.Ctx) error {

	var id uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	data, err := o.service.role.GetByIDAndPermission(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), data, nil))
}

func (o *handler) Create(ctx *fiber.Ctx) error {
	var request Request

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	data, err := o.service.role.Create(ctx.Context(), roles.Model{
		Role: models.Role{
			Key:         request.Key,
			Name:        request.Name,
			Description: request.Description,
			IsActive:    request.IsActive,
			IsRemove:    false,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		},
		Permissions: nil,
	}, request.PermissionIds)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), data, nil))
}

func (o *handler) Update(ctx *fiber.Ctx) error {
	var request Request

	var id uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
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

	var permissions []models.RolePermission
	for _, permissionId := range request.PermissionIds {
		permissions = append(permissions, models.RolePermission{
			PermissionID: permissionId,
		})
	}

	err := o.service.role.Update(ctx.Context(), roles.Model{
		Role: models.Role{
			ID:          id,
			Key:         request.Key,
			Name:        request.Name,
			Description: request.Description,
			IsActive:    request.IsActive,
		},
		Permissions: permissions,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) UpdateIsActive(ctx *fiber.Ctx) error {
	var request Request

	var id uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
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

	err := o.service.role.UpdateIsActive(ctx.Context(), models.Role{
		ID:       id,
		IsActive: request.IsActive,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) Delete(ctx *fiber.Ctx) error {

	var id uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	err := o.service.role.Delete(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) AddPermission(ctx *fiber.Ctx) error {
	var request Request

	var id uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
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

	err := o.service.role.AddPermission(ctx.Context(), models.RolePermission{
		RoleID:       id,
		PermissionID: request.ID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) RemovePermission(ctx *fiber.Ctx) error {

	var roleId uint
	if ctx.Params("role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		roleId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	var permissionId uint
	if ctx.Params("permission_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("permission_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		permissionId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	err := o.service.role.DeletePermission(ctx.Context(), models.RolePermission{
		RoleID:       roleId,
		PermissionID: permissionId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}
