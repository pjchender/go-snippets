// Utils package provide some utilities

package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv will load all variables in .env
// utils.LoadEnv()
// myEnv := os.Getenv("MY_ENV")
func LoadEnv(filepath string) {
	err := godotenv.Load(filepath)
	if err != nil {
		panic(err)
	}
	log.Println("env loaded...")
}
