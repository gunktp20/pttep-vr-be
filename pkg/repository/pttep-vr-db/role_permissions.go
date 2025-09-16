package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
)

func (o *Operator) FindRolePermissionsJoinPermissionsInRoleID(ctx context.Context, v []models.RolePermission) ([]models.RolePermission, int64, error) {
	var mock models.RolePermission
	var results []models.RolePermission
	var count int64
	var roleIds []uint
	for _, value := range v {
		roleIds = append(roleIds, value.RoleID)
	}
	client := o.database.GetDB().Table(mock.TableName()).Model(&results).Joins("left join permissions on permissions.id = role_permissions.permission_id")
	client = client.Joins("left join roles on roles.id = role_permissions.role_id").Preload("Permission").Preload("Role")
	err := client.Where("role_permissions.role_id IN (?) AND permissions.is_remove = ?", roleIds, false).Count(&count).Find(&results).Error
	return results, count, err
}

func (o *Operator) InsertOneRolePermissions(ctx context.Context, v models.RolePermission) (models.RolePermission, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (o *Operator) InsertManyRolePermissions(ctx context.Context, v []models.RolePermission) ([]models.RolePermission, error) {
	var mock models.RolePermission
	var temp []*models.RolePermission
	for _, value := range v {
		temp = append(temp, &value)
	}
	err := o.database.GetDB().Table(mock.TableName()).Create(temp).Error
	if err == nil {
		v = []models.RolePermission{}
		for _, value := range temp {
			v = append(v, *value)
		}
	}
	return v, err
}

func (o *Operator) DeleteOneRolePermissions(ctx context.Context, v models.RolePermission) error {
	err := o.database.GetDB().Table(v.TableName()).Delete(&v, "id = ?", v.ID).Error
	return err
}

func (o *Operator) DeleteOneRolePermissionsByRoleIDAndPermissionID(ctx context.Context, v models.RolePermission) error {
	err := o.database.GetDB().Table(v.TableName()).Delete(&v, "role_id = ? AND permission_id = ?", v.RoleID, v.PermissionID).Error
	return err
}

func (o *Operator) DeleteManyRolePermissionsByRoleID(ctx context.Context, v models.RolePermission) error {
	err := o.database.GetDB().Table(v.TableName()).Delete(&v, "role_id = ?", v.RoleID).Error
	return err
}
