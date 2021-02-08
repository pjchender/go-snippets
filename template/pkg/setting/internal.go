package setting

import (
	"os"

	"github.com/pjchender/go-snippets/template/configs"
	"github.com/pjchender/go-snippets/template/pkg/ginmode"
)

// ReadAppSetting 會設定 APP
func (s *Setting) ReadAppSetting() *configs.App {
	modeENV := os.Getenv("MODE")
	if modeENV != "" {
		ginmode.Set(&s.defaultConfig.App, modeENV)
	}

	return &s.defaultConfig.App
}

func (s *Setting) ReadDBSetting() *configs.Database {
	DSN := os.Getenv("DATABASE_URL")
	if DSN != "" {
		s.defaultConfig.Database.DSN = DSN
	}

	return &s.defaultConfig.Database
}

func (s *Setting) ReadAuthSetting() *configs.Auth {
	passwordSalt := os.Getenv("PASSWORD_SALT")
	if passwordSalt != "" {
		s.defaultConfig.Auth.Password.Salt = passwordSalt
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret != "" {
		s.defaultConfig.Auth.JWT.Secret = JWTSecret
	}

	return &s.defaultConfig.Auth
}

func (s *Setting) ReadServerSetting() *configs.HTTPServer {
	serverPort := os.Getenv("HTTP_SERVER_PORT")
	if serverPort != "" {
		s.defaultConfig.HTTPServer.Port = serverPort
	}

	return &s.defaultConfig.HTTPServer
}

func (s *Setting) ReadGRPCSetting() *configs.GRPC {
	serverPort := os.Getenv("GRPC_SERVER_PORT")
	if serverPort != "" {
		s.defaultConfig.GRPC.Server.Port = serverPort
	}

	clientPort := os.Getenv("GRPC_CLIENT_PORT")
	if clientPort != "" {
		s.defaultConfig.GRPC.Client.Port = clientPort
	}

	return &s.defaultConfig.GRPC
}
