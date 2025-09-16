package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"
	"time"
)

func (o *Operator) FindUserRoles(ctx context.Context) ([]models.UserRole, error) {
	var v []models.UserRole
	var mock models.UserRole
	err := o.database.GetDB().Table(mock.TableName()).Find(&v).Error
	return v, err
}

func (o *Operator) FindUserRolesByUserID(ctx context.Context, v models.UserRole) ([]models.UserRole, error) {
	var result []models.UserRole
	err := o.database.GetDB().Table(v.TableName()).Where("user_id = ?", v.UserID).Find(&result).Error
	return result, err
}

func (o *Operator) FindOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Find(&v).Error
	return v, err
}

func (o *Operator) InsertOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (o *Operator) UpdateOneUserRolesRole(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("role_id", v.RoleID).Error
	return v, err
}

func (o *Operator) UpdateOneUserRoles(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	err := o.database.GetDB().Table(v.TableName()).Save(&v).Error
	return v, err
}

func (o *Operator) UpdateUserRolesByEmail(ctx context.Context, v models.UserRole) error {
	update := make(map[string]interface{})
	update["updated_date"] = time.Now()
	update["user_id"] = v.UserID
	return o.database.GetDB().Table(v.TableName()).Where(" user_id = ? AND email LIKE ?", 0, v.Email).Updates(update).Error
}

func (o *Operator) UpdateOneUserRolesIsActive(ctx context.Context, v models.UserRole) error {
	return o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("is_active", v.IsActive).Error
}
func (o *Operator) UpdateOneUserRolesIsRemove(ctx context.Context, v models.UserRole) error {
	return o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("is_remove", v.IsRemove).Error
}

func (o *Operator) FindUserRolesLeftJoinUserAndRolePermissionAndRole(ctx context.Context, p *pagination.Pagination) ([]models.UserRole, int64, error) {
	var mock models.UserRole
	var v []models.UserRole
	var count int64
	client := o.database.GetDB().Table(mock.TableName()).Model(&mock)
	client = client.Joins("left join users on users.id = user_roles.user_id")
	client = client.Joins("left join roles on roles.id = user_roles.role_id")
	client = client.Preload("User").Preload("Role")
	client = client.Where("user_roles.is_remove = ? AND roles.is_remove = ? AND users.is_remove = ? ", false, false, false).Count(&count)
	if p != nil {
		client = client.Offset(int(p.Offset())).Limit(int(p.Limit())) //.Order(p.Order())
	}
	err := client.Find(&v).Error
	return v, count, err
}

func (o *Operator) FindOneUserRolesLeftJoinUserAndRolePermissionAndRole(ctx context.Context, v models.UserRole) (models.UserRole, error) {
	client := o.database.GetDB().Table(v.TableName()).Model(&v)
	client = client.Joins("left join users on users.id = user_roles.user_id")
	client = client.Joins("left join roles on roles.id = user_roles.role_id")
	client = client.Preload("User").Preload("Role")
	err := client.Where("user_roles.id = ?", v.ID).Find(&v).Error
	return v, err
}
