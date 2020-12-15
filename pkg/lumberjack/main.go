package main

import (
	"log"

	"github.com/natefinch/lumberjack"
	"github.com/pjchender/go-snippets/pkg/lumberjack/global"
	"github.com/pjchender/go-snippets/pkg/lumberjack/logger"
)

func init() {
	err := setupLogger()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	global.Logger.Infof("%s: go-snippets/%s", "pkg", "blog-service")
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
