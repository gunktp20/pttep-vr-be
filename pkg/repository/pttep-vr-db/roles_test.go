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

func TestRoles(t *testing.T) {
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

	t.Run("FindRoles", func(t *testing.T) {
		_, _, err := o.FindRoles(ctx, p)
		assert.Error(t, err)
	})

	t.Run("FindRoles_NilPagination", func(t *testing.T) {
		_, _, err := o.FindRoles(ctx, nil)
		assert.Error(t, err)
	})

	t.Run("FindOneRoles", func(t *testing.T) {
		_, err := o.FindOneRoles(ctx, 1)
		assert.Error(t, err)
	})

	t.Run("FindOneRolesByIsDefault", func(t *testing.T) {
		_, err := o.FindOneRolesByIsDefault(ctx, true)
		assert.Error(t, err)
	})

	t.Run("InsertOneRoles", func(t *testing.T) {
		_, err := o.InsertOneRoles(ctx, models.Role{})
		assert.Error(t, err)
	})

	t.Run("UpdateOneRoles", func(t *testing.T) {
		_, err := o.UpdateOneRoles(ctx, models.Role{})
		assert.Error(t, err)
	})

	t.Run("UpdateOneRolesName", func(t *testing.T) {
		_, err := o.UpdateOneRolesName(ctx, models.Role{ID: 1, Name: "test"})
		assert.Error(t, err)
	})

	t.Run("UpdateOneRolesIsActive", func(t *testing.T) {
		_, err := o.UpdateOneRolesIsActive(ctx, models.Role{ID: 1, IsActive: true})
		assert.Error(t, err)
	})

	t.Run("UpdateOneRolesIsRemove", func(t *testing.T) {
		_, err := o.UpdateOneRolesIsRemove(ctx, models.Role{ID: 1, IsRemove: true})
		assert.Error(t, err)
	})
}
