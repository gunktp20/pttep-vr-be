package models

import "time"

type ActivityLog struct {
	ID          uint      `json:"id" bson:"_id" gorm:"column:id"`
	UserID      uint      `json:"user_id" bson:"user_id" gorm:"column:user_id"`
	BaseID      uint      `json:"base_id" bson:"base_id" gorm:"column:base_id"`
	SessionID   uint      `json:"session_id" bson:"session_id" gorm:"column:session_id"`
	QuestionID  uint      `json:"question_id" bson:"question_id" gorm:"column:question_id"`
	ChoiceID    uint      `json:"choice_id" bson:"choice_id" gorm:"column:choice_id"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}

func (ActivityLog) CollectionName() string {
	return "activity_logs"
}
