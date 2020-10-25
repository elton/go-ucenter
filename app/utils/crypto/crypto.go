package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 text的MD5 签名
func MD5(text string) string {
	hashed := md5.New()
	hashed.Write([]byte(text))
	return hex.EncodeToString(hashed.Sum(nil))
}

// SHA256 text 的sha256签名
func SHA256(text string) string {
	hashed := sha256.New()
	hashed.Write([]byte(text))
	return hex.EncodeToString(hashed.Sum(nil))
}
