package models

type SessionWithDetails struct {
	SessionID        uint    `json:"session_id"`
	BaseName         string  `json:"base_name"`
	SessionName      string  `json:"session_name"`
	CodeGame         string  `json:"code_game"`
	QuestionID       uint    `json:"question_id"`
	QuestionName     string  `json:"question_name"`
	QuestionScore    float64 `json:"question_score"`
	ScorePerChoice   int     `json:"score_per_choice"`
	QuestionTypeName string  `json:"question_type_name"`
	QuestionOrder    int     `json:"question_order"`
}

type SessionList struct {
	SessionID   uint   `json:"session_id"`
	BaseID      uint   `json:"base_id"`
	CodeGame    string `json:"code_game"`
	SessionName string `json:"session_name"`
	BaseName    string `json:"base_name"`
}

type QuestionList struct {
	QuestId          uint   `json:"quest_id"`
	QuestName        string `json:"quest_name"`
	Score            int64  `json:"score" bson:"score" gorm:"column:score"`
	MaxScore         int64  `json:"max_score" bson:"max_score" gorm:"column:max_score"`
	MinScore         int64  `json:"min_score" bson:"min_score" gorm:"column:min_score"`
	ScorePerChoice   int64  `json:"score_per_choice" bson:"score_per_choice" gorm:"column:score_per_choice"`
	Time             string `json:"time" bson:"time" gorm:"column:time"`
	QuestionTypeName string `json:"question_type_name"`
}

type ReportReq struct {
	QuestId   uint   `json:"quest_id"`
	SessionID uint   `json:"session_id"`
	Score     int64  `json:"score" bson:"score" gorm:"column:score"`
	Username  string `json:"username" bson:"username" gorm:"column:username"`
	Email     string `json:"email" bson:"email" gorm:"column:email"`
	BaseName  string `json:"baseName" bson:"baseName" gorm:"column:email"`
	Status    string `json:"status" bson:"status" gorm:"column:status"`
}

type GameTransactionList struct {
	SessionID   uint   `json:"session_id"`
	BaseID      uint   `json:"base_id"`
	CodeGame    string `json:"code_game"`
	SessionName string `json:"session_name"`
	BaseName    string `json:"base_name"`
}

type DashboardReq struct {
	StartDate string `json:"start_date" bson:"start_date" gorm:"column:start_date"`
	EndDate   string `json:"end_date" bson:"end_date" gorm:"column:end_date"`
}

type BaseReq struct {
	Base string `json:"base" bson:"base" gorm:"column:base"`
}

type QuestionPlayCount struct {
	QuestionName string `json:"question_name"`
	PlayCount    int    `json:"play_count"`
}

type QuestionTimeStats struct {
	QuestionName string  `json:"question_name"`
	TotalTime    float64 `json:"total_time"`
}

type UserTotalScore struct {
	Username   string  `json:"username"`
	TotalScore float64 `json:"total_score"`
}

type ReportExportReq struct {
	Username  string  `json:"username" bson:"username" gorm:"column:username"`
	Email     string  `json:"email" bson:"email" gorm:"column:email"`
	BaseName  string  `json:"base_name" bson:"base_name" gorm:"column:base_name"`
	Session   string  `json:"session" bson:"session" gorm:"column:session"`
	Score     float64 `json:"score" bson:"score" gorm:"column:score"`
	Status    string  `json:"status" bson:"status" gorm:"column:status"`
	StartDate string  `json:"start_date" bson:"start_date" gorm:"column:start_date"`
	EndDate   string  `json:"end_date" bson:"end_date" gorm:"column:end_date"`
	Limit     int     `json:"limit" bson:"limit" gorm:"column:limit"`
	Page      int     `json:"page" bson:"page" gorm:"column:page"`
}

type GameTransactionResponse struct {
	Username    string  `json:"username"`
	Score       float64 `json:"score"`
	IsPass      string  `json:"is_pass"`
	BaseName    string  `json:"base_name"`
	SessionName string  `json:"session_name"`
	QuestName   string  `json:"quest_name"`
	Code        string  `json:"code"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	CreatedDate string  `json:"created_date"`
}
