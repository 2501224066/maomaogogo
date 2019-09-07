package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Md5 加密
func Md5(param string) string {
	h := md5.New()
	h.Write([]byte(param))
	return hex.EncodeToString(h.Sum(nil))
}

// RandNum 随机数
func RandNum(len int) string {
	arr := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())

	var str strings.Builder
	for i := 0; i < len; i++ {
		fmt.Fprintf(&str, "%d", arr[rand.Intn(10)])
	}
	return str.String()
}
