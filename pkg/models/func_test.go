package models_test

import (
	"pttep-vr-api/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("Models", func(t *testing.T) {
		activityLog := models.ActivityLog{}
		assert.Equal(t, "activity_logs", activityLog.CollectionName())
		assert.Equal(t, "activity_logs", activityLog.TableName())

		base := models.Base{}
		assert.Equal(t, "bases", base.CollectionName())
		assert.Equal(t, "bases", base.TableName())

		choice := models.Choice{}
		assert.Equal(t, "choices", choice.CollectionName())
		assert.Equal(t, "choices", choice.TableName())

		gameTransaction := models.GameTransaction{}
		assert.Equal(t, "game_transactions", gameTransaction.CollectionName())
		assert.Equal(t, "game_transactions", gameTransaction.TableName())

		gameTransactionTemp := models.GameTransactionTemp{}
		assert.Equal(t, "game_transactions_temp", gameTransactionTemp.CollectionName())
		assert.Equal(t, "game_transactions_temp", gameTransactionTemp.TableName())

		question := models.Question{}
		assert.Equal(t, "questions", question.CollectionName())
		assert.Equal(t, "questions", question.TableName())

		questionType := models.QuestionType{}
		assert.Equal(t, "question_types", questionType.CollectionName())
		assert.Equal(t, "question_types", questionType.TableName())

		session := models.Session{}
		assert.Equal(t, "sessions", session.CollectionName())
		assert.Equal(t, "sessions", session.TableName())

		user := models.User{}
		assert.Equal(t, "users", user.CollectionName())
		assert.Equal(t, "users", user.TableName())

		userTemp := models.UserTemp{}
		assert.Equal(t, "users_temp", userTemp.CollectionName())
		assert.Equal(t, "users_temp", userTemp.TableName())

	})
}
