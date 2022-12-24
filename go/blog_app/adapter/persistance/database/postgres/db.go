package postgres

import (
	"blog_app/adapter/config"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

var (
	connOnce sync.Once
	conn     *DB
)

type DB struct {
	Conn *gorm.DB
}

func ConnectDB() *DB {
	connOnce.Do(func() {
		conn = &DB{
			Conn: newConnection(),
		}
	})
	return conn
}

func newConnection() *gorm.DB {
	dsn := buildDNS()
	conf := config.Get().DB
	isDebug := conf.Debug
	
}

func buildDNS() string {
	conf := config.Get().DB
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s connect_timeout=%s",
		conf.HostName,
		conf.Port,
		conf.UserName,
		conf.Database,
		conf.SSLMode,
		conf.Password,
		conf.ConnTimeout,
	)
}
