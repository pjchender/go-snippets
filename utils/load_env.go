package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv will load all variables in .env
// utils.LoadEnv()
// myEnv := os.Getenv("MY_ENV")
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Println("env loaded...")
}
