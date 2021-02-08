package setting

import (
	"github.com/pjchender/go-snippets/template/configs"
	"gorm.io/gorm"
)

func (s *Setting) ReadGormSetting() *gorm.Config {
	return configs.Gorm
}
