package roles

import (
	"context"
	models "pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"
)

func (o *Service) Create(ctx context.Context, v Model, permissionIDs []uint) (models.Role, error) {
	v.Role.IsRemove = false
	_role, err := o.repository.InsertOneRoles(ctx, v.Role)
	if err != nil {
		return _role, err
	}

	var permissions []models.RolePermission
	for _, permissionID := range permissionIDs {
		permissions = append(permissions, models.RolePermission{
			PermissionID: permissionID,
		})
	}

	if len(permissions) > 0 {
		var value []models.RolePermission
		for _, permission := range permissions {
			value = append(value, models.RolePermission{
				RoleID:       _role.ID,
				PermissionID: permission.PermissionID,
			})
		}
		_, err := o.repository.InsertManyRolePermissions(ctx, value)
		if err != nil {
			return models.Role{}, err
		}
	}

	return _role, nil
}

func (o *Service) Delete(ctx context.Context, id uint) error {
	_, err := o.repository.UpdateOneRolesIsRemove(ctx, models.Role{
		ID:       id,
		IsRemove: true,
	})
	if err != nil {
		return err
	}

	return nil

}

func (o *Service) DeletePermission(ctx context.Context, v models.RolePermission) error {
	err := o.repository.DeleteOneRolePermissionsByRoleIDAndPermissionID(ctx, v)
	if err != nil {
		return err
	}
	return nil
}

func (o *Service) GetByID(ctx context.Context, id uint) (models.Role, error) {
	return o.repository.FindOneRoles(ctx, id)
}

func (o *Service) GetByIDAndPermission(ctx context.Context, id uint) (Model, error) {
	_role, err := o.GetByID(ctx, id)
	if err != nil {
		return Model{}, err
	}
	var permissions []models.RolePermission
	permissions = append(permissions, models.RolePermission{
		RoleID: _role.ID,
	})

	permissions, _, err = o.repository.FindRolePermissionsJoinPermissionsInRoleID(ctx, permissions)
	if err != nil {
		return Model{}, err
	}

	var data Model
	data.Role = _role
	for _, permission := range permissions {
		if permission.RoleID == _role.ID {
			permission.Role = nil
			data.Permissions = append(data.Permissions, permission)
		}
	}

	return data, nil
}

func (o *Service) Get(ctx context.Context, paginate *pagination.Pagination) ([]models.Role, int64, error) {
	return o.repository.FindRoles(ctx, paginate)
}

func (o *Service) GetAndPermission(ctx context.Context, paginate *pagination.Pagination) ([]Model, int64, error) {
	_roles, _count, err := o.Get(ctx, paginate)
	if err != nil {
		return []Model{}, 0, err
	}
	var permissions []models.RolePermission
	for _, role := range _roles {
		permissions = append(permissions, models.RolePermission{
			RoleID: role.ID,
		})
	}

	permissions, _, err = o.repository.FindRolePermissionsJoinPermissionsInRoleID(ctx, permissions)
	if err != nil {
		return []Model{}, 0, err
	}

	var data []Model
	for _, role := range _roles {
		d := Model{
			Role: role,
		}
		for _, permission := range permissions {
			if permission.RoleID == role.ID {
				permission.Role = nil
				d.Permissions = append(d.Permissions, permission)
			}
		}
		data = append(data, d)
	}

	return data, _count, nil
}

func (o *Service) UpdateIsActive(ctx context.Context, v models.Role) error {
	_, err := o.repository.UpdateOneRolesIsActive(ctx, v)
	if err != nil {
		return err
	}
	return nil
}

func (o *Service) Update(ctx context.Context, v Model) error {
	_role, err := o.repository.UpdateOneRoles(ctx, v.Role)
	if err != nil {
		return err
	}
	err = o.repository.DeleteManyRolePermissionsByRoleID(ctx, models.RolePermission{
		RoleID: _role.ID,
	})
	if err != nil {
		return err
	}

	if len(v.Permissions) > 0 {
		var permissions []models.Permission
		for _, permission := range v.Permissions {
			permissions = append(permissions, models.Permission{
				ID: permission.PermissionID,
			})
		}
		var value []models.RolePermission
		for _, v := range permissions {
			value = append(value, models.RolePermission{
				RoleID:       _role.ID,
				PermissionID: v.ID,
			})
		}
		_, err := o.repository.InsertManyRolePermissions(ctx, value)
		if err != nil {
			return err
		}

	}

	return nil
}

func (o *Service) AddPermission(ctx context.Context, v models.RolePermission) error {
	_, err := o.repository.InsertOneRolePermissions(ctx, v)
	if err != nil {
		return err
	}

	return nil
}
