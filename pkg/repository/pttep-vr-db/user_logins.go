package pttep_vr_db

import (
	"context"
	"pttep-vr-api/pkg/models"
)

func (o *Operator) FindUserLogins(ctx context.Context) ([]models.UserLogin, error) {
	var mock models.UserLogin
	var results []models.UserLogin
	err := o.database.GetDB().Table(mock.TableName()).Find(&results).Error
	return results, err
}
func (o *Operator) FindOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	err := o.database.GetDB().Table(v.TableName()).Find(&v).Error
	return v, err
}

func (o *Operator) FindOneUserLoginsByUsername(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("username LIKE ? AND login_type_id = ?", v.Username, v.LoginTypeID).Find(&v).Error
	return v, err
}

func (o *Operator) FindOneUserLoginsByUsernameAndPassword(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("username LIKE ? AND password LIKE ? AND login_type_id = ?", v.Username, v.Password, v.LoginTypeID).Find(&v).Error
	return v, err
}

func (o *Operator) InsertOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (o *Operator) UpdateOneUserLogins(ctx context.Context, v models.UserLogin) (models.UserLogin, error) {
	update := make(map[string]interface{})
	update["username"] = v.Username
	if v.Password != "" {
		update["password"] = v.Password
	}
	err := o.database.GetDB().Table(v.TableName()).Where("user_id = ? AND login_type_id = ?", v.UserID, v.LoginTypeID).Updates(update).Error
	return v, err
}
