package models

import "time"

type GameTransaction struct {
	ID          uint      `json:"id" bson:"_id" gorm:"column:id"`
	QuestionId  uint      `json:"question_id" bson:"question_id" gorm:"column:question_id"`
	SessionId   uint      `json:"session_id" bson:"session_id" gorm:"column:session_id"`
	Start       string    `json:"start" bson:"start" gorm:"column:start"`
	Stop        string    `json:"stop" bson:"stop" gorm:"column:stop"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
	Time        int       `json:"time" bson:"time" gorm:"column:time"`
	Score       float64   `json:"score" bson:"score" gorm:"column:score"`
	PassPercent float64   `json:"pass_percent" bson:"pass_percent" gorm:"column:pass_percent"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsPass      bool      `json:"is_pass" bson:"is_pass" gorm:"column:is_pass"`
	Username    string    `json:"username" bson:"username" gorm:"column:username"`
}

func (GameTransaction) TableName() string {
	return "game_transactions"
}

func (GameTransaction) CollectionName() string {
	return "game_transactions"
}
