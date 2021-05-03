package auth

import "errors"

func SetSecret(s string) (err error) {
	if secret != "" {
		err = errors.New("You can set secret only one time")
	}

	secret = s

	return
}

func SetKey(k string) (err error) {
	if len(key) != 0 {
		err = errors.New("You can set key only one time")
	}

	key = []byte(k)

	return
}
