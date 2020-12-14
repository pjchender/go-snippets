package global

import "github.com/pjchender/go-snippets/pkg/viper/setting"

var (
	// ServerSetting 將透過 viper 讀取得 server 設定灌到全域
	ServerSetting *setting.ServerSettingS
	// AppSetting 將透過 viper 讀取得 app 設定灌到全域
	AppSetting *setting.AppSettingS
	// DatabaseSetting 將透過 viper 讀取得 database 設定灌到全域
	DatabaseSetting *setting.DatabaseSettingS
)
