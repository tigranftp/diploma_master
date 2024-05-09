package passwords

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func HashPassword(plainPassword string) string {
	salt := make([]byte, 8)
	rand.Read(salt)
	return fmt.Sprintf("%x", hashPass(salt, plainPassword))
}

func hashPass(salt []byte, plainPassword string) []byte {
	res := make([]byte, len(salt))
	_ = copy(res, salt)
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(res, hashedPass...)
}

func CheckPassword(passHash []byte, plainPassword string) bool {
	passHash, _ = hex.DecodeString(string(passHash))
	salt := passHash[0:8]
	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}
