package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateSecret(s int) (string, error) {
	b, err := generateRandomBytes(s)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func main() {
	// Parse command-line arguments
	length := flag.Int("length", 32, "Length of the secret to generate")
	flag.Parse()

	secret, err := generateSecret(*length)
	if err != nil {
		fmt.Println("Error generating secret:", err)
		return
	}

	fmt.Println(secret)
}
