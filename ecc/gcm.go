package ecc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

func PointTo256bit(point string) [32]byte {
	return sha256.Sum256([]byte(point))
}

// EncryptGCM Encryption using GCM. Returns the encrypted string and the nonce.
func EncryptGCM(text string, secretKey [32]byte) (string, []byte, error) {
	block, err := aes.NewCipher(secretKey[:])
	if err != nil {
		return "", nil, err
	}

	rawText := []byte(text)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, rawText, nil)
	return string(ciphertext), nonce, err
}

// DecryptGCM decryption using GCM. Returns the decrypted text.
func DecryptGCM(cipherText string, nonce []byte, secretKey [32]byte) (string, error) {
	block, err := aes.NewCipher(secretKey[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	text, err := gcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return "", err
	}
	return string(text), nil
}


