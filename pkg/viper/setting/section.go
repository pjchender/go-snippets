package setting

import "time"

// ServerSettingS 用來將 config.yaml 中的 server 載入
type ServerSettingS struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettingS 用來將 config.yaml 中的 app 載入
type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

// DatabaseSettingS 用來將 config.yaml 中的 database 載入
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

// ReadSection 會將 config.yaml 中的檔案轉成 go 可以讀取的 struct
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
