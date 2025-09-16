package pttep_vr_db

import (
	"pttep-vr-api/pkg/models"
)

func (c *Operator) InsertOneGameTransaction(v models.GameTransaction) (models.GameTransaction, error) {
	err := c.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (c *Operator) FindSettingGameList() ([]models.SessionList, error) {

	var result []models.SessionList
	var v models.Session
	err := c.database.GetDB().Table(v.TableName() + " AS s").
		Select("s.id AS session_id, s.base_id AS base_id, s.code_game, s.name AS session_name, b.name AS base_name").
		Joins("LEFT JOIN bases b ON s.base_id = b.id").
		Find(&result).Error

	if err != nil {
		return []models.SessionList{}, err
	}

	return result, nil
}

func (c *Operator) FindSettingGameById(id uint) ([]models.QuestionList, error) {

	var result []models.QuestionList
	var v models.Question
	err := c.database.GetDB().Table(v.TableName()+" AS q").Where("session_id = ?", id).
		Order("q.order ASC").
		Select("q.id AS quest_id, q.name AS quest_name, q.score, q.min_score, q.max_score, q.score_per_choice, q.time, q.order, qt.name AS question_type_name").
		Joins("LEFT JOIN question_types qt ON qt.id = q.question_type_id").
		Find(&result).Error

	if err != nil {
		return []models.QuestionList{}, err
	}

	return result, nil
}

func (c *Operator) UpdateSettingGameById(update models.Question) (models.Question, error) {

	var v models.Question
	// Perform the update
	err := c.database.GetDB().Table(v.TableName()).
		Where("id = ?", update.ID).
		UpdateColumns(map[string]interface{}{
			"score":            update.Score,
			"min_score":        update.MinScore,
			"max_score":        update.MaxScore,
			"score_per_choice": update.ScorePerChoice,
			"time":             update.Time,
		}).Error

	if err != nil {
		return models.Question{}, err
	}

	return models.Question{}, nil
}

func (c *Operator) InsertOneSettingGame(v models.Session) (models.Session, error) {
	err := c.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (c *Operator) InsertOneSettingGameQuest(v models.Question) (models.Question, error) {
	err := c.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (c *Operator) InsertOneGameTransactionTemp(v models.GameTransactionTemp) (models.GameTransactionTemp, error) {
	err := c.database.GetDB().Table(v.TableName()).Create(&v).Error
	return v, err
}

func (c *Operator) FindSettingGameUser(key string) ([]models.SessionWithDetails, error) {

	var result []models.SessionWithDetails
	var v models.Session

	err := c.database.GetDB().Table(v.TableName()+" AS s").
		Select("s.id AS session_id, b.name AS base_name, s.name AS session_name, s.code_game, q.id AS question_id, q.name AS question_name, q.score AS question_score, q.score_per_choice, qt.name AS question_type_name, q.order AS question_order").
		Joins("LEFT JOIN bases b ON s.base_id = b.id").
		Joins("LEFT JOIN questions q ON s.id = q.session_id").
		Joins("LEFT JOIN question_types qt ON qt.id = q.question_type_id").
		Where("s.code_game = ?", key).
		Order("q.order ASC").
		Scan(&result).Error

	if err != nil {
		return []models.SessionWithDetails{}, err
	}

	return result, nil

}

func (c *Operator) FindRportTransGmaeUser() ([]models.GameTransaction, error) {

	var result []models.GameTransaction
	var v models.Session
	err := c.database.GetDB().Table(v.TableName() + " AS s").
		Select("s.id AS session_id, s.base_id AS base_id, s.code_game, s.name AS session_name, b.name AS base_name").
		Joins("LEFT JOIN bases b ON s.base_id = b.id").
		Find(&result).Error

	if err != nil {
		return []models.GameTransaction{}, err
	}

	return result, nil
}

func (c *Operator) FindDashboardTransGmaeUser() ([]models.GameTransaction, error) {

	var result []models.GameTransaction
	var v models.Session
	err := c.database.GetDB().Table(v.TableName() + " AS s").
		Select("s.id AS session_id, s.base_id AS base_id, s.code_game, s.name AS session_name, b.name AS base_name").
		Joins("LEFT JOIN bases b ON s.base_id = b.id").
		Find(&result).Error

	if err != nil {
		return []models.GameTransaction{}, err
	}

	return result, nil
}

func (c *Operator) SumTimeTransGmaeUser(startDate, endDate string) (models.DashboardTotal, error) {

	var result models.DashboardTotal
	var v models.GameTransaction
	err := c.database.GetDB().Table(v.TableName()).
		Where("created_date >= ? AND created_date < ?", startDate, endDate).
		Select("SUM(time) AS total").
		Find(&result).Error

	if err != nil {
		return models.DashboardTotal{}, err
	}

	return result, nil
}

func (c *Operator) GetDistinctUsernameCountByDateRange(startDate, endDate string) (int64, error) {
	var count int64
	var v models.GameTransaction
	err := c.database.GetDB().Table(v.TableName()).
		Where("created_date >= ? AND created_date < ?", startDate, endDate).
		Select("COUNT(DISTINCT username)").
		Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *Operator) GetAvgOfMaxAverageScore(startDate, endDate string) (float64, error) {
	var result models.DashboardAvg
	var gt models.GameTransaction

	subQuery := c.database.GetDB().Table(gt.TableName()).
		Select("game_transactions.username, MAX(average_score) AS max_average_score").
		Joins("JOIN (SELECT username, question_id, AVG(score) AS average_score FROM game_transactions WHERE created_date >= ? AND created_date < ? GROUP BY username, question_id) AS subquery ON subquery.username = game_transactions.username", startDate, endDate).
		Group("game_transactions.username")

	err := c.database.GetDB().Table(gt.TableName()).
		Select("SUM(final_query.max_average_score) / COUNT(final_query.username) AS avg_score").
		Joins("JOIN (?) AS final_query ON final_query.username = game_transactions.username", subQuery).
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	return result.AvgScore, nil
}

func (c *Operator) GetMaxFalseIsPassByQuestion(startDate, endDate string) ([]models.GameFailTransaction, error) {
	var result []models.GameFailTransaction
	var v models.GameTransaction

	err := c.database.GetDB().Table(v.TableName()+" AS gt").
		Select("gt.question_id, q.name, COUNT(CASE WHEN gt.is_pass = false THEN 1 END) AS false_count").
		Joins("LEFT JOIN questions q ON gt.question_id = q.id").
		Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate).
		Group("gt.question_id, q.name").
		Order("false_count DESC").
		Limit(1).
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Operator) GetGraphScoreSevenDay(base string) ([]models.GraphTransaction, error) {
	var result []models.GraphTransaction
	var v models.GameTransaction

	err := c.database.GetDB().Table(v.TableName()+" AS gt").
		Select("DATE(gt.created_date) AS date, ROUND(SUM(gt.score), 2) AS score").
		Joins("LEFT JOIN sessions s ON s.id = gt.session_id").
		Joins("LEFT JOIN bases b ON b.id = s.base_id").
		Where("b.name = ?", base).
		Where("gt.created_date >= CURDATE() - INTERVAL 7 DAY").
		Group("DATE(gt.created_date)").
		Order("DATE(gt.created_date) ASC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Operator) GetPassRateByDate(startDate, endDate string) (models.DashboardPassRate, error) {
	var passRate models.DashboardPassRate
	var v models.GameTransaction

	err := c.database.GetDB().Model(&v).
		Select("COUNT(CASE WHEN is_pass = TRUE THEN 1 END) / COUNT(*) * 100 AS pass_rate").
		Where("created_date >= ? AND created_date < ?", startDate, endDate).
		Scan(&passRate).Error

	if err != nil {
		return passRate, err
	}

	return passRate, nil
}

func (c *Operator) GetMostQuestionByDate(startDate, endDate string) ([]models.QuestionPlayCount, error) {
	var result []models.QuestionPlayCount
	var gt models.GameTransaction
	var q models.Question
	var s models.Session
	var qt models.QuestionType

	err := c.database.GetDB().Table(gt.TableName()+" AS gt").
		Select("q.name AS question_name, COUNT(gt.id) AS play_count").
		Joins("LEFT JOIN "+q.TableName()+" AS q ON gt.question_id = q.id").
		Joins("LEFT JOIN "+s.TableName()+" AS s ON gt.session_id = s.id").
		Joins("LEFT JOIN "+qt.TableName()+" AS qt ON q.question_type_id = qt.id").
		Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate).
		Group("q.name").
		Order("play_count DESC").
		Limit(10).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *Operator) GetTopQuestionsByTime(startDate, endDate string) ([]models.QuestionTimeStats, error) {
	var result []models.QuestionTimeStats
	var gt models.GameTransaction
	var q models.Question
	var s models.Session
	var qt models.QuestionType

	err := c.database.GetDB().Table(gt.TableName()+" AS gt").
		Select("q.name AS question_name, SUM(gt.time) AS total_time").
		Joins("LEFT JOIN "+q.TableName()+" AS q ON gt.question_id = q.id").
		Joins("LEFT JOIN "+s.TableName()+" AS s ON gt.session_id = s.id").
		Joins("LEFT JOIN "+qt.TableName()+" AS qt ON q.question_type_id = qt.id").
		Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate).
		Group("q.name").
		Order("total_time DESC").
		Limit(10).
		Scan(&result).
		Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Operator) GetTotalScoreForUsersWithDateRange(startDate, endDate string) ([]models.UserTotalScore, error) {
	var result []models.UserTotalScore

	subquery := c.database.GetDB().Table("game_transactions gt").
		Select("gt.username, MAX(gt.score) AS max_score").
		Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate).
		Group("gt.username, gt.question_id")

	err := c.database.GetDB().Table("(?) AS max_scores", subquery).
		Select("u.name AS username, SUM(max_scores.max_score) AS total_score").
		Joins("LEFT JOIN users u ON u.code = max_scores.username").
		Group("u.name").
		Order("total_score DESC").
		Limit(10).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (c *Operator) GetTopQuestionsByPlayCount(startDate, endDate string) ([]models.QuestionPlayCount, error) {
	var result []models.QuestionPlayCount
	var gt models.GameTransaction
	var q models.Question
	var s models.Session
	var qt models.QuestionType

	err := c.database.GetDB().Table(gt.TableName()+" AS gt").
		Select("q.name AS question_name, COUNT(gt.question_id) AS play_count").
		Joins("LEFT JOIN "+q.TableName()+" AS q ON gt.question_id = q.id").
		Joins("LEFT JOIN "+s.TableName()+" AS s ON gt.session_id = s.id").
		Joins("LEFT JOIN "+qt.TableName()+" AS qt ON q.question_type_id = qt.id").
		Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate).
		Group("q.name").
		Order("play_count DESC").
		Limit(3).
		Scan(&result).
		Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Operator) FindGameTransactions(username string, score float64, isPass bool, baseName string, sessionName string, status string, email string, startDate string, endDate string, limit int, page int) ([]models.GameTransactionResponse, int, error) {
	var result []models.GameTransactionResponse
	var gt models.GameTransaction
	var totalCount int64

	query := c.database.GetDB().Table(gt.TableName() + " AS gt").
		Select("gt.username, gt.score, gt.is_pass, gt.created_date, b.name AS base_name, s.name AS session_name, q.name AS quest_name, u.email, u.code, u.name").
		Joins("LEFT JOIN sessions s ON gt.session_id = s.id").
		Joins("LEFT JOIN bases b ON s.base_id = b.id").
		Joins("LEFT JOIN questions q ON gt.question_id = q.id").
		Joins("LEFT JOIN question_types qt ON q.question_type_id = qt.id").
		Joins("LEFT JOIN users u ON gt.username = u.code").
		Order("gt.created_date DESC")

	if username != "" {
		query = query.Where("gt.username = ?", username)
	}
	if score != 0 {
		query = query.Where("gt.score = ?", score)
	}
	if status != "ALL" {
		query = query.Where("gt.is_pass = ?", isPass)
	}
	if baseName != "" {
		query = query.Where("b.name = ?", baseName)
	}
	if sessionName != "" {
		query = query.Where("s.name = ?", sessionName)
	}
	if email != "" {
		query = query.Where("u.email = ?", email)
	}
	if startDate != "" && endDate != "" {
		query = query.Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate)
	}

	var countQuery = c.database.GetDB().Table(gt.TableName() + " AS gt").
		Joins("LEFT JOIN sessions s ON gt.session_id = s.id").
		Joins("LEFT JOIN bases b ON s.base_id = b.id").
		Joins("LEFT JOIN questions q ON gt.question_id = q.id").
		Joins("LEFT JOIN question_types qt ON q.question_type_id = qt.id").
		Joins("LEFT JOIN users u ON gt.username = u.code")

	if username != "" {
		countQuery = countQuery.Where("gt.username = ?", username)
	}
	if score != 0 {
		countQuery = countQuery.Where("gt.score = ?", score)
	}
	if status != "ALL" {
		countQuery = countQuery.Where("gt.is_pass = ?", isPass)
	}
	if baseName != "" {
		countQuery = countQuery.Where("b.name = ?", baseName)
	}
	if sessionName != "" {
		countQuery = countQuery.Where("s.name = ?", sessionName)
	}
	if email != "" {
		countQuery = countQuery.Where("u.email = ?", email)
	}
	if startDate != "" && endDate != "" {
		countQuery = countQuery.Where("gt.created_date >= ? AND gt.created_date < ?", startDate, endDate)
	}

	err := countQuery.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		offset := (page - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}

	err = query.Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, int(totalCount), nil
}
