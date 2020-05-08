package utils

import (
	"crypto/md5"
	"encoding/hex"

)

// MD5 加密算法，将密码进行加密后保存

func MD5(str string) string  {
	pass:=md5.New()
	pass.Write([]byte(str))
	return hex.EncodeToString(pass.Sum(nil))
}