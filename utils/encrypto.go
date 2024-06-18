package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
)

func encrypt(message string) string {
	sha := sha512.New()
	_, err := io.WriteString(sha, message)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(sha.Sum(nil))
}
