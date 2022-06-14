package go_futu_api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func RsaEncrypt(privateKey *rsa.PrivateKey, origData []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, origData)
}

func RsaDecrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext) //RSA算法解密
}

func GenRsaPair(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return priv, nil
}
