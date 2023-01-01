package postgres

import (
	"blog_app/adapter/config"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pingcap/log"
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
	// isDebug := conf.Debug

	// var logLevel logger.LogLevel
	// if isDebug {
	// 	logLevel = logger.Info
	// } else {
	// 	logLevel = logger.Warn
	// }

	// gl := logger.Default.LogMode(logLevel)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	db.DB().SetMaxOpenConns(conf.MaxOpen)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	db.DB().SetConnMaxLifetime(conf.MaxLifeTime)
	return db
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
