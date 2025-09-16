package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
)

func (o *Operator) FindLoginTypes(ctx context.Context) ([]models.LoginType, error) {
	var mock models.LoginType
	var results []models.LoginType
	err := o.database.GetDB().Table(mock.TableName()).Find(&results).Error
	return results, err
}

func (o *Operator) FindOneLoginTypes(ctx context.Context, v models.LoginType) (models.LoginType, error) {
	err := o.database.GetDB().Table(v.TableName()).Find(&v).Error
	return v, err
}

func (o *Operator) InsertOneLoginTypes(ctx context.Context, v models.LoginType) (models.LoginType, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}
