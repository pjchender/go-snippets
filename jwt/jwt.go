package jwt

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Email string
	Name  string
}

type UserClaims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

func GenerateJWT(user *User) string {
	jwtSecret := os.Getenv("JWT_SECRET")

	claims := UserClaims{
		UserID: user.ID.String(),
		Email:  user.Email,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Jubo Inc.",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		panic(err)
	}

	return signedToken
}

// ValidateJWT 會驗證 JWT 是否合法
// 如果不是合法的 JWT 或已經過了 ExpiresAt 的話，使用 jwt.ParseWithClaims 都會直接報 Error，且 token.Valid 為 false
// 也可以使用 claims.Valid() 判斷 claims 是否有效
func ValidateJWT(authToken string) (*UserClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	// 如果 token 有錯或過期的話，err 就不會是 nil
	token, err := jwt.ParseWithClaims(authToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("invalid signature method")
			return nil, errors.New("invalid authorization token")
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid authorization token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !(ok && token.Valid) {
		log.Println("invalid authorization token")
		return nil, errors.New("invalid authorization token")
	}

	return claims, nil
}
