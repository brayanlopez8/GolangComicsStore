package tools

import ()

func digestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}

//GenerateKeyMarvel get configuration portal marvel by api calls
func GenerateKeyMarvel() string {
	t := "1"
	privateKey := "STRING DE LLAVE PRIVADA"
	publicKey := "STRING DE LLAVE PUBLICA"
	return digestString(t, publicKey, privateKey)
}
