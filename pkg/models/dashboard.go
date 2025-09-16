package models

type DashboardTotal struct {
	Total int64 `json:"total"`
}

type DashboardAvg struct {
	AvgScore float64 `json:"avg_score"`
}

type GameFailTransaction struct {
	QuestionID int    `json:"question_id"`
	Name       string `json:"name"`
	FalseCount int    `json:"false_count"`
}

type GraphTransaction struct {
	Date  string  `json:"date"`
	Score float64 `json:"score"`
}

type DashboardPassRate struct {
	PassRate float64 `json:"pass_rate"`
}
