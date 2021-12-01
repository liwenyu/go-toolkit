package algo

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 algo
// return 32 bit
func MD5(str string) string {
	return MD5Encrypt(str)
}

// md5 algo
// return 16 bit
func Md516bit(str string) string {
	res := MD5Encrypt(str)
	if len(res) == 0 {
		return ""
	}
	return res[8:24]
}

// MD5 algo
func MD5Encrypt(str string) string {
	if len(str) == 0 {
		return ""
	}
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
