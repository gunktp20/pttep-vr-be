package pttep_vr_db

import "pttep-vr-api/pkg/models"

func (o *Operator) InsertOneUsersTemp(v models.UserTemp) (models.UserTemp, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (o *Operator) FindOneUsersTempByUsername(s string) (models.UserTemp, error) {
	var v models.UserTemp
	err := o.database.GetDB().Table(v.TableName()).Where("username = ?", s).Find(&v).Error
	return v, err
}
