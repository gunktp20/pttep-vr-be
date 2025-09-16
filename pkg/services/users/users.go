package users

import (
	"context"
	"errors"
	"fmt"
	"pttep-vr-api/pkg/constant"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/roles"
	"pttep-vr-api/pkg/utils/pagination"
	"pttep-vr-api/pkg/utils/random"
	"time"
)

func (o *Service) AddLogin(ctx context.Context, userLogin models.UserLogin) (models.UserLogin, error) {
	userLogin.IsRemove = false
	userLogin.IsActive = true
	userLogin.CreatedDate = time.Now()
	userLogin.UpdatedDate = time.Now()

	var loginType models.LoginType
	var err error
	loginType.SetID(userLogin.LoginTypeID)
	loginType, err = o.repository.FindOneLoginTypes(ctx, loginType)
	if err != nil {
		return models.UserLogin{}, err
	}

	switch constant.TYPELoginType(loginType.Key) {
	case constant.LoginType.Email(), constant.LoginType.Contractor():
		userLogin, err := o.repository.FindOneUserLoginsByUsername(ctx, userLogin)
		if err != nil {
			return models.UserLogin{}, err
		}
		if fmt.Sprintf("%v", userLogin.ID) != "0" {
			return models.UserLogin{}, errors.New("This username already exists in the system.")
		}
	default:
		return models.UserLogin{}, errors.New("invalid login type")
	}

	user, err := o.repository.InsertOneUserLogins(ctx, userLogin)
	if err != nil {
		return models.UserLogin{}, err
	}

	return user, nil
}

func (o *Service) AddRole(ctx context.Context, v models.UserRole) error {
	user, _ := o.repository.FindOneUsersByEmail(models.User{
		Email: v.Email,
	})
	v.UserID = user.ID

	_, err := o.repository.InsertOneUserRoles(ctx, v)
	if err != nil {
		return err
	}
	return nil
}

func (o *Service) ChangeRole(ctx context.Context, v models.UserRole) error {
	user, _ := o.repository.FindOneUsersByEmail(models.User{
		Email: v.Email,
	})
	v.UserID = user.ID

	_, err := o.repository.UpdateOneUserRoles(ctx, v)
	if err != nil {
		return err
	}

	return nil
}

func (o *Service) CreateTemp(ctx context.Context, user models.UserTemp) (models.UserTemp, error) {
	//random number 4 digit
	for {
		user.Username = random.New([]random.CharacterSet{random.Number}, 4)
		u, err := o.repository.FindOneUsersTempByUsername(user.Username)
		if err != nil {
			return models.UserTemp{}, err
		}
		if u.ID == 0 {
			break
		}
	}

	var err error
	user, err = o.repository.InsertOneUsersTemp(user)
	if err != nil {
		fmt.Println(err)
		return models.UserTemp{}, err
	}
	return user, nil
}

func (o *Service) Create(ctx context.Context, user models.User) (models.User, error) {

	check, err := o.repository.FindOneUsersByEmail(user)
	if err != nil {
		return models.User{}, err
	}
	if check.ID != 0 {
		return models.User{}, errors.New("user already exists")
	}

	//random number 6 digit
	for {
		user.Code = random.New([]random.CharacterSet{random.Number}, 6)
		u, err := o.repository.FindOneUsersTempByUsername(user.Code)
		if err != nil {
			return models.User{}, err
		}
		if u.ID == 0 {
			break
		}
	}

	user.IsActive = true
	user.IsRemove = false
	user.CreatedDate = time.Now()
	user.UpdatedDate = time.Now()

	user, err = o.repository.InsertOneUsers(user)
	if err != nil {
		return models.User{}, err
	}
	//user_id in user_role where email = ?
	// _ = o.pkg.userManagement.Role.UpdateUserIdByEmail(ctx, models.UserRole{
	// 	Email:  user.Email,
	// 	UserID: user.ID,
	// })

	_ = o.repository.UpdateUserRolesByEmail(ctx, models.UserRole{
		Email:  user.Email,
		UserID: user.ID,
	})

	//add role player (default role)
	// role, err := o.pkg.userManagement.Role.GetDefault(ctx)
	role, err := o.repository.FindOneRolesByIsDefault(ctx, true)
	if err == nil {
		_, err := o.repository.InsertOneUserRoles(ctx, models.UserRole{
			Name:        role.Name,
			UserID:      user.ID,
			Email:       user.Email,
			RoleID:      role.ID,
			IsActive:    true,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		})
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (o *Service) GetPermission(ctx context.Context, v models.User) ([]models.Permission, error) {

	userRoles, err := o.repository.FindUserRolesByUserID(ctx, models.UserRole{
		UserID: v.ID,
	})
	if err != nil {
		return nil, err
	}

	var rolePermissions []models.RolePermission
	for _, role := range userRoles {
		rolePermissions = append(rolePermissions, models.RolePermission{
			RoleID: role.RoleID,
		})
	}

	rolePermissions, _, err = o.repository.FindRolePermissionsJoinPermissionsInRoleID(ctx, rolePermissions)
	if err != nil {
		return nil, err
	}

	var permissions []models.Permission
	check := make(map[uint]bool)
	for _, rolePermission := range rolePermissions {
		if _, ok := check[rolePermission.PermissionID]; !ok {
			permissions = append(permissions, *rolePermission.Permission)
			check[rolePermission.PermissionID] = true
		}
	}
	return permissions, nil
}

func (o *Service) GetRole(ctx context.Context, paginate *pagination.Pagination) ([]Model, int64, error) {
	userRoles, count, err := o.repository.FindUserRolesLeftJoinUserAndRolePermissionAndRole(ctx, paginate)
	if err != nil {
		return nil, 0, err
	}

	var rolePermissions []models.RolePermission
	for _, v := range userRoles {
		rolePermissions = append(rolePermissions, models.RolePermission{
			RoleID: v.Role.ID,
		})
	}
	rolePermissions, _, err = o.repository.FindRolePermissionsJoinPermissionsInRoleID(ctx, rolePermissions)
	if err != nil {
		return nil, 0, err
	}

	var data []Model
	for _, v := range userRoles {
		var _role roles.Model
		if v.Role != nil {
			_role.Role = *v.Role
		}

		for _, permission := range rolePermissions {
			if permission.RoleID == v.Role.ID {
				permission.Role = nil
				_role.Permissions = append(_role.Permissions, permission)
			}
		}
		d := Model{

			UserRole: v,
			Roles:    _role,
		}
		if v.User != nil {
			d.User = *v.User
		}

		d.UserRole.Role = nil
		d.UserRole.User = nil
		data = append(data, d)
	}

	return data, count, nil
}

func (o *Service) GetRoleByUserRole(ctx context.Context, value models.UserRole) (Model, error) {
	userRoles, err := o.repository.FindOneUserRolesLeftJoinUserAndRolePermissionAndRole(ctx, value)
	if err != nil {
		return Model{}, err
	}
	var rolePermissions []models.RolePermission
	rolePermissions = append(rolePermissions, models.RolePermission{
		RoleID: userRoles.Role.ID,
	})
	rolePermissions, _, err = o.repository.FindRolePermissionsJoinPermissionsInRoleID(ctx, rolePermissions)
	if err != nil {
		return Model{}, err
	}

	_role := roles.Model{
		Role: *userRoles.Role,
	}
	_userRole := userRoles
	for _, permission := range rolePermissions {
		if permission.RoleID == userRoles.Role.ID {
			permission.Role = nil
			_role.Permissions = append(_role.Permissions, permission)
		}
	}
	data := Model{
		User:     *userRoles.User,
		UserRole: _userRole,
		Roles:    _role,
	}
	data.UserRole.User = nil
	data.UserRole.Role = nil

	return data, nil
}

func (o *Service) GetUserByCode(ctx context.Context, code string) (models.User, error) {
	return o.repository.FindOneUserByCode(code)
}

func (o *Service) Get(ctx context.Context, v models.User) (models.User, error) {

	user, err := o.repository.FindOneUsers(v.ID)
	if err != nil {
		return user, err
	}

	var newCode string = ""
	if user.CodeExpiredAt == nil || (user.CodeExpiredAt != nil && time.Now().After(*user.CodeExpiredAt)) {
		for {
			newCode = random.New([]random.CharacterSet{random.Number}, 6)
			u, err := o.repository.FindOneUserByCode(newCode)
			if err != nil {
				return models.User{}, err
			}
			if u.ID == 0 {
				break
			}
		}

		expiredAt := time.Now().Add(2 * time.Minute)

		user.Code = newCode
		user.CodeExpiredAt = &expiredAt

		v, err = o.repository.UpdateOneUsers(user)
		if err != nil {
			return models.User{}, err
		}
	}

	return o.repository.FindOneUsers(v.ID)
}

func (o *Service) RemoveRole(ctx context.Context, v models.UserRole) error {
	return o.repository.UpdateOneUserRolesIsRemove(ctx, v)
}

func (o *Service) Update(ctx context.Context, v models.User) (models.User, error) {
	var err error
	v, err = o.repository.UpdateOneUsers(v)
	if err != nil {
		return models.User{}, err
	}

	_, err = o.repository.UpdateOneUserLogins(ctx, models.UserLogin{
		UserID:      v.ID,
		LoginTypeID: v.LoginTypeID,
		Username:    v.Email,
	})
	if err != nil {
		return models.User{}, err
	}

	return v, nil
}
