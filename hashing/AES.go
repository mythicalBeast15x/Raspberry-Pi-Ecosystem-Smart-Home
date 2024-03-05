package hashing

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"io"
)

func generateRandomKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// CBC mode requires an initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Pad plaintext if needed
	plaintext = PKCS7Padding(plaintext, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	copy(ciphertext[:aes.BlockSize], iv)

	// Encode ciphertext to Base64
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	return ciphertextBase64, nil
}

func decrypt(ciphertextBase64 string, key []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// Unpad plaintext
	plaintext = PKCS7Unpadding(plaintext)

	return plaintext, nil
}

// PKCS7Padding pads data using the PKCS#7 scheme
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7Unpadding removes padding from data that was padded using the PKCS#7 scheme
func PKCS7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func loadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

/*
func main() {
	// Load environment variables from .env file
	err := loadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	// Generate a random encryption key of length 16 bytes (128 bits)
	key, err := generateRandomKey(16)
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}

	// Use the generated key in your encryption/decryption logic
	fmt.Println("Generated key:", key)
	// Example usage: Encrypt and decrypt data
	data := []byte("Hello, AES encryption!")

	fmt.Println("Plaintext:", string(data))

	// Encrypt the data
	ciphertextBase64, err := encrypt(data, []byte(key))
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Ciphertext:", ciphertextBase64)

	// Decrypt the ciphertext
	plaintext, err := decrypt(ciphertextBase64, []byte(key))
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted:", string(plaintext))
}
*/
