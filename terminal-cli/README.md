## 在 Terminal 產生問題，並取得使用者在 Terminal 輸入的內容

##### keywords: `bufio`, `fmt`, `log`, `os`, `strings`

> [add_person.go](https://github.com/protocolbuffers/protobuf/blob/master/examples/add_person.go) @ GitHub: protobuf

```go
func main() {
	// 建立 Reader
	r := bufio.NewReader(os.Stdin)

	// 詢問姓名
	fmt.Println("What's your name?")
	name, err := r.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	name = strings.TrimSpace(name)

	// 詢問年齡
	fmt.Println("What's your age?")
	age, err := r.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	age = strings.TrimSpace(age)

	// 輸出結果：Hello Aaron, you are 32 years old
	fmt.Printf("Hello %s, you are %s years old", name, age)
}
```
