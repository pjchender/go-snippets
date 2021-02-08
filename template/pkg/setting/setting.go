package setting

import (
	"github.com/jinzhu/configor"
	"github.com/pjchender/go-snippets/template/configs"
	log "github.com/sirupsen/logrus"
)

type Setting struct {
	defaultConfig *configs.Default
}

func NewSetting() (*Setting, error) {
	defaultConfig := configs.Default{}
	err := configor.Load(&defaultConfig)
	if err != nil {
		log.Error("[setting] configor.Load failed: ", err)
		return nil, err
	}

	return &Setting{defaultConfig: &defaultConfig}, nil
}
