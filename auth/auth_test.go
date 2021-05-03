package auth_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/isso0424/portfolio_api/auth"
	"github.com/stretchr/testify/assert"
)

func encrypt(text, key []byte) (data []byte, err error) {
	ciph, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, e := io.ReadFull(rand.Reader, nonce); e != nil {
		e = err
		return
	}

	data = gcm.Seal(nonce, nonce, text, nil)

	return
}

func TestAuthorization(t *testing.T) {
	_, dir, _, _ := runtime.Caller(0)
	p := filepath.Join(filepath.Dir(dir), "test_key.dat")
	fp, err := os.OpenFile(p, os.O_RDONLY, 0755)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		t.Fatal(err)
	}

	key := string(data[:len(data) - 1])
	err = auth.SetKey(key)
	if err != nil {
		t.Fatal(err)
	}

	secret := "password"
	err = auth.SetSecret(secret)
	if err != nil {
		t.Fatal(err)
	}

	rawToken := fmt.Sprintf("%s%d", secret, time.Now().Add(20 * time.Minute).Unix())

	token, err := encrypt([]byte(rawToken), []byte(key))
	if err != nil {
		t.Fatal(err)
	}

	result, err := auth.ValidateToken(string(token))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, result)

	incorrectToken := fmt.Sprintf("%s%d", "invalid", time.Now().Add(20 * time.Minute).Unix())
	token, err = encrypt([]byte(incorrectToken), []byte(key))
	result, err = auth.ValidateToken(string(token))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, result)

	revokedToken := fmt.Sprintf("%s%d", secret, time.Now().Add(-20 * time.Minute).Unix())
	token, err = encrypt([]byte(revokedToken), []byte(key))
	result, err = auth.ValidateToken(string(token))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, false, result)
}
