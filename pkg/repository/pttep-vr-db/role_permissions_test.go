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

func TestRolePermissions(t *testing.T) {
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

	t.Run("FindRolePermissionsJoinPermissionsInRoleID", func(t *testing.T) {
		rolePermissions := []models.RolePermission{
			{RoleID: 1},
			{RoleID: 2},
		}
		_, _, err := o.FindRolePermissionsJoinPermissionsInRoleID(ctx, rolePermissions)
		assert.Error(t, err)
	})

	t.Run("FindRolePermissionsJoinPermissionsInRoleID_EmptySlice", func(t *testing.T) {
		rolePermissions := []models.RolePermission{}
		_, _, err := o.FindRolePermissionsJoinPermissionsInRoleID(ctx, rolePermissions)
		assert.Error(t, err)
	})

	t.Run("InsertOneRolePermissions", func(t *testing.T) {
		rolePermission := models.RolePermission{
			RoleID:       1,
			PermissionID: 1,
		}
		_, err := o.InsertOneRolePermissions(ctx, rolePermission)
		assert.Error(t, err)
	})

	t.Run("InsertManyRolePermissions", func(t *testing.T) {
		rolePermissions := []models.RolePermission{
			{RoleID: 1, PermissionID: 1},
			{RoleID: 1, PermissionID: 2},
			{RoleID: 2, PermissionID: 1},
		}
		_, err := o.InsertManyRolePermissions(ctx, rolePermissions)
		assert.Error(t, err)
	})

	t.Run("InsertManyRolePermissions_EmptySlice", func(t *testing.T) {
		rolePermissions := []models.RolePermission{}
		_, err := o.InsertManyRolePermissions(ctx, rolePermissions)
		assert.Error(t, err)
	})

	t.Run("DeleteOneRolePermissions", func(t *testing.T) {
		rolePermission := models.RolePermission{
			ID: 1,
		}
		err := o.DeleteOneRolePermissions(ctx, rolePermission)
		assert.Error(t, err)
	})

	t.Run("DeleteOneRolePermissionsByRoleIDAndPermissionID", func(t *testing.T) {
		rolePermission := models.RolePermission{
			RoleID:       1,
			PermissionID: 1,
		}
		err := o.DeleteOneRolePermissionsByRoleIDAndPermissionID(ctx, rolePermission)
		assert.Error(t, err)
	})

	t.Run("DeleteManyRolePermissionsByRoleID", func(t *testing.T) {
		rolePermission := models.RolePermission{
			RoleID: 1,
		}
		err := o.DeleteManyRolePermissionsByRoleID(ctx, rolePermission)
		assert.Error(t, err)
	})
}
