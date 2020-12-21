package database

import (
	"log"

	"github.com/pjchender/go-snippets/template/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormDatabase 中包含 *gorm.DB
type GormDatabase struct {
	DB *gorm.DB
}

// NewGormDatabase 會以 Gorm 和 datbase 建立連線
func NewGormDatabase(dsn string, gormConfig *gorm.Config) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}

// AutoMigrate 會執行 gorm 提供的 auto migration
func (d *GormDatabase) AutoMigrate() {
	// enable format UUID as PK
	d.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := d.DB.AutoMigrate(
		&model.Product{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

// DropAllTables 會移除 database 中所有的 tables
func (d *GormDatabase) DropAllTables() {
	if err := d.DB.Migrator().DropTable(
		"products",
	); err != nil {
		log.Fatal(err.Error())
	}
}
