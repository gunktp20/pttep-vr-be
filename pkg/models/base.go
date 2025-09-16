package models

import "time"

type Base struct {
	ID          uint      `json:"id" bson:"_id" gorm:"column:id"`
	Name        string    `json:"name" bson:"name" gorm:"column:name"`
	Description string    `json:"description" bson:"description" gorm:"column:description"`
	Area        string    `json:"area" bson:"area" gorm:"column:area"`
	Location    string    `json:"location" bson:"location" gorm:"column:location"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (Base) TableName() string {
	return "bases"
}

func (Base) CollectionName() string {
	return "bases"
}
