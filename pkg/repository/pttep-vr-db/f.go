package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/gormDB"
	"pttep-vr-api/pkg/utils/pagination"

	"gorm.io/gorm"
)

func New(o gormDB.Interface) *Operator {
	return &Operator{
		database: o,
	}
}

type Operator struct {
	database gormDB.Interface
}

type Interface interface {
	InsertOneGameTransaction(models.GameTransaction) (models.GameTransaction, error)
	InsertOneGameTransactionTemp(models.GameTransactionTemp) (models.GameTransactionTemp, error)
	FindSettingGameUser(string) ([]models.SessionWithDetails, error)
	FindSettingGameList() ([]models.SessionList, error)
	FindSettingGameById(uint) ([]models.QuestionList, error)
	UpdateSettingGameById(models.Question) (models.Question, error)
	FindGameTransactions(string, float64, bool, string, string, string, string, string, string, int, int) ([]models.GameTransactionResponse, int, error)
	SumTimeTransGmaeUser(string, string) (models.DashboardTotal, error)
	GetDistinctUsernameCountByDateRange(string, string) (int64, error)
	GetAvgOfMaxAverageScore(string, string) (float64, error)
	GetMaxFalseIsPassByQuestion(string, string) ([]models.GameFailTransaction, error)
	GetGraphScoreSevenDay(string) ([]models.GraphTransaction, error)
	GetPassRateByDate(string, string) (models.DashboardPassRate, error)
	GetMostQuestionByDate(string, string) ([]models.QuestionPlayCount, error)
	GetTopQuestionsByTime(string, string) ([]models.QuestionTimeStats, error)
	GetTotalScoreForUsersWithDateRange(string, string) ([]models.UserTotalScore, error)
	GetTopQuestionsByPlayCount(string, string) ([]models.QuestionPlayCount, error)

	DB() *gorm.DB

	InsertOneSettingGame(models.Session) (models.Session, error)
	InsertOneSettingGameQuest(models.Question) (models.Question, error)
	FindRportTransGmaeUser() ([]models.GameTransaction, error)
	FindDashboardTransGmaeUser() ([]models.GameTransaction, error)

	FindUsers() ([]models.User, error)
	FindOneUsers(uint) (models.User, error)
	FindOneUsersByEmail(models.User) (models.User, error)
	InsertOneUsers(models.User) (models.User, error)
	UpdateOneUsers(models.User) (models.User, error)
	FindOneUserByCode(string) (models.User, error)

	InsertOneUsersTemp(models.UserTemp) (models.UserTemp, error)
	FindOneUsersTempByUsername(string) (models.UserTemp, error)

	// Login Types
	FindLoginTypes(ctx context.Context) ([]models.LoginType, error)
	FindOneLoginTypes(ctx context.Context, v models.LoginType) (models.LoginType, error)
	InsertOneLoginTypes(ctx context.Context, v models.LoginType) (models.LoginType, error)

	// Permissions
	FindPermissions(ctx context.Context, p *pagination.Pagination) ([]models.Permission, int64, error)
	FindPermissionsIn(ctx context.Context, v []models.Permission) ([]models.Permission, int64, error)
	FindOnePermissions(ctx context.Context, id uint) (models.Permission, error)
	InsertOnePermissions(ctx context.Context, v models.Permission) (models.Permission, error)
	UpdateOnePermissions(ctx context.Context, v models.Permission) (models.Permission, error)
	UpdateOnePermissionsName(ctx context.Context, v models.Permission) (models.Permission, error)
	UpdateOnePermissionsIsActive(ctx context.Context, v models.Permission) (models.Permission, error)
	UpdateOnePermissionsIsRemove(ctx context.Context, v models.Permission) (models.Permission, error)

	// Role Permissions
	FindRolePermissionsJoinPermissionsInRoleID(ctx context.Context, v []models.RolePermission) ([]models.RolePermission, int64, error)
	InsertOneRolePermissions(ctx context.Context, v models.RolePermission) (models.RolePermission, error)
	InsertManyRolePermissions(ctx context.Context, v []models.RolePermission) ([]models.RolePermission, error)
	DeleteOneRolePermissions(ctx context.Context, v models.RolePermission) error
	DeleteOneRolePermissionsByRoleIDAndPermissionID(ctx context.Context, v models.RolePermission) error
	DeleteManyRolePermissionsByRoleID(ctx context.Context, v models.RolePermission) error

	// Roles
	FindRoles(ctx context.Context, p *pagination.Pagination) ([]models.Role, int64, error)
	FindOneRoles(ctx context.Context, id uint) (models.Role, error)
	FindOneRolesByIsDefault(ctx context.Context, status bool) (models.Role, error)
	InsertOneRoles(ctx context.Context, v models.Role) (models.Role, error)
	UpdateOneRoles(ctx context.Context, v models.Role) (models.Role, error)
	UpdateOneRolesName(ctx context.Context, v models.Role) (models.Role, error)
	UpdateOneRolesIsActive(ctx context.Context, v models.Role) (models.Role, error)
	UpdateOneRolesIsRemove(ctx context.Context, v models.Role) (models.Role, error)

	// User Login
	FindUserLogins(ctx context.Context) ([]models.UserLogin, error)
	FindOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error)
	FindOneUserLoginsByUsername(ctx context.Context, v models.UserLogin) (models.UserLogin, error)
	FindOneUserLoginsByUsernameAndPassword(ctx context.Context, v models.UserLogin) (models.UserLogin, error)
	InsertOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error)
	UpdateOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error)

	// User Role
	FindUserRoles(ctx context.Context) ([]models.UserRole, error)
	FindUserRolesByUserID(ctx context.Context, v models.UserRole) ([]models.UserRole, error)
	FindOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error)
	InsertOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error)
	UpdateOneUserRolesRole(ctx context.Context, v models.UserRole) (models.UserRole, error)
	UpdateOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error)
	UpdateUserRolesByEmail(ctx context.Context, v models.UserRole) error
	UpdateOneUserRolesIsActive(ctx context.Context, v models.UserRole) error
	UpdateOneUserRolesIsRemove(ctx context.Context, v models.UserRole) error
	FindUserRolesLeftJoinUserAndRolePermissionAndRole(ctx context.Context, p *pagination.Pagination) ([]models.UserRole, int64, error)
	FindOneUserRolesLeftJoinUserAndRolePermissionAndRole(ctx context.Context, v models.UserRole) (models.UserRole, error)
}
