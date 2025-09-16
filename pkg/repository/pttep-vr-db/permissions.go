package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"
)

func (o *Operator) FindPermissions(ctx context.Context, p *pagination.Pagination) ([]models.Permission, int64, error) {
	var mock models.Permission
	var results []models.Permission
	var count int64
	client := o.database.GetDB().Table(mock.TableName()).Where("is_remove = ?", false).Count(&count)
	if p != nil {
		client = client.Offset(int(p.Offset())).Limit(int(p.Limit())) //.Order(p.Order())
	}
	err := client.Find(&results).Error
	return results, count, err
}

func (o *Operator) FindPermissionsIn(ctx context.Context, v []models.Permission) ([]models.Permission, int64, error) {
	var mock models.Permission
	var results []models.Permission
	var permissionIds []uint
	for _, value := range v {
		permissionIds = append(permissionIds, value.ID)
	}

	var count int64
	err := o.database.GetDB().Table(mock.TableName()).Where("id IN (?)", permissionIds).Count(&count).Find(&results).Error
	return results, count, err
}

func (o *Operator) FindOnePermissions(ctx context.Context, id uint) (models.Permission, error) {
	var v models.Permission
	err := o.database.GetDB().Table(v.TableName()).Find(&v, " id = ?", id).Error
	return v, err
}

func (o *Operator) InsertOnePermissions(ctx context.Context, v models.Permission) (models.Permission, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (o *Operator) UpdateOnePermissions(ctx context.Context, v models.Permission) (models.Permission, error) {
	err := o.database.GetDB().Table(v.TableName()).Save(&v).Error
	return v, err
}

func (o *Operator) UpdateOnePermissionsName(ctx context.Context, v models.Permission) (models.Permission, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("name", v.Name).Error
	return v, err
}

func (o *Operator) UpdateOnePermissionsIsActive(ctx context.Context, v models.Permission) (models.Permission, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("is_active", v.IsActive).Error
	return v, err
}

func (o *Operator) UpdateOnePermissionsIsRemove(ctx context.Context, v models.Permission) (models.Permission, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Update("is_remove", v.IsRemove).Error
	return v, err
}
