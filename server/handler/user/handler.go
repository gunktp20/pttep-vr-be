package user

import (
	"fmt"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/users"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/pkg/utils/jwt"
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
	user *users.Service
}

func newHandler(user *users.Service) *handler {
	return &handler{
		service: &service{
			user: user,
		},
	}
}

func (o *handler) AddRole(ctx *fiber.Ctx) error {
	var request Request

	var id uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
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

	err := o.service.user.AddRole(ctx.Context(), models.UserRole{
		Name:        request.Name,
		UserID:      id,
		Email:       request.Email,
		RoleID:      request.RoleID,
		IsActive:    request.IsActive,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) AddRole2(ctx *fiber.Ctx) error {
	var request Request

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	err := o.service.user.AddRole(ctx.Context(), models.UserRole{
		Name:        request.Name,
		UserID:      0,
		Email:       request.Email,
		RoleID:      request.RoleID,
		IsActive:    request.IsActive,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) ChangeRole(ctx *fiber.Ctx) error {
	var request Request

	var userId uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	var userRoleId uint
	if ctx.Params("user_role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userRoleId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	err := o.service.user.ChangeRole(ctx.Context(), models.UserRole{
		ID:          userRoleId,
		Name:        request.Name,
		UserID:      userId,
		RoleID:      request.RoleID,
		IsActive:    request.IsActive,
		UpdatedDate: time.Now(),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) GetRole(ctx *fiber.Ctx) error {
	var paginate pagination.Pagination
	if err := ctx.QueryParser(&paginate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	data, count, err := o.service.user.GetRole(ctx.Context(), paginate.Get())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	result := ResultGet{
		List:  data,
		Total: int(count),
	}
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) GetRoleByUserRole(ctx *fiber.Ctx) error {

	var userId uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	var userRoleId uint
	if ctx.Params("user_role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userRoleId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	result, err := o.service.user.GetRoleByUserRole(ctx.Context(), models.UserRole{
		UserID: userId,
		ID:     userRoleId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) RemoveRole(ctx *fiber.Ctx) error {

	var userId uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	var userRoleId uint
	if ctx.Params("user_role_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_role_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userRoleId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	err := o.service.user.RemoveRole(ctx.Context(), models.UserRole{
		UserID:   userId,
		ID:       userRoleId,
		IsRemove: true,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

func (o *handler) AddUserLogin(ctx *fiber.Ctx) error {
	var userId uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	var loginTypeId uint
	if ctx.Params("login_type_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("login_type_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		loginTypeId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("login type is required")))
	}

	var request models.UserLogin
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	request.UserID = userId
	request.LoginTypeID = loginTypeId
	userLogin, err := o.service.user.AddLogin(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	//create token
	jwtObj := jwt.JWTAuthService()
	token := jwtObj.GenerateToken(userLogin.Username, true)
	result := Result{
		Token:       token,
		Username:    userLogin.Username,
		UserID:      userLogin.UserID,
		LoginTypeID: loginTypeId,
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) Get(ctx *fiber.Ctx) error {
	var id uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	result, err := o.service.user.Get(ctx.Context(), models.User{ID: id})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) GetUserByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	if code == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("code is required")))
	}

	result, err := o.service.user.GetUserByCode(ctx.Context(), code)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) GetPermission(ctx *fiber.Ctx) error {
	var id uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		id = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	permission, err := o.service.user.GetPermission(ctx.Context(), models.User{ID: id})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), permission, nil))
}

func (o *handler) Update(ctx *fiber.Ctx) error {
	var userId uint
	if ctx.Params("user_id") != "" {
		_id, err := strconv.Atoi(ctx.Params("user_id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
		}
		userId = uint(_id)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, fmt.Errorf("id required")))
	}

	var request models.User
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	request.ID = userId
	_, err := o.service.user.Update(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}
