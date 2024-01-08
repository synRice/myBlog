// util/util.go
package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() string {
	// 生成随机的 Token
	token := make([]byte, 32)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}
