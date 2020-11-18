# Configor

> [github.com/jinzhu/configor](github.com/jinzhu/configor)

```go
import (
	"github.com/jinzhu/configor"
)

func main() {
	configuration := config.Get()

	data, _ := json.MarshalIndent(configuration, "", "  ")
	fmt.Println(string(data))
}
```
