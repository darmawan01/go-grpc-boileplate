package db

import (
	"fmt"
	"time"

	"go_grpc_boileplate/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open() (db *gorm.DB, err error) {
	db, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: fmt.Sprintf(
				"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
				configs.Config.DB.User, configs.Config.DB.Pass, configs.Config.DB.Name, configs.Config.DB.Host, configs.Config.DB.Port,
			),
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)

	if err != nil {
		return
	}

	if !configs.Config.IsProduction() {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}
	dbGorm, err := db.DB()
	if err != nil {
		return
	}
	dbGorm.SetMaxOpenConns(configs.Config.DB.MaxOpenConn)
	dbGorm.SetMaxIdleConns(configs.Config.DB.MaxIdleConn)
	dbGorm.SetConnMaxLifetime(time.Duration(configs.Config.DB.MaxLifeTime) * time.Minute)

	return
}
