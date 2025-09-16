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

func TestUserRoles(t *testing.T) {
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

	t.Run("FindUserRoles", func(t *testing.T) {
		_, err := o.FindUserRoles(ctx)
		assert.Error(t, err)
	})

	t.Run("FindUserRolesByUserID", func(t *testing.T) {
		_, err := o.FindUserRolesByUserID(ctx, models.UserRole{UserID: 1})
		assert.Error(t, err)
	})

	t.Run("FindOneUserRoles", func(t *testing.T) {
		_, err := o.FindOneUserRoles(ctx, models.UserRole{ID: 1})
		assert.Error(t, err)
	})

	t.Run("InsertOneUserRoles", func(t *testing.T) {
		_, err := o.InsertOneUserRoles(ctx, models.UserRole{})
		assert.Error(t, err)
	})

	t.Run("UpdateOneUserRolesRole", func(t *testing.T) {
		_, err := o.UpdateOneUserRolesRole(ctx, models.UserRole{ID: 1, RoleID: 1})
		assert.Error(t, err)
	})

	t.Run("UpdateOneUserRoles", func(t *testing.T) {
		_, err := o.UpdateOneUserRoles(ctx, models.UserRole{})
		assert.Error(t, err)
	})

	t.Run("UpdateUserRolesByEmail", func(t *testing.T) {
		err := o.UpdateUserRolesByEmail(ctx, models.UserRole{UserID: 1, Email: "test@test.com"})
		assert.Error(t, err)
	})

	t.Run("UpdateOneUserRolesIsActive", func(t *testing.T) {
		err := o.UpdateOneUserRolesIsActive(ctx, models.UserRole{ID: 1, IsActive: true})
		assert.Error(t, err)
	})

	t.Run("UpdateOneUserRolesIsRemove", func(t *testing.T) {
		err := o.UpdateOneUserRolesIsRemove(ctx, models.UserRole{ID: 1, IsRemove: true})
		assert.Error(t, err)
	})

	t.Run("FindUserRolesLeftJoinUserAndRolePermissionAndRole", func(t *testing.T) {
		_, _, err := o.FindUserRolesLeftJoinUserAndRolePermissionAndRole(ctx, p)
		assert.Error(t, err)
	})

	t.Run("FindUserRolesLeftJoinUserAndRolePermissionAndRole_NilPagination", func(t *testing.T) {
		_, _, err := o.FindUserRolesLeftJoinUserAndRolePermissionAndRole(ctx, nil)
		assert.Error(t, err)
	})

	t.Run("FindOneUserRolesLeftJoinUserAndRolePermissionAndRole", func(t *testing.T) {
		_, err := o.FindOneUserRolesLeftJoinUserAndRolePermissionAndRole(ctx, models.UserRole{ID: 1})
		assert.Error(t, err)
	})
}
