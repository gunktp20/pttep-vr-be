package models

import "time"

const permission = "permissions"

type Permission struct {
	ID uint `gorm:"column:id" json:"id,omitempty" bson:"_id,omitempty"`

	Key         string    `gorm:"column:key" json:"key" bson:"key"`
	Name        string    `gorm:"column:name" json:"name" bson:"name"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (Permission) TableName() string {
	return permission
}

func (Permission) CollectionName() string {
	return permission
}
func (_self *Permission) SetID(v uint) {
	_self.ID = v
}
