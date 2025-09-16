package models

import "time"

type Session struct {
	ID          uint      `json:"id" bson:"_id" gorm:"column:id"`
	CodeGame    string    `json:"code_game" bson:"code_game" gorm:"column:code_game"`
	BaseID      uint      `json:"base_id" bson:"base_id" gorm:"column:base_id"`
	Name        string    `json:"name" bson:"name" gorm:"column:name"`
	PassScore   float64   `json:"pass_score" bson:"pass_score" gorm:"column:pass_score"`
	PassPercent float64   `json:"pass_percent" bson:"pass_percent" gorm:"column:pass_percent"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (Session) TableName() string {
	return "sessions"
}

func (Session) CollectionName() string {
	return "sessions"
}
