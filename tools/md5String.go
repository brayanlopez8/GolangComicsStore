package tools

import (
	"crypto/md5"
	"fmt"
	"os"
)

func digestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}

//GenerateKeyMarvel get configuration portal marvel by api calls
func GenerateKeyMarvel() string {
	t := "1"
	privateKey := os.Getenv("PUBLICKEY")
	fmt.Println("PUBLICKEY: " + privateKey)
	publicKey := os.Getenv("PRIVATEKEY")
	fmt.Println("PRIVATEKEY: " + publicKey)
	return digestString(t, publicKey, privateKey)
}
