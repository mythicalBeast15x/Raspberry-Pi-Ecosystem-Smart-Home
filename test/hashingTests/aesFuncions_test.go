package hashingTests

/*
func TestEncryptDecrypt(t *testing.T) {
	// Test cases
	testCases := []struct {
		plaintext string
	}{
		{plaintext: "Hello, world!"},
		{plaintext: "This is a test."},
		{plaintext: ""},
	}

	// Generate a random encryption key of length 16 bytes (128 bits)
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Error generating random key: %v", err)
=======
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
>>>>>>> ZigBee_Network_Group
	}

	// Iterate over test cases
	for _, tc := range testCases {
		plaintext := []byte(tc.plaintext)

		// Encrypt the plaintext
		ciphertextBase64, err := encrypt(plaintext, key)
		if err != nil {
			t.Errorf("Encryption error: %v", err)
		}

		// Decrypt the ciphertext
		decrypted, err := decrypt(ciphertextBase64, key)
		if err != nil {
			t.Errorf("Decryption error: %v", err)
		}

		// Compare plaintext and decrypted values
		if !bytes.Equal(plaintext, decrypted) {
			t.Errorf("Decrypt(Encrypt(plaintext)) = %s; want %s", decrypted, plaintext)
		}
	}
}

*/
