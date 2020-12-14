package setting

// DatabaseSettingS 用來定義 DatabaseSetting 的 struct
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// GetDatabaseSetting 是用來取得 Database 的設定值
// 如果有使用 viper 的話，則可以直接透過 viper 讀取設定檔的內容
func GetDatabaseSetting() *DatabaseSettingS {

	databaseSetting := DatabaseSettingS{
		DBType:       "postgres",
		UserName:     "postgres",
		Password:     "",
		Host:         "127.0.0.1",
		DBName:       "blog_service",
		TablePrefix:  "blog_",
		Charset:      "utf8",
		ParseTime:    true,
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	}

	return &databaseSetting
}
