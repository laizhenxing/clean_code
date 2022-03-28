package cypher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// HashPassword 计算密码哈希
func HashPassword(salt string, password string) string {
	return MD5(SHA256(password) + salt)
}

// MD5 计算字符串的md5哈希值
func MD5(s string) string {
	ret := md5.Sum([]byte(s))
	return hex.EncodeToString(ret[:])
}

// SHA1 计算字符串的sha1哈希值
func SHA1(s string) string {
	ret := sha1.Sum([]byte(s))
	return hex.EncodeToString(ret[:])
}

// SHA256 计算字符串的sha256哈希值
func SHA256(s string) string {
	ret := sha256.Sum256([]byte(s))
	return hex.EncodeToString(ret[:])
}

// SHA512 计算字符串的sha512哈希值
func SHA512(s string) string {
	ret := sha512.Sum512([]byte(s))
	return hex.EncodeToString(ret[:])
}

// EncodeToString base64加密
func EncodeToString(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// DecodeString base64解密
func DecodeString(encodeString string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	return decodeBytes, err
}
