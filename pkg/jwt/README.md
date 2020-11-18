# jwt

> [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

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
