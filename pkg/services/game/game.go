package game

import (
	"context"
	"log"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/server/response"
	"time"
)

func (s *Service) CreateGameTransactionPlayer(ctx context.Context, req models.GameTransactionReq) error {
	for _, question := range req.QuestionReq {
		game := models.GameTransaction{
			QuestionId:  question.QuestionId,
			SessionId:   req.SessionId,
			Start:       req.Start,
			Stop:        req.Stop,
			Time:        req.Time,
			PassPercent: req.PassPercent,
			IsPass:      question.IsPass,
			Score:       question.Score,
			CreatedDate: time.Now(),
			Username:    req.Username,
		}
		_, err := s.repository.InsertOneGameTransaction(game)
		if err != nil {
			log.Print("Error InsertOneGameTransaction", err)
			return err
		}
	}
	return nil
}

func (s *Service) CreateGameTransactionPlayerTemp(ctx context.Context, game models.GameTransactionTemp) (response.Response, error) {

	_, err := s.repository.InsertOneGameTransactionTemp(game)
	if err != nil {
		//log.Print("Error InsertOneGameTransaction Temp", err)
		return response.Response{
			Code:    "101",
			Message: "Please try again later.",
		}, nil
	}

	return response.Response{
		Code:    "0",
		Message: "Success",
	}, nil
}

func (s *Service) GetSettingGameByUser(ctx context.Context, key string) ([]models.SessionWithDetails, error) {
	return s.repository.FindSettingGameUser(key)
}

func (s *Service) GetGameList(ctx context.Context) ([]models.SessionList, error) {
	return s.repository.FindSettingGameList()
}

func (s *Service) GetSettingGameById(ctx context.Context, id uint) ([]models.QuestionList, error) {
	return s.repository.FindSettingGameById(id)
}

func (s *Service) UpdateSettingGameById(ctx context.Context, req models.Question) (models.Question, error) {
	return s.repository.UpdateSettingGameById(req)
}

func (s *Service) ReportGame(ctx context.Context, req models.ReportExportReq) ([]models.GameTransactionResponse, int, error) {
	username := req.Username
	score := req.Score
	status := req.Status
	baseName := req.BaseName
	sessionName := req.Session
	email := req.Email
	startDate := req.StartDate
	endDate := req.EndDate
	var isPass bool

	if status == "PASS" {
		isPass = true
	} else if status == "FAIL" {
		isPass = false
	}

	results, count, err := s.repository.FindGameTransactions(username, score, isPass, baseName, sessionName, status, email, startDate, endDate, req.Limit, req.Page)
	if err != nil {
		log.Print("Error FindGameTransactions:", err)
		return []models.GameTransactionResponse{}, 0, err
	}

	return results, count, nil
}

func (s *Service) DashBoardSumTime(ctx context.Context, req models.DashboardReq) (models.DashboardTotal, error) {
	return s.repository.SumTimeTransGmaeUser(req.StartDate, req.EndDate)
}

func (s *Service) GetDistinctUsernameCountByDateRange(ctx context.Context, req models.DashboardReq) (models.DashboardTotal, error) {
	results, err := s.repository.GetDistinctUsernameCountByDateRange(req.StartDate, req.EndDate)
	if err != nil {
		log.Print("Error GetDistinctUsernameCountByDateRange", err)
		return models.DashboardTotal{}, err
	}
	res := models.DashboardTotal{
		Total: results,
	}

	return res, nil
}

func (s *Service) GetAvgScore(ctx context.Context, req models.DashboardReq) (models.DashboardAvg, error) {
	results, err := s.repository.GetAvgOfMaxAverageScore(req.StartDate, req.EndDate)
	if err != nil {
		log.Print("Error GetAvgOfMaxAverageScore", err)
		return models.DashboardAvg{}, err
	}
	res := models.DashboardAvg{
		AvgScore: results,
	}

	return res, nil
}

func (s *Service) GetMaxFalseQuestion(ctx context.Context, req models.DashboardReq) ([]models.GameFailTransaction, error) {
	return s.repository.GetMaxFalseIsPassByQuestion(req.StartDate, req.EndDate)
}

func (s *Service) GetGraphScoreSevenDay(ctx context.Context, req models.BaseReq) ([]models.GraphTransaction, error) {
	return s.repository.GetGraphScoreSevenDay(req.Base)
}

func (s *Service) GetPassRateByDate(ctx context.Context, req models.DashboardReq) (models.DashboardPassRate, error) {
	return s.repository.GetPassRateByDate(req.StartDate, req.EndDate)
}

func (s *Service) GetMostQuestionByDate(ctx context.Context, req models.DashboardReq) ([]models.QuestionPlayCount, error) {
	return s.repository.GetMostQuestionByDate(req.StartDate, req.EndDate)
}

func (s *Service) GetTopQuestionsByTimeByDate(ctx context.Context, req models.DashboardReq) ([]models.QuestionTimeStats, error) {
	return s.repository.GetTopQuestionsByTime(req.StartDate, req.EndDate)
}

func (s *Service) GetTotalScoreForUsersbyDate(ctx context.Context, req models.DashboardReq) ([]models.UserTotalScore, error) {
	return s.repository.GetTotalScoreForUsersWithDateRange(req.StartDate, req.EndDate)
}

func (s *Service) GetQuestionPlayCount(ctx context.Context, req models.DashboardReq) ([]models.QuestionPlayCount, error) {
	return s.repository.GetTopQuestionsByPlayCount(req.StartDate, req.EndDate)
}
