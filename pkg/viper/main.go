package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pjchender/go-snippets/pkg/viper/global"
	"github.com/pjchender/go-snippets/pkg/viper/setting"
)

// STEP 2：在 init 中載入 setting
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatal(err)
	}
}

// STEP 3：透過 global 取用載入好的 setting
func main() {
	fmt.Printf("ServerSetting: %+v \n", global.ServerSetting)
}

// STEP 1：建立載入 setting 的 function
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
