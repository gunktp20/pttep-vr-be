package registration

import (
	"fmt"
	"net/http"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/users"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/pkg/utils/jwt"
	"pttep-vr-api/server/response"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service *service
}

type service struct {
	users *users.Service
}

func newHandler(users *users.Service) *handler {
	return &handler{
		service: &service{
			users: users,
		},
	}
}

func (o *handler) NormalTemp(ctx *fiber.Ctx) error {
	var request Request
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("ping_error"), nil, err))
	}

	user, err := o.service.users.CreateTemp(ctx.Context(), models.UserTemp{
		Name:        request.Name,
		Surname:     request.Surname,
		Email:       request.Email,
		Tel:         request.Tel,
		Group:       request.Group,
		Company:     request.Company,
		IsActive:    true,
		IsRemove:    false,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	})
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("register"), nil, err))
	}

	var result Result
	result.Code = user.Username
	return ctx.Status(http.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

func (o *handler) Normal(ctx *fiber.Ctx) error {

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

	var request models.User
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	var result Result
	jwtObj := jwt.JWTAuthService()
	result.Token = jwtObj.GenerateToken(request.Email, true)

	request.LoginTypeID = loginTypeId

	user, err := o.service.users.Create(ctx.Context(), request)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("register"), nil, err))
	}

	result.Username = user.Email
	result.UserID = user.ID
	result.LoginTypeID = user.LoginTypeID
	return ctx.Status(http.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}
