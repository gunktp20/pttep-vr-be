package models

import "time"

type Question struct {
	ID             uint `json:"id" bson:"_id" gorm:"column:id"`
	SessionID      uint `json:"session_id" bson:"session_id" gorm:"column:session_id"`
	QuestionTypeID uint `json:"question_type_id" bson:"question_type_id" gorm:"column:question_type_id"`

	Name           string `json:"name" bson:"name" gorm:"column:name"`
	Score          int64  `json:"score" bson:"score" gorm:"column:score"`
	MaxScore       int64  `json:"max_score" bson:"max_score" gorm:"column:max_score"`
	MinScore       int64  `json:"min_score" bson:"min_score" gorm:"column:min_score"`
	ScorePerChoice int64  `json:"score_per_choice" bson:"score_per_choice" gorm:"column:score_per_choice"`
	Time           string `json:"time" bson:"time" gorm:"column:time"`

	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (Question) TableName() string {
	return "questions"
}

func (Question) CollectionName() string {
	return "questions"
}
