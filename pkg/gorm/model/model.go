package model

import (
	"fmt"
	"time"

	"github.com/pjchender/go-snippets/pkg/gorm/global"
	"github.com/pjchender/go-snippets/pkg/gorm/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDBEngine 用來與資料庫建立連線，並回傳 gorm.DB
// https://gorm.io/docs/connecting_to_the_database.html
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		databaseSetting.Host,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Truncate(time.Second)
		},
	})

	if err != nil {
		return nil, err
	}

	if global.RunMode == "debug" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
