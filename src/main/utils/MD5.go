package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func VerifyMD5(sourceValue, encryptValue string) bool {
	md5Encrypt := EncodeMD5(sourceValue)
	if encryptValue == md5Encrypt {
		return true
	}
	return false
}
