package global

import "github.com/pjchender/go-snippets/pkg/lumberjack/logger"

var (
	// Logger 用來把 pkg/logger 的物件保存在 global package 中
	Logger *logger.Logger
)
