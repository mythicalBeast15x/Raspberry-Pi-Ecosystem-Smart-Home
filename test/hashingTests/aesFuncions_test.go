package Hashing

import (
	"CMPSC488SP24SecThursday/hashing"
	"testing"
)

func TestEncryptDecryptGCM(t *testing.T) {
	key := "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47" // 32 bytes for AES-256
	originalText := "Hello, World!"
	encryptedText, err := hashing.EncryptGCM(originalText, key) // Call EncryptGCM
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err) // Use Fatalf to stop the test if there's an error
	}
	decryptedText, err := hashing.DecryptGCM(encryptedText, key) // Call DecryptGCM
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	if decryptedText != originalText {
		t.Errorf("Decrypted text does not match original. Got %s, want %s", decryptedText,
			originalText)
	}
}
