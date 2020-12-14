package main

import (
	"fmt"

	"github.com/pjchender/go-snippets/helpers/errcode"
)

func main() {
	fmt.Printf("%+v \n", errcode.TooManyRequest) // 錯誤 10000007, 錯誤訊息：請求過多
}
