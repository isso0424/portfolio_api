package auth

import (
	"strconv"
	"time"
)

var (
	secret string
	key    []byte
)

func ValidateToken(token string) (result bool, err error) {
	parsed, err := parseToken([]byte(token))
	if err != nil {
		return
	}

	secretLen := len(secret)
	if len(parsed) < secretLen+1 {
		return
	}

	target, epochStr := parsed[:secretLen], parsed[secretLen:]

	epoch, err := strconv.ParseInt(epochStr, 10, 64)
	if err != nil {
		return
	}

	nowEpoch := time.Now().Unix()

	result = epoch > nowEpoch && target == secret

	return
}
