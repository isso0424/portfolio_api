package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func parseToken(token []byte) (parsed string, err error) {
	ciphe, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(ciphe)
	if err != nil {
		return
	}

	nonceSize := gcm.NonceSize()
	if len(token) < nonceSize {
		err = errors.New("size of nonce is invalid")

		return
	}

	nonce, cipherText := token[:nonceSize], token[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return
	}

	parsed = string(plaintext)

	return
}
