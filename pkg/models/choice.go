package models

import "time"

type Choice struct {
	ID          uint      `json:"id" bson:"_id" gorm:"column:id"`
	QuestionID  uint      `json:"question_id" bson:"question_id" gorm:"column:question_id"`
	Name        string    `json:"name" bson:"name" gorm:"column:name"`
	Value       int       `json:"value" bson:"value" gorm:"column:value"`
	Score       int       `json:"score" bson:"score" gorm:"column:score"`
	IsAnswer    bool      `json:"is_answer" bson:"is_answer" gorm:"column:is_answer"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (Choice) TableName() string {
	return "choices"
}

func (Choice) CollectionName() string {
	return "choices"
}
