package role

import (
	"net/http"
	"pttep-vr-api/pkg/services/roles"
	"pttep-vr-api/server/route"
)

const (
	prefix = "/roles"
)

func Route(service *roles.Service) []*route.Route {
	handler := newHandler(service)
	routes := []*route.Route{
		{
			Name:        "Get Roles",
			Description: "Get Roles",
			Method:      http.MethodGet,
			Path:        prefix,
			Middleware:  nil,
			HandlerFunc: handler.Get,
			Test:        false,
		},
		{
			Name:        "Get Roles And Permissions",
			Description: "Get Roles And Permissions",
			Method:      http.MethodGet,
			Path:        prefix + "/permissions",
			Middleware:  nil,
			HandlerFunc: handler.GetAndPermission,
			Test:        false,
		},
		{
			Name:        "Get One Roles",
			Description: "Get One Roles",
			Method:      http.MethodGet,
			Path:        prefix + "/:role_id",
			Middleware:  nil,
			HandlerFunc: handler.GetByID,
			Test:        false,
		},
		{
			Name:        "Get One Roles And Permissions",
			Description: "Get One Roles And Permissions",
			Method:      http.MethodGet,
			Path:        prefix + "/:role_id/permissions",
			Middleware:  nil,
			HandlerFunc: handler.GetByIDAndPermission,
			Test:        false,
		},
		{
			Name:        "Create Roles",
			Description: "Create Roles",
			Method:      http.MethodPost,
			Path:        prefix,
			Middleware:  nil,
			HandlerFunc: handler.Create,
			Test:        false,
		},
		{
			Name:        "Update Roles",
			Description: "Update Roles",
			Method:      http.MethodPut,
			Path:        prefix + "/:role_id",
			Middleware:  nil,
			HandlerFunc: handler.Update,
			Test:        false,
		},
		{
			Name:        "Update Roles",
			Description: "Update Roles",
			Method:      http.MethodPut,
			Path:        prefix + "/:role_id/status",
			Middleware:  nil,
			HandlerFunc: handler.UpdateIsActive,
			Test:        false,
		},
		{
			Name:        "Add Permission",
			Description: "Add Permission",
			Method:      http.MethodPut,
			Path:        prefix + "/:role_id/permissions",
			Middleware:  nil,
			HandlerFunc: handler.AddPermission,
			Test:        false,
		},
		{
			Name:        "Delete Roles",
			Description: "Delete Roles",
			Method:      http.MethodDelete,
			Path:        prefix + "/:role_id",
			Middleware:  nil,
			HandlerFunc: handler.Delete,
			Test:        false,
		},
		{
			Name:        "Remove Permission",
			Description: "Remove Permission",
			Method:      http.MethodDelete,
			Path:        prefix + "/:role_id/permissions/:permission_id",
			Middleware:  nil,
			HandlerFunc: handler.RemovePermission,
			Test:        false,
		},
	}
	return routes
}
