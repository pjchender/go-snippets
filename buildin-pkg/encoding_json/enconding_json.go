package jsonsandbox

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name          string `json:"name"`
	Capacity      uint   `json:"capacity"`
	Configuration string `json:",string"` // 轉換後 configuration 中也是 string
}

type Configuration struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

func Unmarshal() {
	// configuration 也是 JSON
	jsonString := `{"name":"battery sensor","capacity":40,"configuration":"{\\"name\\":\\"config\\",\\"state\\":\\"low battery\\"}"}`

	reading := SensorReading{}

	/**
	 * 將 JSON 轉成 struct 需要先把 string 轉成 byte slice，
	 * 然後再透過 Unmarshal 把空的 Struct 帶入
	 **/
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
	// {Name:battery sensor Capacity:40 Time:2019-01-21 19:07:28 +0000 UTC}
}

type Book struct {
	BookID        uint   `json:"-"`              // 轉換時總是忽略掉此欄位
	Title         string `json:"title"`          // 轉換後 JSON 的欄位名稱是 title
	Author        string `json:"author"`         // 轉換後 JSON 的欄位名稱是 author
	Name          string `json:"name,omitempty"` // 當 name 是空值時轉換後則無該欄位
	Age           uint   `json:",omitempty"`     // 當 Age 有值時，則JSON 的欄位名稱是 "Age"（大寫開頭），否則不顯示該欄位
	Price         uint   `json:"_,"`             // 轉換後 JSON 的欄位名稱是 _
	Configuration string `json:",string"`        // 轉換後 configuration 中也是 string
}

func Marshal() {
	book := Book{BookID: 2, Title: "Learning Go", Author: "Gopher", Name: "", Age: 30, Price: 31900}

	/* 將 struct 轉成 byte slice，再透過 string 變成 JSON 格式 */
	byteSlice, _ := json.MarshalIndent(book, "", "  ")
	fmt.Println(string(byteSlice))

	//{
	//	"title": "Learning Go",
	//	"author": "Gopher",
	//	"_": 31900
	//}
}
