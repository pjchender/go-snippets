package global

import "gorm.io/gorm"

var (
	// DBEngine 是透過 gorm.Open 與資料庫連線後取得
	DBEngine *gorm.DB
)
