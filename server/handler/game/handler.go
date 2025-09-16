package game

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/services/game"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/server/response"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

type handler struct {
	service *game.Service
}

func newHandler(service *game.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreatTransactionGame(ctx *fiber.Ctx) error {
	var transReq models.GameTransactionReq

	if err := ctx.BodyParser(&transReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("invalid req body"), nil, err))
	}

	err := h.service.CreateGameTransactionPlayer(context.Background(), transReq)
	if err != nil {
		log.Print("Error CreateGameTransactionPlayer", err)
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err create game player"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, err))
}

func (h *handler) Dashboard(ctx *fiber.Ctx) error {
	var transReq models.GameTransactionReq

	if err := ctx.BodyParser(&transReq); err != nil {
		return ctx.Status(http.StatusOK).JSON(DefaultResponse{
			Code:    100,
			Message: "Invalid request body",
		})
	}

	// resp, err := h.service.CreateGameTransactionPlayer(context.Background(), gameTrans)
	// if err != nil {
	// 	return ctx.Status(http.StatusOK).JSON(resp)
	// }

	return ctx.Status(http.StatusOK).JSON("")
}

func (h *handler) CreatTransactionGameTemp(ctx *fiber.Ctx) error {
	var transReq GameTransactionReqTemp

	if err := ctx.BodyParser(&transReq); err != nil {
		return ctx.Status(http.StatusOK).JSON(DefaultResponse{
			Code:    100,
			Message: "Invalid request body",
		})
	}

	gameTrans := models.GameTransactionTemp{
		SessionName:  transReq.SessionName,
		QuestionName: transReq.QuestionName,
		BaseName:     transReq.BaseName,
		Start:        transReq.Start,
		Stop:         transReq.Stop,
		Time:         transReq.Time,
		Score:        transReq.Score,
		PassPercent:  transReq.PassPercent,
		IsPass:       transReq.IsPass,
		Username:     transReq.Username,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
	}

	fmt.Println("gameTrans :", gameTrans)

	resp, err := h.service.CreateGameTransactionPlayerTemp(context.Background(), gameTrans)
	if err != nil {
		log.Print("Error CreateGameTransactionPlayerTemp", err)
		return ctx.Status(http.StatusOK).JSON(resp)
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (h *handler) GetSettingGameUser(ctx *fiber.Ctx) error {
	key := ctx.Params("key")

	resp, err := h.service.GetSettingGameByUser(context.Background(), key)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("get_setting_game_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) GetSettingGameId(ctx *fiber.Ctx) error {
	id := ctx.Params("key")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("parse int err"), nil, err))
	}

	resp, err := h.service.GetSettingGameById(context.Background(), uint(idInt))
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("get_setting_game_id_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) GetSettingGame(ctx *fiber.Ctx) error {

	resp, err := h.service.GetGameList(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("get_setting_game_list_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) SettingGame(ctx *fiber.Ctx) error {
	var questReq SettingGameReq

	if err := ctx.BodyParser(&questReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))

	}

	update := models.Question{
		ID:             questReq.QuestId,
		Score:          questReq.Score,
		MaxScore:       questReq.MaxScore,
		MinScore:       questReq.MinScore,
		ScorePerChoice: questReq.ScorePerChoice,
		Time:           questReq.Time,
	}

	resp, err := h.service.UpdateSettingGameById(context.Background(), update)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err update game by id"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), nil, err))
}

func (h *handler) ReportTransGame(ctx *fiber.Ctx) error {
	var reportReq models.ReportExportReq
	if err := ctx.BodyParser(&reportReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, count, err := h.service.ReportGame(context.Background(), reportReq)

	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("report_list_not_found"), nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), map[string]interface{}{
		"data":  resp,
		"count": count,
	}, nil))
}

func (h *handler) DashboardSumTimeUser(ctx *fiber.Ctx) error {
	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}
	resp, err := h.service.DashBoardSumTime(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_sum_time_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardSumUser(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetDistinctUsernameCountByDateRange(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_sum_user_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardAvgScore(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetAvgScore(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_sum_avg_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardGameFail(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetMaxFalseQuestion(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_game_fail_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardGraph(ctx *fiber.Ctx) error {

	var dashboardReq models.BaseReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetGraphScoreSevenDay(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_graph_fail_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardPassRate(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetPassRateByDate(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_pass_rate_fail_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardMostQuestion(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetMostQuestionByDate(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_most_quest_fail_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardTopQuestionsByTimeByDate(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetTopQuestionsByTimeByDate(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_quest_time_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardTopScoreUserByTimeByDate(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetTotalScoreForUsersbyDate(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_total_score_user_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) DashboardTopQuest(ctx *fiber.Ctx) error {

	var dashboardReq models.DashboardReq

	if err := ctx.BodyParser(&dashboardReq); err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("err parse json"), nil, err))
	}

	resp, err := h.service.GetQuestionPlayCount(context.Background(), dashboardReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("dashboard_top_quest_user_not_found"), resp, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), resp, nil))
}

func (h *handler) ExportTransGame(ctx *fiber.Ctx) error {
	var reportReq models.ReportExportReq
	if err := ctx.BodyParser(&reportReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.New(ctx.Context(), errorMessage.Get("err_parse_json"), nil, err))
	}

	resp, _, err := h.service.ReportGame(context.Background(), reportReq)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.New(ctx.Context(), errorMessage.Get("rport_list_not_found"), resp, err))
	}

	f := excelize.NewFile()

	sheetName := "Transactions"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.New(ctx.Context(), errorMessage.Get("err_new_sheet"), resp, err))
	}

	f.SetCellValue(sheetName, "A1", "name")
	f.SetCellValue(sheetName, "B1", "email")
	f.SetCellValue(sheetName, "C1", "location")
	f.SetCellValue(sheetName, "D1", "session")
	f.SetCellValue(sheetName, "E1", "quest_name")
	f.SetCellValue(sheetName, "F1", "score")
	f.SetCellValue(sheetName, "G1", "is_pass")
	f.SetCellValue(sheetName, "H1", "created_date")

	for i, t := range resp {
		row := i + 2

		var passStr string
		var name string

		if t.IsPass == "1" {
			passStr = "Pass"
		} else {
			passStr = "Failed"
		}

		if t.Name == "" {
			name = t.Username
		} else {
			name = t.Name
		}

		date := t.CreatedDate

		layout := time.RFC3339

		time, err := time.Parse(layout, date)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(response.New(ctx.Context(), errorMessage.Get("err parse date"), resp, err))
		}

		dateFormat := time.Format("2006-01-02 15:04:05")

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), name)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), t.Email)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), t.BaseName)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), t.SessionName)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), t.QuestName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), t.Score)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), passStr)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), dateFormat)

	}

	f.SetActiveSheet(index)

	buf, err := f.WriteToBuffer()
	if err != nil {
		log.Printf("Error writing Excel to buffer: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(ctx.Context(), errorMessage.Get("err_save_excel"), nil, err))
	}

	bufByte := buf.Bytes()
	file := base64.StdEncoding.EncodeToString(bufByte)

	return ctx.Status(fiber.StatusOK).JSON(response.New(ctx.Context(), errorMessage.Get("success"), file, nil))
}
