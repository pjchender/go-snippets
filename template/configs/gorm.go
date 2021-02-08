package configs

import (
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var Gorm = &gorm.Config{
	Logger:                                   gormLogger.Default.LogMode(gormLogger.Error),
	DisableForeignKeyConstraintWhenMigrating: true,
}
