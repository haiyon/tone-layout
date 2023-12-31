package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - hashes a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash - checks if a password matches a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// PKCS7Padding - PKCS7 padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding - PKCS7 unpadding
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt - AES Encrypt
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt - AES Decrypt
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

var aes_key = []byte("DB+XPxNgDWNe3ydQqQyrGYtiStqYjA==")

// AesEncrypt64 - AES Encrypt base64
func AesEncrypt64(origData []byte) string {
	crypted, err := AesEncrypt(origData, aes_key)
	if err != nil {
		fmt.Println("AesEncrypt64 err:", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(crypted)
}

// AesDecrypt64 - AES Decrypt base64
func AesDecrypt64(cryptedStr string) string {
	crypted, err := base64.StdEncoding.DecodeString(cryptedStr)
	if err != nil {
		fmt.Println("AesDecrypt64 err:", err)
		return ""
	}
	origData, err := AesDecrypt(crypted, aes_key)
	if err != nil {
		fmt.Println("AesDecrypt64 err:", err)
		return ""
	}
	return string(origData)
}
