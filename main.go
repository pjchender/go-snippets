package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/jwt"
	"github.com/pjchender/go-snippets/utils"
)

func main() {
	utils.LoadEnv()
	userID, _ := uuid.NewRandom()

	user := jwt.User{
		ID:    userID,
		Email: "aaronchen@jubo.health",
		Name:  "aaronchen",
	}

	token := jwt.GenerateJWT(&user)
	fmt.Println(token)

	claims, err := jwt.ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwOGE5ZDc1Ni1iMmY4LTQxYWItYjVmMy0xNjA3NmJkZjBjODgiLCJlbWFpbCI6ImFhcm9uY2hlbkBqdWJvLmhlYWx0aCIsIm5hbWUiOiJhYXJvbmNoZW4iLCJpc3MiOiJKdWJvLCBJbmMuIn0.J6CX_2vIFcYRJ0aTlnGNjJy7-OkfBXP7jTrauAd6fG4")
	if err != nil {
		fmt.Println("jwt.ValidateJWT", err)
	}

	fmt.Println(claims)
}
