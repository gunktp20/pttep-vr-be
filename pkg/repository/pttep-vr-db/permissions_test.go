package pttep_vr_db_test

import (
	"context"
	"pttep-vr-api/pkg/models"
	pttep_vr_db "pttep-vr-api/pkg/repository/pttep-vr-db"
	"pttep-vr-api/pkg/utils/gormDB"
	"pttep-vr-api/pkg/utils/pagination"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestPermissions(t *testing.T) {
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
	p := &pagination.Pagination{}

	t.Run("FindPermissions", func(t *testing.T) {
		_, _, err := o.FindPermissions(ctx, p)
		assert.Error(t, err)
	})

	t.Run("FindPermissions_NilPagination", func(t *testing.T) {
		_, _, err := o.FindPermissions(ctx, nil)
		assert.Error(t, err)
	})

	t.Run("FindPermissionsIn", func(t *testing.T) {
		_, _, err := o.FindPermissionsIn(ctx, []models.Permission{{ID: 1}})
		assert.Error(t, err)
	})

	t.Run("FindOnePermissions", func(t *testing.T) {
		_, err := o.FindOnePermissions(ctx, 1)
		assert.Error(t, err)
	})

	t.Run("InsertOnePermissions", func(t *testing.T) {
		_, err := o.InsertOnePermissions(ctx, models.Permission{})
		assert.Error(t, err)
	})

	t.Run("UpdateOnePermissions", func(t *testing.T) {
		_, err := o.UpdateOnePermissions(ctx, models.Permission{})
		assert.Error(t, err)
	})

	t.Run("UpdateOnePermissionsName", func(t *testing.T) {
		_, err := o.UpdateOnePermissionsName(ctx, models.Permission{ID: 1, Name: "test"})
		assert.Error(t, err)
	})

	t.Run("UpdateOnePermissionsIsActive", func(t *testing.T) {
		_, err := o.UpdateOnePermissionsIsActive(ctx, models.Permission{ID: 1, IsActive: true})
		assert.Error(t, err)
	})

	t.Run("UpdateOnePermissionsIsRemove", func(t *testing.T) {
		_, err := o.UpdateOnePermissionsIsRemove(ctx, models.Permission{ID: 1, IsRemove: true})
		assert.Error(t, err)
	})
}
