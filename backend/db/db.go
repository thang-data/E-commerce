package db

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thang-data/backend/config"
	"github.com/thang-data/backend/log"
	"github.com/thang-data/backend/log/gormlogrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var sqlDB *sql.DB
var db *gorm.DB
var err error

func Init() {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	// gormlogrus

	level := log.ParseLevel(cfg)
	sqlDB, err = sql.Open("pgx", dsn)
	if err != nil {
		logrus.Panic(err.Error())
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{Logger: gormlogrus.New().LogMode(logger.LogLevel(level))})
	if err != nil {
		logrus.Panic(err.Error())
	}
	dsnRead := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		cfg.DBHostReadonly, cfg.DBPortReadonly, cfg.DBUserReadonly, cfg.DBPasswordReadonly, cfg.DBNameReadonly)
	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.Open(dsnRead)},
		Policy:   dbresolver.RandomPolicy{},
	}))
}
func Connect() *gorm.DB {
	if db == nil {
		Init()
	}
	return db
}
