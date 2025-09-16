package mock

import (
	"context"
	"errors"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"

	"gorm.io/gorm"
)

// MockRepository implements pttep_vr_db.Interface
type MockRepository struct {
	ShouldReturnError bool
	ReturnData        interface{}
}

// Game Transaction methods
func (m *MockRepository) InsertOneGameTransaction(game models.GameTransaction) (models.GameTransaction, error) {
	if m.ShouldReturnError {
		return models.GameTransaction{}, errors.New("mock error")
	}
	return game, nil
}

func (m *MockRepository) InsertOneGameTransactionTemp(game models.GameTransactionTemp) (models.GameTransactionTemp, error) {
	if m.ShouldReturnError {
		return models.GameTransactionTemp{}, errors.New("mock error")
	}
	return game, nil
}

func (m *MockRepository) FindSettingGameUser(key string) ([]models.SessionWithDetails, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	if m.ReturnData != nil {
		return m.ReturnData.([]models.SessionWithDetails), nil
	}
	return []models.SessionWithDetails{{SessionID: 1}}, nil
}

func (m *MockRepository) FindSettingGameList() ([]models.SessionList, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.SessionList{{SessionID: 1, SessionName: "Test Session"}}, nil
}

func (m *MockRepository) FindSettingGameById(id uint) ([]models.QuestionList, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.QuestionList{{QuestId: 1, QuestName: "Test?"}}, nil
}

func (m *MockRepository) UpdateSettingGameById(req models.Question) (models.Question, error) {
	if m.ShouldReturnError {
		return models.Question{}, errors.New("mock error")
	}
	return req, nil
}

func (m *MockRepository) FindGameTransactions(username string, score float64, isPass bool, baseName, sessionName, status, email, startDate, endDate string, limit, page int) ([]models.GameTransactionResponse, int, error) {
	if m.ShouldReturnError {
		return nil, 0, errors.New("mock error")
	}
	return []models.GameTransactionResponse{{Username: username}}, 1, nil
}

func (m *MockRepository) SumTimeTransGmaeUser(startDate, endDate string) (models.DashboardTotal, error) {
	if m.ShouldReturnError {
		return models.DashboardTotal{}, errors.New("mock error")
	}
	return models.DashboardTotal{Total: 100}, nil
}

func (m *MockRepository) GetDistinctUsernameCountByDateRange(startDate, endDate string) (int64, error) {
	if m.ShouldReturnError {
		return 0, errors.New("mock error")
	}
	return 50, nil
}

func (m *MockRepository) GetAvgOfMaxAverageScore(startDate, endDate string) (float64, error) {
	if m.ShouldReturnError {
		return 0, errors.New("mock error")
	}
	return 75.5, nil
}

func (m *MockRepository) GetMaxFalseIsPassByQuestion(startDate, endDate string) ([]models.GameFailTransaction, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.GameFailTransaction{{QuestionID: 1}}, nil
}

func (m *MockRepository) GetGraphScoreSevenDay(base string) ([]models.GraphTransaction, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.GraphTransaction{{Date: "2024-01-01"}}, nil
}

func (m *MockRepository) GetPassRateByDate(startDate, endDate string) (models.DashboardPassRate, error) {
	if m.ShouldReturnError {
		return models.DashboardPassRate{}, errors.New("mock error")
	}
	return models.DashboardPassRate{PassRate: 85.5}, nil
}

func (m *MockRepository) GetMostQuestionByDate(startDate, endDate string) ([]models.QuestionPlayCount, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.QuestionPlayCount{{QuestionName: "Quest 1", PlayCount: 100}}, nil
}

func (m *MockRepository) GetTopQuestionsByTime(startDate, endDate string) ([]models.QuestionTimeStats, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.QuestionTimeStats{{QuestionName: "Quest 2", TotalTime: 30.5}}, nil
}

func (m *MockRepository) GetTotalScoreForUsersWithDateRange(startDate, endDate string) ([]models.UserTotalScore, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.UserTotalScore{{Username: "user1", TotalScore: 500}}, nil
}

func (m *MockRepository) GetTopQuestionsByPlayCount(startDate, endDate string) ([]models.QuestionPlayCount, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.QuestionPlayCount{{QuestionName: "Quest1", PlayCount: 150}}, nil
}

// Database method
func (m *MockRepository) DB() *gorm.DB {
	return nil // Mock ไม่ต้องใช้ DB จริง
}

// Settings methods
func (m *MockRepository) InsertOneSettingGame(session models.Session) (models.Session, error) {
	if m.ShouldReturnError {
		return models.Session{}, errors.New("mock error")
	}
	return session, nil
}

func (m *MockRepository) InsertOneSettingGameQuest(question models.Question) (models.Question, error) {
	if m.ShouldReturnError {
		return models.Question{}, errors.New("mock error")
	}
	return question, nil
}

func (m *MockRepository) FindRportTransGmaeUser() ([]models.GameTransaction, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.GameTransaction{}, nil
}

func (m *MockRepository) FindDashboardTransGmaeUser() ([]models.GameTransaction, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.GameTransaction{}, nil
}

// User methods
func (m *MockRepository) FindUsers() ([]models.User, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.User{}, nil
}

func (m *MockRepository) FindOneUsers(id uint) (models.User, error) {
	if m.ShouldReturnError {
		return models.User{}, errors.New("mock error")
	}
	return models.User{ID: id}, nil
}

func (m *MockRepository) FindOneUsersByEmail(user models.User) (models.User, error) {
	if m.ShouldReturnError {
		return models.User{}, errors.New("mock error")
	}
	return user, nil
}

func (m *MockRepository) InsertOneUsers(user models.User) (models.User, error) {
	if m.ShouldReturnError {
		return models.User{}, errors.New("mock error")
	}
	return user, nil
}

func (m *MockRepository) UpdateOneUsers(user models.User) (models.User, error) {
	if m.ShouldReturnError {
		return models.User{}, errors.New("mock error")
	}
	return user, nil
}

func (m *MockRepository) FindOneUserByCode(code string) (models.User, error) {
	if m.ShouldReturnError {
		return models.User{}, errors.New("mock error")
	}
	return models.User{Code: code}, nil
}

// User Temp methods
func (m *MockRepository) InsertOneUsersTemp(userTemp models.UserTemp) (models.UserTemp, error) {
	if m.ShouldReturnError {
		return models.UserTemp{}, errors.New("mock error")
	}
	return userTemp, nil
}

func (m *MockRepository) FindOneUsersTempByUsername(username string) (models.UserTemp, error) {
	if m.ShouldReturnError {
		return models.UserTemp{}, errors.New("mock error")
	}
	return models.UserTemp{Username: username}, nil
}

// Login Types methods
func (m *MockRepository) FindLoginTypes(ctx context.Context) ([]models.LoginType, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.LoginType{}, nil
}

func (m *MockRepository) FindOneLoginTypes(ctx context.Context, v models.LoginType) (models.LoginType, error) {
	if m.ShouldReturnError {
		return models.LoginType{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) InsertOneLoginTypes(ctx context.Context, v models.LoginType) (models.LoginType, error) {
	if m.ShouldReturnError {
		return models.LoginType{}, errors.New("mock error")
	}
	return v, nil
}

// Permission methods
func (m *MockRepository) FindPermissions(ctx context.Context, p *pagination.Pagination) ([]models.Permission, int64, error) {
	if m.ShouldReturnError {
		return nil, 0, errors.New("mock error")
	}
	return []models.Permission{}, 0, nil
}

func (m *MockRepository) FindPermissionsIn(ctx context.Context, v []models.Permission) ([]models.Permission, int64, error) {
	if m.ShouldReturnError {
		return nil, 0, errors.New("mock error")
	}
	return v, int64(len(v)), nil
}

func (m *MockRepository) FindOnePermissions(ctx context.Context, id uint) (models.Permission, error) {
	if m.ShouldReturnError {
		return models.Permission{}, errors.New("mock error")
	}
	return models.Permission{ID: id}, nil
}

func (m *MockRepository) InsertOnePermissions(ctx context.Context, v models.Permission) (models.Permission, error) {
	if m.ShouldReturnError {
		return models.Permission{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOnePermissions(ctx context.Context, v models.Permission) (models.Permission, error) {
	if m.ShouldReturnError {
		return models.Permission{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOnePermissionsName(ctx context.Context, v models.Permission) (models.Permission, error) {
	if m.ShouldReturnError {
		return models.Permission{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOnePermissionsIsActive(ctx context.Context, v models.Permission) (models.Permission, error) {
	if m.ShouldReturnError {
		return models.Permission{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOnePermissionsIsRemove(ctx context.Context, v models.Permission) (models.Permission, error) {
	if m.ShouldReturnError {
		return models.Permission{}, errors.New("mock error")
	}
	return v, nil
}

// Role Permission methods
func (m *MockRepository) FindRolePermissionsJoinPermissionsInRoleID(ctx context.Context, v []models.RolePermission) ([]models.RolePermission, int64, error) {
	if m.ShouldReturnError {
		return nil, 0, errors.New("mock error")
	}
	return v, int64(len(v)), nil
}

func (m *MockRepository) InsertOneRolePermissions(ctx context.Context, v models.RolePermission) (models.RolePermission, error) {
	if m.ShouldReturnError {
		return models.RolePermission{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) InsertManyRolePermissions(ctx context.Context, v []models.RolePermission) ([]models.RolePermission, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) DeleteOneRolePermissions(ctx context.Context, v models.RolePermission) error {
	if m.ShouldReturnError {
		return errors.New("mock error")
	}
	return nil
}

func (m *MockRepository) DeleteOneRolePermissionsByRoleIDAndPermissionID(ctx context.Context, v models.RolePermission) error {
	if m.ShouldReturnError {
		return errors.New("mock error")
	}
	return nil
}

func (m *MockRepository) DeleteManyRolePermissionsByRoleID(ctx context.Context, v models.RolePermission) error {
	if m.ShouldReturnError {
		return errors.New("mock error")
	}
	return nil
}

// Role methods
func (m *MockRepository) FindRoles(ctx context.Context, p *pagination.Pagination) ([]models.Role, int64, error) {
	if m.ShouldReturnError {
		return nil, 0, errors.New("mock error")
	}
	return []models.Role{}, 0, nil
}

func (m *MockRepository) FindOneRoles(ctx context.Context, id uint) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return models.Role{ID: id}, nil
}

func (m *MockRepository) FindOneRolesByIsDefault(ctx context.Context, status bool) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return models.Role{IsDefault: status}, nil
}

func (m *MockRepository) InsertOneRoles(ctx context.Context, v models.Role) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneRoles(ctx context.Context, v models.Role) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneRolesName(ctx context.Context, v models.Role) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneRolesIsActive(ctx context.Context, v models.Role) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneRolesIsRemove(ctx context.Context, v models.Role) (models.Role, error) {
	if m.ShouldReturnError {
		return models.Role{}, errors.New("mock error")
	}
	return v, nil
}

// User Login methods
func (m *MockRepository) FindUserLogins(ctx context.Context) ([]models.UserLogin, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.UserLogin{}, nil
}

func (m *MockRepository) FindOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	if m.ShouldReturnError {
		return models.UserLogin{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) FindOneUserLoginsByUsername(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	if m.ShouldReturnError {
		return models.UserLogin{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) FindOneUserLoginsByUsernameAndPassword(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	if m.ShouldReturnError {
		return models.UserLogin{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) InsertOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	if m.ShouldReturnError {
		return models.UserLogin{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	if m.ShouldReturnError {
		return models.UserLogin{}, errors.New("mock error")
	}
	return v, nil
}

// User Role methods
func (m *MockRepository) FindUserRoles(ctx context.Context) ([]models.UserRole, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.UserRole{}, nil
}

func (m *MockRepository) FindUserRolesByUserID(ctx context.Context, v models.UserRole) ([]models.UserRole, error) {
	if m.ShouldReturnError {
		return nil, errors.New("mock error")
	}
	return []models.UserRole{v}, nil
}

func (m *MockRepository) FindOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	if m.ShouldReturnError {
		return models.UserRole{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) InsertOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	if m.ShouldReturnError {
		return models.UserRole{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneUserRolesRole(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	if m.ShouldReturnError {
		return models.UserRole{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	if m.ShouldReturnError {
		return models.UserRole{}, errors.New("mock error")
	}
	return v, nil
}

func (m *MockRepository) UpdateUserRolesByEmail(ctx context.Context, v models.UserRole) error {
	if m.ShouldReturnError {
		return errors.New("mock error")
	}
	return nil
}

func (m *MockRepository) UpdateOneUserRolesIsActive(ctx context.Context, v models.UserRole) error {
	if m.ShouldReturnError {
		return errors.New("mock error")
	}
	return nil
}

func (m *MockRepository) UpdateOneUserRolesIsRemove(ctx context.Context, v models.UserRole) error {
	if m.ShouldReturnError {
		return errors.New("mock error")
	}
	return nil
}

func (m *MockRepository) FindUserRolesLeftJoinUserAndRolePermissionAndRole(ctx context.Context, p *pagination.Pagination) ([]models.UserRole, int64, error) {
	if m.ShouldReturnError {
		return nil, 0, errors.New("mock error")
	}
	return []models.UserRole{}, 0, nil
}

func (m *MockRepository) FindOneUserRolesLeftJoinUserAndRolePermissionAndRole(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	if m.ShouldReturnError {
		return models.UserRole{}, errors.New("mock error")
	}
	return v, nil
}
