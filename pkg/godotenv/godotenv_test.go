package godotenv

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func ExampleLoadEnv() {
	LoadEnv("./../.env")
	myEnv := os.Getenv("MY_ENV")
	fmt.Println(myEnv)
	// Output: Gopher
}

func TestLoadEnv(t *testing.T) {
	LoadEnv("./../.env")
	myEnv := os.Getenv("MY_ENV")
	assert.Equal(t, "Gopher", myEnv)
}
