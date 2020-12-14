package setting

import (
	"github.com/spf13/viper"
)

// Setting 是用來保存 viper 的 struct
type Setting struct {
	vp *viper.Viper
}

// NewSetting 會說明設定檔的名稱、類型和路徑，最終則會回傳帶有 Viper 實例的 Setting 物件
func NewSetting() (*Setting, error) {
	// 初始化 viper 實例
	vp := viper.New()

	// 設定 config 檔的類型、檔名和路徑
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("configs/")

	// 讀取 config 檔
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// 取出 config 的 key
	// fmt.Println("server.httpport", vp.Get("server.httpport")) // 8000
	// fmt.Println("server", vp.Get("server"))                   // map[httpport:8000 readtimeout:60 runmode:debug writetimeout:60]

	return &Setting{vp}, nil
}
