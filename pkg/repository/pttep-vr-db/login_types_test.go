package pttep_vr_db_test

import (
	"context"
	"pttep-vr-api/pkg/models"
	pttep_vr_db "pttep-vr-api/pkg/repository/pttep-vr-db"
	"pttep-vr-api/pkg/utils/gormDB"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestLoginTypes(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbMock, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gormDBMock := gormDB.NewMockInterface(ctrl)
	gormDBMock.EXPECT().GetDB().Return(dbMock).AnyTimes()
	o := pttep_vr_db.New(gormDBMock)
	assert.NotNil(t, o.DB())

	ctx := context.Background()

	t.Run("FindLoginTypes", func(t *testing.T) {
		_, err := o.FindLoginTypes(ctx)
		assert.Error(t, err)
	})

	t.Run("FindOneLoginTypes", func(t *testing.T) {
		_, err := o.FindOneLoginTypes(ctx, models.LoginType{})
		assert.Error(t, err)
	})

	t.Run("InsertOneLoginTypes", func(t *testing.T) {
		_, err := o.InsertOneLoginTypes(ctx, models.LoginType{})
		assert.Error(t, err)
	})
}
