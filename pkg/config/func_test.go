package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"pttep-vr-api/pkg/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	fmt.Println("configPath:", configPath)
	defer func() {
		//err := os.Remove(configPath)
		//if err != nil {
		//	fmt.Println(err)
		//}
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Println(err)
		}
	}()

	err := os.WriteFile(configPath, []byte(`
app:
  name: ${APP_NAME}
  version: ${APP_VERSION}
  state: ${APP_ENVIRONMENT}
  timezone: ${APP_TIMEZONE}
  config:
    host: ${APP_HOST}
    port: ${APP_PORT}
    path: ${APP_PATH_ROOT}
    allows:
      origins:
        - "*"
      response:
        error: ${APP_RESPONSE_ERROR}
database:
  host: ${DB_HOST}
  port: ${DB_PORT}
  name: ${DB_NAME}
  username: ${DB_USERNAME}
  password: ${DB_PASSWORD}
`), 0644)
	assert.NoError(t, err)

	os.Setenv("APP_ENVIRONMENT", "LOCAL")
	defer os.Unsetenv("APP_ENVIRONMENT")
	os.Setenv("APP_HOST", "localhost")
	defer os.Unsetenv("APP_HOST")
	os.Setenv("APP_NAME", "PTTEP-VR-API")
	defer os.Unsetenv("APP_NAME")
	os.Setenv("APP_PATH_ROOT", "")
	defer os.Unsetenv("APP_PATH_ROOT")
	os.Setenv("APP_PORT", "8080")
	defer os.Unsetenv("APP_PORT")
	os.Setenv("APP_RESPONSE_ERROR", "true")
	defer os.Unsetenv("APP_RESPONSE_ERROR")
	os.Setenv("APP_TIMEZONE", "Asia/Bangkok")
	defer os.Unsetenv("APP_TIMEZONE")
	os.Setenv("APP_VERSION", "0.0.1")
	defer os.Unsetenv("APP_VERSION")
	os.Setenv("DB_HOST", "db_host")
	defer os.Unsetenv("DB_HOST")
	os.Setenv("DB_NAME", "db_name")
	defer os.Unsetenv("DB_NAME")
	os.Setenv("DB_PASSWORD", "db_password")
	defer os.Unsetenv("DB_PASSWORD")
	os.Setenv("DB_PORT", "3306")
	defer os.Unsetenv("DB_PORT")
	os.Setenv("DB_USERNAME", "db_username")
	defer os.Unsetenv("DB_USERNAME")

	t.Run("Replace Env In Config", func(t *testing.T) {
		os.Setenv("PORT", "8080")
		defer os.Unsetenv("PORT")

		err = config.Init(configPath)
		assert.NoError(t, err)

		assert.Equal(t, 8080, config.Get().App.Config.Port)
	})

	t.Run("Replace Env In Config", func(t *testing.T) {
		os.Setenv("PORT", "PORT")
		defer os.Unsetenv("PORT")

		err = config.Init(configPath)
		assert.Error(t, err)
	})

	t.Run("Replace Env In Config", func(t *testing.T) {
		err := config.Init(tmpDir)
		assert.Error(t, err)
	})
}

func Test2(t *testing.T) {

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	defer func() {
		//err := os.Remove(configPath)
		//if err != nil {
		//	fmt.Println(err)
		//}
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Println(err)
		}
	}()

	err := os.WriteFile(configPath, []byte(`
app:
  name: APP_NAME
  version: ${APP_VERSION}
  state: ${APP_ENVIRONMENT}
  timezone: ${APP_TIMEZONE}
  config:
    host: ${APP_HOST}
    port: ${APP_PORT}
    path: ${APP_PATH_ROOT}
    allows:
      origins:
        - "*"
      response: ${APP_RESPONSE_ERROR}
`), 0644)
	assert.NoError(t, err)

	os.Setenv("APP_RESPONSE_ERROR", "ERROR")
	defer os.Unsetenv("APP_RESPONSE_ERROR")

	err = config.Init(configPath)
	assert.Error(t, err)

}
