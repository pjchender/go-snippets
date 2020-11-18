# Generate and Compare Password with bcrypt

- password 加鹽與雜湊（hashed and salted）

```go
func main() {
	utils.LoadEnv()

	pass := "jubo01"
	digestedPassword := password.CreatePassword(pass, 10)

	isSamePassword := password.ComparePassword(digestedPassword, "jubo01")
	fmt.Println("isSamePassword", isSamePassword)
}
```
