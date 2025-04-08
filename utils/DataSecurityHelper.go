package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AESCipher struct {
	Key []byte
	Iv  []byte
}

func NewAESCipher(key, iv string) *AESCipher {
	return &AESCipher{
		Key: []byte(key), // 32 bytes
		Iv:  []byte(iv),  // 16 bytes
	}
}

func (ac *AESCipher) Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(ac.Key)
	if err != nil {
		return "", err
	}
	padText := pkcs7Pad([]byte(plainText), aes.BlockSize)
	cipherText := make([]byte, len(padText))
	mode := cipher.NewCBCEncrypter(block, ac.Iv)
	mode.CryptBlocks(cipherText, padText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (ac *AESCipher) Decrypt(cipherText string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(ac.Key)
	if err != nil {
		return "", err
	}
	plainText := make([]byte, len(decoded))
	mode := cipher.NewCBCDecrypter(block, ac.Iv)
	mode.CryptBlocks(plainText, decoded)
	plainText, err = pkcs7Unpad(plainText)
	return string(plainText), err
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	pad := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(pad)}, pad)...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	padLen := int(data[length-1])
	return data[:(length - padLen)], nil
}
