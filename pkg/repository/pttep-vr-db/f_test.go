package pttep_vr_db_test

import (
	"pttep-vr-api/pkg/models"
	pttep_vr_db "pttep-vr-api/pkg/repository/pttep-vr-db"
	"pttep-vr-api/pkg/utils/gormDB"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test(t *testing.T) {
	//db, mock, err := sqlmock.New()
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbMock, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)
	//gormDBMock, err := gormDB.NewMock(dbMock)
	//assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gormDBMock := gormDB.NewMockInterface(ctrl)
	gormDBMock.EXPECT().GetDB().Return(dbMock)
	o := pttep_vr_db.New(gormDBMock)
	assert.NotNil(t, o.DB())

	t.Run("InsertOneUsersTemp", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneUsersTemp(models.UserTemp{})
		assert.Error(t, err)
	})
	t.Run("FindOneUsersTempByUsername", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindOneUsersTempByUsername("test")
		assert.Error(t, err)
	})

	t.Run("FindUsers", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindUsers()
		assert.Error(t, err)
	})
	t.Run("FindOneUsers", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindOneUsers(1)
		assert.Error(t, err)
	})
	t.Run("FindOneUsersByEmail", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindOneUsersByEmail(models.User{Email: "test"})
		assert.Error(t, err)
	})
	t.Run("InsertOneUsers", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneUsers(models.User{Email: "test"})
		assert.Error(t, err)
	})
	t.Run("UpdateOneUsers", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.UpdateOneUsers(models.User{Email: "test"})
		assert.Error(t, err)
	})
	t.Run("InsertOneGameTransaction", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneGameTransaction(models.GameTransaction{})
		assert.Error(t, err)
	})
	t.Run("FindSettingGameList", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindSettingGameList()
		assert.Error(t, err)
	})
	t.Run("FindSettingGameById", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindSettingGameById(1)
		assert.Error(t, err)
	})

	t.Run("UpdateSettingGameById", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.UpdateSettingGameById(models.Question{})
		assert.Error(t, err)
	})

	t.Run("InsertOneSettingGame", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneSettingGame(models.Session{})
		assert.Error(t, err)
	})

	t.Run("InsertOneSettingGameQuest", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneSettingGameQuest(models.Question{})
		assert.Error(t, err)
	})
	t.Run("InsertOneGameTransactionTemp", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneGameTransactionTemp(models.GameTransactionTemp{})
		assert.Error(t, err)
	})
	t.Run("FindSettingGameUser", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindSettingGameUser("test")
		assert.Error(t, err)
	})
	t.Run("FindRportTransGmaeUser", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindRportTransGmaeUser()
		assert.Error(t, err)
	})
	t.Run("FindDashboardTransGmaeUser", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindDashboardTransGmaeUser()
		assert.Error(t, err)
	})
	t.Run("SumTimeTransGmaeUser", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.SumTimeTransGmaeUser(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetDistinctUsernameCountByDateRange", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetDistinctUsernameCountByDateRange(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetAvgOfMaxAverageScore", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetAvgOfMaxAverageScore(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetMaxFalseIsPassByQuestion", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetMaxFalseIsPassByQuestion(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetGraphScoreSevenDay", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetGraphScoreSevenDay("base")
		assert.Error(t, err)
	})
	t.Run("GetGraphScoreSevenDay", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetPassRateByDate(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetMostQuestionByDate", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetMostQuestionByDate(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetTopQuestionsByTime", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetTopQuestionsByTime(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetTotalScoreForUsersWithDateRange", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetTotalScoreForUsersWithDateRange(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("GetTopQuestionsByPlayCount", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.GetTopQuestionsByPlayCount(time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		assert.Error(t, err)
	})
	t.Run("FindGameTransactions", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, _, err := o.FindGameTransactions("test", 1, true, "test", "test", "test", "test", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), 10, 1)
		assert.Error(t, err)
	})

}
