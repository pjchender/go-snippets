package global

import (
	"github.com/pjchender/go-snippets/template/configs"
	"gorm.io/gorm"
)

var (
	AppSetting        *configs.App
	HTTPServerSetting *configs.HTTPServer
	GRPCSetting       *configs.GRPC
	DatabaseSetting   *configs.Database
	AuthSetting       *configs.Auth
	GormSetting       *gorm.Config
)
