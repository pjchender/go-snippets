# GoDotEnv

> [joho/godotenv](https://pkg.go.dev/github.com/joho/godotenv)

```go
import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  secretKey := os.Getenv("SECRET_KEY")
}
```
