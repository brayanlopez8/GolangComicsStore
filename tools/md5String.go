package tools

import (
	"crypto/md5"
	"fmt"
)

func digestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}

//GenerateKeyMarvel get configuration portal marvel by api calls
func GenerateKeyMarvel() string {
	t := "1"
	privateKey := ReadConfig("PUBLICKEY") //"STRING DE LLAVE PRIVADA"
	publicKey := ReadConfig("PRIVATEKEY") //"STRING DE LLAVE PUBLICA"
	return digestString(t, publicKey, privateKey)
}
