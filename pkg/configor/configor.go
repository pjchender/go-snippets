package configor

import (
	"os"

	"github.com/jinzhu/configor"
)

// Configuration 的 struct field tag 中可以使用 default 欄位來定義預設值
// slice 和 struct 同樣可以定義預設值（https://github.com/jinzhu/configor/issues/29）
type Configuration struct {
	IsBool bool   `default:"true"`
	Mode   string `default:"production"`
	Cors   struct {
		AllowOrigins []string `default:"['GET','POST']"`
		AllowMethods []string
		AllowHeaders []string
	}
	StringArray []string `default:"['foo','bar']" env:"STRING_ARRAY"`
	StructArray []struct {
		Name string `default:"poop"`
	} `default:"[{'name':'foo'},{}]" env:"STRUCT_ARRAY"`
}

// Get can get the configuration
func Get() *Configuration {
	config := &Configuration{}

	if err := configor.Load(config); err != nil {
		panic(err)
	}

	mode := os.Getenv("MODE")
	if mode != "" {
		config.Mode = mode
	}

	return config
}
