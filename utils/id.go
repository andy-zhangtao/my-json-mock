package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// GenerateID 生成唯一ID
func GenerateID(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
