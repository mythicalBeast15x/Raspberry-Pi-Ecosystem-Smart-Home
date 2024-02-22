package Hashing

import (
	"CMPSC488SP24SecThursday/hashing"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47" // 32 bytes for AES-256
	originalText := "Hello, World!"
	encryptedText, err := hashing.Encrypt(originalText, key)
	if err != nil {
		t.Errorf("Failed to encrypt: %v", err)
	}
	decryptedText, err := hashing.Decrypt(encryptedText, key)
	if err != nil {
		t.Errorf("Failed to decrypt: %v", err)
	}
	if decryptedText != originalText {
		t.Errorf("Decrypted text does not match original. Got %s, want %s", decryptedText,
			originalText)
	}
}
