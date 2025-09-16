package pttep_vr_db_test

import (
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

func UserTest(t *testing.T) {
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

	// User
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
		_, err := o.FindOneUsersByEmail(models.User{})
		assert.Error(t, err)
	})
	t.Run("InsertOneUsers", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.InsertOneUsers(models.User{})
		assert.Error(t, err)
	})
	t.Run("UpdateOneUsers", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.UpdateOneUsers(models.User{})
		assert.Error(t, err)
	})
	t.Run("FindOneUserByCode", func(t *testing.T) {
		gormDBMock.EXPECT().GetDB().Return(dbMock)
		_, err := o.FindOneUserByCode("1234")
		assert.Error(t, err)
	})

}
