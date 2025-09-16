package authentication

import (
	"fmt"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/authentications"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/server/response"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service *service
}

type service struct {
	authentication *authentications.Service
}

func newHandler(authentication *authentications.Service) *handler {
	return &handler{
		service: &service{
			authentication: authentication,
		},
	}
}

// Login
// @Summary User login
// @Description Authenticate user with username and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} response.Response{data=Result} "Login successful"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Invalid credentials"
// @Router /auth/login [post]
func (h *handler) Login(ctx *fiber.Ctx) error {

	//if pttep check case
	var request models.UserLogin
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("error"), nil, err))
	}

	//todo if pttep unpack token and get information
	//todo step1, check user in db, if invalid create this (user, userLogin)

	//service
	user, token, err := h.service.authentication.Login(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("authentication_invalid_username_or_password"), nil, err))
	}
	result := Result{
		Token: token,
		User:  user,
	}
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}

// Logout
// @Summary User logout
// @Description Logout current user session
// @Tags Authentication
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response "Logout successful"
// @Router /auth/logout [post]
func (h *handler) Logout(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Get("Authorization"))
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, nil))
}

// GetTypes
// @Summary Get login types
// @Description Retrieve available login types
// @Tags Authentication
// @Produce json
// @Success 200 {object} response.Response{data=ResultGetTypes} "Login types retrieved successfully"
// @Failure 404 {object} response.Response "Login types not found"
// @Router /auth/types [get]
func (h *handler) GetTypes(ctx *fiber.Ctx) error {
	result := ResultGetTypes{}
	types, err := h.service.authentication.GetTypes(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("authentication_type_not_found"), result, err))
	}
	result.List = types
	result.Total = len(types)
	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), result, nil))
}
