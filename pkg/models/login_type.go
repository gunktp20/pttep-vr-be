package models

import "time"

const loginTypes = "login_types"

type LoginType struct {
	ID uint `gorm:"column:id" json:"id,omitempty" bson:"_id,omitempty"`

	Key         string    `json:"key" bson:"key" gorm:"colum:key"`
	Name        string    `json:"name" bson:"name" gorm:"column:name"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (LoginType) TableName() string {
	return loginTypes
}

func (LoginType) CollectionName() string {
	return loginTypes
}

func (_self *LoginType) SetID(v uint) {
	_self.ID = v
}
