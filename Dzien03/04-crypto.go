package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("hello world")
	md1 := md5.Sum(data)
	fmt.Printf("MD5: %x\n", md1)

	md2 := sha256.Sum256(data)
	fmt.Printf("SHA256: %x\n", md2)

	md3 := sha512.Sum512(data)
	fmt.Printf("SHA512: %x\n", md3)

	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Losowy klucz: %x\n", key)

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decoded))

}
