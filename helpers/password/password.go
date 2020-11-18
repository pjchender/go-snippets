package password

import (
	"bytes"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// CreatePassword returns a hashed version of the given password.
func CreatePassword(password string, strength int) string {
	saltedPassword := addSalt(password)

	digestedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, strength)
	if err != nil {
		panic(err)
	}
	return string(digestedPassword)
}

// ComparePassword compares a hashed password with its possible plaintext equivalent.
func ComparePassword(digestedPassword, password string) bool {
	if digestedPassword == "" {
		return false
	}

	saltedPassword := addSalt(password)

	err := bcrypt.CompareHashAndPassword([]byte(digestedPassword), saltedPassword)

	return err == nil
}

func addSalt(password string) []byte {
	passwordBuf := bytes.Buffer{}
	passwordBuf.WriteString(password)

	salt := os.Getenv("PASSWORD_SALT")
	if salt != "" {
		passwordBuf.WriteString(salt)
	}

	return passwordBuf.Bytes()
}
