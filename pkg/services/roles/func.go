package roles

import (
	"context"
	"pttep-vr-api/pkg/config"
	pttep_vr_db "pttep-vr-api/pkg/repository/pttep-vr-db"

	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"
)

type Service struct {
	config     *config.Config
	repository pttep_vr_db.Interface
}

type pkg struct {
	userManagement UmInterface
}

func New(config *config.Config, repository pttep_vr_db.Interface) *Service {
	return &Service{
		config:     config,
		repository: repository,
	}
}

type Model struct {
	Role        models.Role
	Permissions []models.RolePermission
}

type UmInterface interface {
	Create(context.Context, models.Role) (models.Role, error)
	AddManyPermission(context.Context, models.Role, []models.Permission) error
	Delete(context.Context, models.Role) error
	RemoveOnePermissionByRoleIDAndPermission(ctx context.Context, v models.RolePermission) error
	GetAll(context.Context, *pagination.Pagination) ([]models.Role, int64, error)
	GetPermission(context.Context, []models.RolePermission) ([]models.RolePermission, int64, error)
	GetByID(context.Context, uint) (models.Role, error)
	UpdateIsActive(context.Context, models.Role) error
	Update(context.Context, models.Role) (models.Role, error)
	RemoveAllPermission(context.Context, models.RolePermission) error
	AddOnePermission(context.Context, models.RolePermission) error
}
