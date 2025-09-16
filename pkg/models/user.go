package models

import "time"

type User struct {
	ID                uint       `json:"id" bson:"id" gorm:"column:id"`
	Code              string     `json:"code" bson:"code" gorm:"column:code"`
	CodeExpiredAt     *time.Time `json:"code_expired_at" gorm:"column:code_expired_at"`
	Email             string     `json:"email" bson:"email" gorm:"column:email"`
	Name              string     `json:"name" bson:"name" gorm:"column:name"`
	Surname           string     `json:"surname" bson:"surname" gorm:"column:surname"`
	Tel               string     `json:"tel" bson:"tel" gorm:"column:tel"`
	Position          string     `json:"position" bson:"position" gorm:"column:position"`
	Company           string     `json:"company" bson:"company" gorm:"column:company"`
	WorkingWith       string     `json:"working_with" bson:"working_with" gorm:"column:working_with"`
	WorkForDepartment string     `json:"work_for_department" bson:"work_for_department" gorm:"column:work_for_department"`
	Location          string     `json:"location" bson:"location" gorm:"column:location"`
	LastIpLogin       string     `json:"last_ip_login" bson:"last_ip_login" gorm:"column:last_ip_login"`
	LastLoginDate     time.Time  `json:"last_login_date" bson:"last_login_date" gorm:"column:last_login_date"`
	IsActive          bool       `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove          bool       `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	LoginTypeID       uint       `json:"login_type_id" bson:"login_type_id" gorm:"column:login_type_id"`
	CreatedDate       time.Time  `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate       time.Time  `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (User) TableName() string {
	return "users"
}

func (User) CollectionName() string {
	return "users"
}
