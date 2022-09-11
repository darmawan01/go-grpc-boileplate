package db

import (
	"fmt"
	"time"

	"go_grpc_boileplate/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConn struct {
	Info       configs.ConnInfo
	SilentMode bool
}

func (conn *DBConn) Open() (db *gorm.DB, err error) {
	db, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: fmt.Sprintf(
				"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
				conn.Info.User,
				conn.Info.Pass,
				conn.Info.Name,
				conn.Info.Host,
				conn.Info.Port,
			),
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)

	if err != nil {
		return
	}

	if !conn.SilentMode {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	dbGorm, err := db.DB()
	if err != nil {
		return
	}
	dbGorm.SetMaxOpenConns(conn.Info.MaxOpenConn)
	dbGorm.SetMaxIdleConns(conn.Info.MaxIdleConn)
	dbGorm.SetConnMaxLifetime(time.Duration(conn.Info.MaxLifeTime) * time.Minute)

	return
}
