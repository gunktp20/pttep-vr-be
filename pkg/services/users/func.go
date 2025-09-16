package users

import (
	"context"
	"pttep-vr-api/pkg/config"
	"pttep-vr-api/pkg/models"
	pttep_vr_db "pttep-vr-api/pkg/repository/pttep-vr-db"
	"pttep-vr-api/pkg/services/roles"

	"pttep-vr-api/pkg/utils/pagination"
)

type Service struct {
	config     *config.Config
	repository pttep_vr_db.Interface
}

type UserManagement struct {
	Authentication UmAuthenticationInterface
	Role           UmRoleInterface
	User           UmUserInterface
}

func New(config *config.Config, repository pttep_vr_db.Interface) *Service {
	return &Service{
		config:     config,
		repository: repository,
	}
}

type Model struct {
	User     models.User
	UserRole models.UserRole
	Roles    roles.Model
}

type RepositoryInterface interface {
	FindOneUsersByEmail(models.User) (models.User, error)
	FindOneUsersTempByUsername(string) (models.UserTemp, error)
	InsertOneUsersTemp(models.UserTemp) (models.UserTemp, error)
	InsertOneUsers(models.User) (models.User, error)
	FindOneUsers(uint) (models.User, error)
	UpdateOneUsers(models.User) (models.User, error)
}

type UmRoleInterface interface {
	UpdateUserIdByEmail(context.Context, models.UserRole) error
	GetDefault(context.Context) (models.Role, error)
	GetPermission(context.Context, []models.RolePermission) ([]models.RolePermission, int64, error)
}

type UmUserInterface interface {
	AddRole(context.Context, models.UserRole) error
	ChangeRole(context.Context, models.UserRole) error
	GetUserRole(context.Context, models.UserRole) ([]models.UserRole, error)
	UpdateUserLogins(context.Context, models.UserLogin) (models.UserLogin, error)
	RemoveRole(context.Context, models.UserRole) error
	GetRoleByUserRole(context.Context, models.UserRole) (models.UserRole, error)
	GetRole(context.Context, *pagination.Pagination) ([]models.UserRole, int64, error)
}

type UmAuthenticationInterface interface {
	Add(context.Context, models.UserLogin) (models.UserLogin, error)
}
