package gormDB

import (
	"fmt"

	"gorm.io/gorm"
)

type Interface interface {
	GetDB() *gorm.DB
	Connect(Dialector) error
}

func New(host string, port int, username, password, dbname string) *Client {
	return &Client{
		config: config{
			username: username,
			password: password,
			host:     host,
			port:     port,
			name:     dbname,
		},
	}
}

type Client struct {
	config config
	client *gorm.DB
}

type config struct {
	username string
	password string
	host     string
	port     int
	name     string
}

func (c *Client) GetDB() *gorm.DB {
	return c.client
}

func (c *Client) Connect(dlt Dialector) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.config.username, c.config.password, c.config.host, c.config.port, c.config.name)
	var err error
	c.client, err = gorm.Open(dlt(dsn), &gorm.Config{})
	return err
}

// Factory สำหรับสร้าง gorm.Dialector (ใช้ mysql.Open, sqlite.Open, sqlmock ฯลฯ)
type Dialector func(dsn string) gorm.Dialector
