# enum with string value

#### package mode

```go
// modified from: https://github.com/gotify/server/blob/master/mode/mode.go
package mode

type Mode string

var (
	Dev  Mode = "dev"
	Prod Mode = "prod"
	Test Mode = "test"
)

var mode = Dev   // mode 預設值是 Dev

func Set(newMode Mode) {
	mode = newMode
}

func Get() Mode {
	return mode
}
```

#### main

```go
func main() {
  // 取得預設的 mode
	currentMode := mode.Get()
	fmt.Println("currentMode", currentMode) // dev

	// 使用 Set 更新 mode
	mode.Set(mode.Prod)
	updatedMode := mode.Get()
	fmt.Println("updatedMode", updatedMode) // prod

	// 轉成字串
	modeString := string(updatedMode)
	fmt.Println("modeString", modeString) // prod
}
```
