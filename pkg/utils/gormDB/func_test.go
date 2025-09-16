package gormDB_test

import (
	"pttep-vr-api/pkg/utils/gormDB"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test(t *testing.T) {

	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	v := func(dns string) gorm.Dialector {
		return mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		})
	}

	client := gormDB.New("host", 3306, "username", "password", "dbname")
	err = client.Connect(v)
	assert.NoError(t, err)

	assert.NotNil(t, client.GetDB())

}
