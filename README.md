# Golang Snippets

## godotenv

```go
myEnv := os.Getenv("MY_ENV")
```

## jwt

```go
func main() {
	utils.LoadEnv()
	uuid, _ := uuid.NewRandom()

	user := jwt.User{
		ID:    uuid,
		Email: "aaronchen@jubo.health",
		Name:  "aaronchen",
	}

    // Generate JWT Token
	token := jwt.GenerateJWT(&user)
    fmt.Println("token", token)
    
    // Validate JWT Token 
	claims, err := jwt.ValidateJWT(token)
	if err != nil {
		fmt.Println("ValidateJWT error", err)
	}
	fmt.Println("claims", claims)
}
```

## password 加鹽與雜湊（hashed and salted）

```go
func main() {
	utils.LoadEnv()

	pass := "jubo01"
	digestedPassword := password.CreatePassword(pass, 10)

	isSamePassword := password.ComparePassword(digestedPassword, "jubo01")
	fmt.Println("isSamePassword", isSamePassword)
}
```

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

## enum with string value

### package mode

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

### main

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

