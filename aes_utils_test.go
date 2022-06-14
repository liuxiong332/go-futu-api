package go_futu_api

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("hello ming")

	c := make([]byte, aes.BlockSize+len(plaintext))
	iv := c[:aes.BlockSize]

	//加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	//打印加密base64后密码
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	//解密
	plaintext, err = AesDecrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}

	if string(plaintext) != "hello ming" {
		t.Error("not equal")
	}

	aesCipher, _ := NewAESCipher(key, iv)

	ciper2 := aesCipher.Encrypt([]byte("hello ming"))
	fmt.Println(base64.StdEncoding.EncodeToString(ciper2))

	plain2 := aesCipher.Decrypt(ciper2)
	fmt.Println(string(plain2))

	if string(plain2) != "hello ming" {
		t.Error("not equal")
	}
}
