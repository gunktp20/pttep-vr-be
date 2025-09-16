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

func TestUserLogins(t *testing.T) {
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

	t.Run("FindUserLogins", func(t *testing.T) {
		_, err := o.FindUserLogins(ctx)
		assert.Error(t, err)
	})

	t.Run("FindOneUserLogins", func(t *testing.T) {
		_, err := o.FindOneUserLogins(ctx, models.UserLogin{})
		assert.Error(t, err)
	})

	t.Run("FindOneUserLoginsByUsername", func(t *testing.T) {
		_, err := o.FindOneUserLoginsByUsername(ctx, models.UserLogin{Username: "test", LoginTypeID: 1})
		assert.Error(t, err)
	})

	t.Run("FindOneUserLoginsByUsernameAndPassword", func(t *testing.T) {
		_, err := o.FindOneUserLoginsByUsernameAndPassword(ctx, models.UserLogin{
			Username:    "test",
			Password:    "pass",
			LoginTypeID: 1,
		})
		assert.Error(t, err)
	})

	t.Run("InsertOneUserLogins", func(t *testing.T) {
		_, err := o.InsertOneUserLogins(ctx, models.UserLogin{})
		assert.Error(t, err)
	})

	t.Run("UpdateOneUserLogins", func(t *testing.T) {
		_, err := o.UpdateOneUserLogins(ctx, models.UserLogin{UserID: 1, LoginTypeID: 1})
		assert.Error(t, err)
	})

	t.Run("UpdateOneUserLogins_WithPassword", func(t *testing.T) {
		_, err := o.UpdateOneUserLogins(ctx, models.UserLogin{
			UserID:      1,
			LoginTypeID: 1,
			Password:    "newpass",
		})
		assert.Error(t, err)
	})
}
