package models

type GameTransactionReq struct {
	QuestionReq []QuestionReq `json:"question_req" bson:"question_req" gorm:"column:question_req"`
	SessionId   uint          `json:"session_id" bson:"session_id" gorm:"column:session_id"`
	Start       string        `json:"start" bson:"start" gorm:"column:start"`
	Stop        string        `json:"stop" bson:"stop" gorm:"column:stop"`
	Time        int           `json:"time" bson:"time" gorm:"column:time"`
	Score       float64       `json:"score" bson:"score" gorm:"column:score"`
	PassPercent float64       `json:"pass_percent" bson:"pass_percent" gorm:"column:pass_percent"`
	IsPass      bool          `json:"is_pass" bson:"is_pass" gorm:"column:is_pass"`
	Username    string        `json:"username" bson:"username" gorm:"column:username"`
}

type QuestionReq struct {
	QuestionId uint    `json:"question_id" bson:"question_id" gorm:"column:question_id"`
	IsPass     bool    `json:"is_pass" bson:"is_pass" gorm:"column:is_pass"`
	Score      float64 `json:"score" bson:"score" gorm:"column:score"`
}
