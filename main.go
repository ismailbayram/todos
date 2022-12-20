package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	password := "123456"
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := h.Sum(nil)

	fmt.Println(hashedPassword)
	fmt.Println(string(hashedPassword))
	fmt.Println(hex.EncodeToString(hashedPassword))
}
