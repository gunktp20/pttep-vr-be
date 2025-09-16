package models

import "time"

type UserTemp struct {
	ID       uint   `json:"id" bson:"id" gorm:"column:id"`
	Username string `json:"username" bson:"username" gorm:"column:username"`
	Name     string `json:"name" bson:"name" gorm:"column:name"`
	Surname  string `json:"surname" bson:"surname" gorm:"column:surname"`
	Email    string `json:"email" bson:"email" gorm:"column:email"`
	Tel      string `json:"tel" bson:"tel" gorm:"column:tel"`
	Group    string `json:"group" bson:"group" gorm:"column:group"`
	Company  string `json:"company" bson:"company" gorm:"column:company"`

	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (UserTemp) TableName() string {
	return "users_temp"
}

func (UserTemp) CollectionName() string {
	return "users_temp"
}
