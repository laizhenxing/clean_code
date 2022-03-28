package str

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

const (
	// MetaCharacter 包含了数字和大小写字母的元字符
	MetaCharacter = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MetaNumber    = "1234567890"
)

// MakeRandomString 生成指定长度的随机字符串
func MakeRandomString(n int64) string {
	str := MetaCharacter
	var salt string
	rand.Seed(time.Now().Unix())
	for n > 0 {
		salt += string(str[rand.Intn(len(str)-1)])
		n--
	}
	return salt
}

//MakePassSalt 生成用于密码加密的salt字符串
func MakePassSalt() string {
	return MakeRandomString(10)
}

// RandomDigit 返回随机的数字字符串
func RandomDigit(n int) string {
	var letterRunes = []rune(MetaNumber)

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// SubString 字符串截取
func SubString(s string, len int64) string {
	var length int
	strlen := utf8.RuneCountInString(s)
	if int64(strlen) > len {
		length = int(len)
	} else {
		length = strlen
	}

	return string([]rune(s)[:length])
}

// Int32ToString 32位整形转字符串
func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
