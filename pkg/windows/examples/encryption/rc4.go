package main

import (
	"fmt"
	"log"

	"github.com/bluesentinelsec/OffensiveGoLang/pkg/windows/exfil"
)

func main() {
	srcData := []byte("Secret Message")
	key := []byte("password123")

	// encrypt data
	cipherText, err := exfil.CryptRC4(key, srcData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted: ", string(cipherText))

	// decrypt data
	plaintextData, err := exfil.CryptRC4(key, cipherText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted: ", string(plaintextData))

}
