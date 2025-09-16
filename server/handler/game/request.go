package game

type SettingGameReq struct {
	QuestId        uint   `json:"quest_id"`
	Score          int64  `json:"score" bson:"score" gorm:"column:score"`
	MaxScore       int64  `json:"max_score" bson:"max_score" gorm:"column:max_score"`
	MinScore       int64  `json:"min_score" bson:"min_score" gorm:"column:min_score"`
	ScorePerChoice int64  `json:"score_per_choice" bson:"score_per_choice" gorm:"column:score_per_choice"`
	Time           string `json:"time" bson:"time" gorm:"column:time"`
}

type Quest struct {
	QuestId        uint  `json:"quest_id" bson:"quest_id" gorm:"column:quest_id"`
	Score          int64 `json:"score" bson:"score" gorm:"column:score"`
	MaxScore       int64 `json:"max_score" bson:"max_score" gorm:"column:max_score"`
	MinScore       int64 `json:"min_score" bson:"min_score" gorm:"column:min_score"`
	ScorePerChoice int64 `json:"score_per_choice" bson:"score_per_choice" gorm:"column:score_per_choice"`
}

type GameIdReq struct {
	ID int `json:"id" bson:"_id" gorm:"column:_id"`
}

type GameTransactionReqTemp struct {
	QuestionName string  `json:"question_name" bson:"question_name" gorm:"column:question_name"`
	SessionName  string  `json:"session_name" bson:"session_name" gorm:"column:session_name"`
	BaseName     string  `json:"base_name" bson:"base_name" gorm:"column:base_name"`
	Start        string  `json:"start" bson:"start" gorm:"column:start"`
	Stop         string  `json:"stop" bson:"stop" gorm:"column:stop"`
	Time         int     `json:"time" bson:"time" gorm:"column:time"`
	Score        float64 `json:"score" bson:"score" gorm:"column:score"`
	PassPercent  float64 `json:"pass_percent" bson:"pass_percent" gorm:"column:pass_percent"`
	IsPass       bool    `json:"is_pass" bson:"is_pass" gorm:"column:is_pass"`
	Username     string  `json:"username" bson:"username" gorm:"column:username"`
}
