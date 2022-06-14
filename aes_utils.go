package go_futu_api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AesEncrypt 加密函数
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

type AESCipher struct {
	encryptMode cipher.BlockMode
	decryptMode cipher.BlockMode
}

func NewAESCipher(key, iv []byte) (*AESCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptMode := cipher.NewCBCEncrypter(block, iv)

	decryptMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	return &AESCipher{encryptMode: encryptMode, decryptMode: decryptMode}, nil
}

func (a *AESCipher) Encrypt(plaintext []byte) []byte {
	plaintext = PKCS7Padding(plaintext, a.encryptMode.BlockSize())
	crypted := make([]byte, len(plaintext))
	a.encryptMode.CryptBlocks(crypted, plaintext)
	return crypted
}

func (a *AESCipher) Decrypt(ciphertext []byte) []byte {
	origData := make([]byte, len(ciphertext))
	a.decryptMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData
}
