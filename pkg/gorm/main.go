package main

import (
	"log"

	"github.com/pjchender/go-snippets/pkg/gorm/global"
	"github.com/pjchender/go-snippets/pkg/gorm/model"
	"github.com/pjchender/go-snippets/pkg/gorm/setting"
)

func init() {
	err := setupDBEngine()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}

func setupDBEngine() error {
	var err error

	databaseSetting := setting.GetDatabaseSetting()
	global.DBEngine, err = model.NewDBEngine(databaseSetting)
	if err != nil {
		return err
	}

	return nil
}
