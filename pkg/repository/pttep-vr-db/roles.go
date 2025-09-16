package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"
)

func (o *Operator) FindRoles(ctx context.Context, p *pagination.Pagination) ([]models.Role, int64, error) {
	var mock models.Role
	var results []models.Role
	var count int64
	client := o.database.GetDB().Table(mock.TableName()).Where("is_remove = ?", false).Count(&count)
	if p != nil {
		client = client.Offset(int(p.Offset())).Limit(int(p.Limit())) //.Order(p.Order())
	}
	err := client.Find(&results).Error
	return results, count, err
}

func (o *Operator) FindOneRoles(ctx context.Context, id uint) (models.Role, error) {
	var v models.Role
	err := o.database.GetDB().Table(v.TableName()).Find(&v, " id = ?", id).Error
	return v, err
}

func (o *Operator) FindOneRolesByIsDefault(ctx context.Context, status bool) (models.Role, error) {
	var v models.Role
	err := o.database.GetDB().Table(v.TableName()).Find(&v, " is_default = ?", status).Error
	return v, err
}

func (o *Operator) InsertOneRoles(ctx context.Context, v models.Role) (models.Role, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (o *Operator) UpdateOneRoles(ctx context.Context, v models.Role) (models.Role, error) {
	err := o.database.GetDB().Table(v.TableName()).Save(&v).Error
	return v, err
}

func (o *Operator) UpdateOneRolesName(ctx context.Context, v models.Role) (models.Role, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("name", v.Name).Error
	return v, err
}

func (o *Operator) UpdateOneRolesIsActive(ctx context.Context, v models.Role) (models.Role, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("is_active", v.IsActive).Error
	return v, err
}

func (o *Operator) UpdateOneRolesIsRemove(ctx context.Context, v models.Role) (models.Role, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("is_remove", v.IsRemove).Error
	return v, err
}
