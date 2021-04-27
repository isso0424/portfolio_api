package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

var secret string

func ValidateToken(token, key []byte) (result bool, err error) {
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

	result = string(plaintext) == secret

	return
}

func SetSecret(s string) (err error) {
	if secret != "" {
		err = errors.New("You can set secret only one time")
	}

	secret = s

	return
}
