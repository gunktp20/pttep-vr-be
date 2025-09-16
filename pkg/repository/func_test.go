package repository_test

import (
	"pttep-vr-api/pkg/repository"
	"pttep-vr-api/pkg/utils/gormDB"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	ctrl := gomock.NewController(t)
	gormDBMock := gormDB.NewMockInterface(ctrl)

	//// สร้าง sqlmock
	//sqlDB, _, err := sqlmock.New()
	//if err != nil {
	//	t.Fatalf("failed to open sqlmock: %s", err)
	//}
	//defer sqlDB.Close()
	//
	//// mock DialectorFactory
	//mockDialector := func(dsn string) gorm.Dialector {
	//	return mysql.New(mysql.Config{
	//		Conn:                      sqlDB,
	//		SkipInitializeWithVersion: true,
	//	})
	//}

	t.Run("Operator New", func(t *testing.T) {
		//gormDBMock.EXPECT().Connect(mockDialector).Return(nil)
		_repository := repository.New(gormDBMock)
		assert.NotNil(t, _repository)
		assert.NotNil(t, _repository.PTTEPVR())
	})
}
