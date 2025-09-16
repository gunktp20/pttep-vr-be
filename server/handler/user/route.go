package user

import (
	"net/http"
	"pttep-vr-api/pkg/services/users"
	"pttep-vr-api/server/route"
)

const (
	prefix = "/users"
)

func Route(service *users.Service) []*route.Route {
	handler := newHandler(service)
	routes := []*route.Route{
		{
			Name:        "Get Roles",
			Description: "Get Roles",
			Method:      http.MethodGet,
			Path:        prefix + "/roles",
			Middleware:  nil,
			HandlerFunc: handler.GetRole,
			Test:        false,
		},
		{
			Name:        "Get User",
			Description: "Get User",
			Method:      http.MethodGet,
			Path:        prefix + "/:user_id",
			Middleware:  nil,
			HandlerFunc: handler.Get,
			Test:        false,
		},
		{
			Name:        "Get User by Code",
			Description: "Get User by Identification Code",
			Method:      http.MethodGet,
			Path:        prefix + "/identification-code/:code",
			Middleware:  nil,
			HandlerFunc: handler.GetUserByCode,
			Test:        false,
		},
		{
			Name:        "Get User",
			Description: "Get User",
			Method:      http.MethodPut,
			Path:        prefix + "/:user_id",
			Middleware:  nil,
			HandlerFunc: handler.Update,
			Test:        false,
		},
		{
			Name:        "Get Roles",
			Description: "Get Roles",
			Method:      http.MethodGet,
			Path:        prefix + "/:user_id/roles/:user_role_id",
			Middleware:  nil,
			HandlerFunc: handler.GetRoleByUserRole,
			Test:        false,
		},
		{
			Name:        "Add Roles",
			Description: "Add Roles",
			Method:      http.MethodPost,
			Path:        prefix + "/:user_id/roles",
			Middleware:  nil,
			HandlerFunc: handler.AddRole,
			Test:        false,
		},
		{
			Name:        "Add Roles",
			Description: "Add Roles",
			Method:      http.MethodPost,
			Path:        prefix + "/roles",
			Middleware:  nil,
			HandlerFunc: handler.AddRole2,
			Test:        false,
		},
		{
			Name:        "Change Roles",
			Description: "Change Roles",
			Method:      http.MethodPut,
			Path:        prefix + "/:user_id/roles/:user_role_id",
			Middleware:  nil,
			HandlerFunc: handler.ChangeRole,
			Test:        false,
		},
		{
			Name:        "Delete Roles",
			Description: "Delete Roles",
			Method:      http.MethodDelete,
			Path:        prefix + "/:user_id/roles/:user_role_id",
			Middleware:  nil,
			HandlerFunc: handler.RemoveRole,
			Test:        false,
		},
		{
			Name:        "Add User Login",
			Description: "Add User Login",
			Method:      http.MethodPost,
			Path:        prefix + "/:user_id/authentication/:login_type_id",
			Middleware:  nil,
			HandlerFunc: handler.AddUserLogin,
			Test:        false,
		},
		{
			Name:        "Get User Permission",
			Description: "Get User Permission",
			Method:      http.MethodGet,
			Path:        prefix + "/:user_id/permissions",
			Middleware:  nil,
			HandlerFunc: handler.GetPermission,
			Test:        false,
		},
	}
	return routes
}
