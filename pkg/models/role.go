package models

import "time"

const role = "roles"

type Role struct {
	ID uint `gorm:"column:id" json:"id,omitempty" bson:"_id,omitempty"`

	Key         string    `gorm:"column:key" json:"key" bson:"key"`
	Name        string    `gorm:"column:name" json:"name" bson:"name"`
	Description string    `gorm:"column:description" json:"description" bson:"description"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	IsDefault   bool      `json:"is_default" bson:"is_default" gorm:"column:is_default"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (Role) TableName() string {
	return role
}

func (Role) CollectionName() string {
	return role
}

func (_self *Role) SetID(v uint) {
	_self.ID = v
}
