package utils

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"net/url"
	// "strings"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

var key = MD5byte("orange")
var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func MD5byte(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

// md5加密
func MD5(s string) string {
	return hex.EncodeToString(MD5byte(s))
}

// sha1加密
func SHA1(s string) string {
	return hex.EncodeToString(SHA1Byte(s))
}

func SHA1Byte(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

// Base64编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

// AES编码
func AesEncode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBEncrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}

// AES解码
func AesDecode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBDecrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}

// utf-8转gbk
func Utf8ToGBK(str string) string {
	// 字符集转换
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(str)
}

// gbk转utf-8
func GBKToUtf8(str string) string {
	// 字符集转换
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(str)
}

// url编码
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}
