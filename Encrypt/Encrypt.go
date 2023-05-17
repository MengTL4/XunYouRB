package Encrypt

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
)

// DecryptDESECB Decrypt 解密函数
func DecryptDESECB(d, key []byte) string {
	data, _ := base64.StdEncoding.DecodeString(string(d))
	if len(key) > 8 {
		key = key[:8]
	}
	block, _ := des.NewCipher(key)
	bs := block.BlockSize()
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Unpad(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

func ByteBase64(bArr []byte) []byte {
	encoded := base64.StdEncoding.EncodeToString(bArr)
	return []byte(encoded)
}

func ByteMd5(bArr []byte) []byte {
	encoded := md5.Sum(bArr)
	return encoded[:]
}

func EncodeHex(bArr []byte) []byte {
	length := len(bArr)
	cArr := make([]byte, length*2)
	i := 0
	for i2 := 0; i2 < length; i2++ {
		i3 := i + 1
		cArr[i] = "0123456789ABCDEF"[(bArr[i2]&240)>>4]
		i = i3 + 1
		cArr[i3] = "0123456789ABCDEF"[bArr[i2]&15]
	}
	return cArr
}

// Encrypt param参数加密
func Encrypt(data, key []byte) string {
	if len(key) > 8 {
		key = key[:8]
	}
	block, _ := des.NewCipher(key)
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Sign(signStr string) string {
	return string(EncodeHex(ByteMd5(ByteBase64([]byte(signStr)))))
}

func Param(paramStr, keyParam string) string {
	return Encrypt([]byte(paramStr), []byte(keyParam))
}
