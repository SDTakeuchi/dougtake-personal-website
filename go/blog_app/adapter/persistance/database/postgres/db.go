package postgres

import (
	"blog_app/adapter/config"
	"fmt"
	"sync"

	"github.com/pingcap/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	isDebug := config.Get().Debug
	conf := config.Get().DB

	var logLevel logger.LogLevel
	if isDebug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}

	gl := logger.Default.LogMode(logLevel)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gl,
	})
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	sql.SetMaxOpenConns(conf.MaxOpen)
	sql.SetMaxIdleConns(conf.MaxIdle)
	sql.SetConnMaxLifetime(conf.MaxLifeTime)
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
