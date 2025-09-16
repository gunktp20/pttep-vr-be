package pttep_vr_db

import (
	"pttep-vr-api/pkg/models"
	"time"
)

func (o *Operator) FindUsers() ([]models.User, error) {
	var mock models.User
	var results []models.User
	err := o.database.GetDB().Table(mock.TableName()).Find(&results).Error
	return results, err
}
func (o *Operator) FindOneUsers(id uint) (models.User, error) {
	var v models.User
	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", id).Find(&v).Error
	return v, err
}
func (o *Operator) FindOneUsersByEmail(v models.User) (models.User, error) {
	err := o.database.GetDB().Table(v.TableName()).Where("email = ?", v.Email).Find(&v).Error
	return v, err
}
func (o *Operator) InsertOneUsers(v models.User) (models.User, error) {
	err := o.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}
func (o *Operator) UpdateOneUsers(v models.User) (models.User, error) {
	update := make(map[string]interface{})
	update["email"] = v.Email
	update["name"] = v.Name
	update["surname"] = v.Surname
	update["tel"] = v.Tel
	update["position"] = v.Position
	update["company"] = v.Company
	update["working_with"] = v.WorkingWith
	update["work_for_department"] = v.WorkForDepartment
	update["location"] = v.Location
	update["code"] = v.Code
	update["code_expired_at"] = v.CodeExpiredAt
	v.UpdatedDate = time.Now()
	update["updated_date"] = true

	err := o.database.GetDB().Table(v.TableName()).Where("id = ?", v.ID).Updates(update).Error
	return v, err
}

func (o *Operator) FindOneUserByCode(code string) (models.User, error) {
	var v models.User
	err := o.database.GetDB().Table(v.TableName()).Where("code = ?", code).Find(&v).Error
	return v, err
}
