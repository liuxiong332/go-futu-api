package go_futu_api

import (
	"testing"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCgPwlGJrWqTYaoMkI8jXkEI8ewQ7E57G2Fi91WTXMMK7X6GsT9
VmnRcq++Rk/VS+4IPBlfWyVRg0NfQDyuKjed21fUPa9AIbpYWHgP/tojyeYC1+Ra
Xncrt9kLp7nW4FZMJmzwU9hfxIB0nhDQqhJenjdBZuYZfkICfMqyqbVkAwIDAQAB
AoGAJRcSDXOuPrHdBhdD74ILTaL+eFTis3Z+zxdVbsFUbK+9WhtSFxUmPv1dohvi
JIuDl9JZSRHurFRGhsh2gxVwc7JXwWfD0DmD8dvdzr8q85Jml9YVZ7uhHFqSO4cY
I7dlBOd7Uwjnc39E/d+1E/kWVNfKt7opPHgt02zOHLSxkbECQQDS7H3myu3oLOi0
Slpd1MmmHVOo2cqJ1b3H6E8JtEjmHGswWTYvQNAe4yZ+Kffsp5LUYujedncPKvEj
4G+iz44bAkEAwn4Bx30FKTri/tybgSnCWKwTGSX479829Xucrm5pYU/3D5/PeJQL
Ra4YSyg2/hU3ZBrue6CdzYJgGXNGEWhAOQJBALMlOB4A96X+FruidzRA2fBj8j10
lakSSHl1H0RfwpbnRkcvTm0+AEZrqbL4lGGFRplrVNw2BBN25o8RPeArp0cCQEhu
kw0PI1fqhVUzJXqh6a4KT4aDHMWAlMAxi/VuSzKhjDo2Yxbd06DcqFF9JZXUou9W
FFDYTUyW7GEuC/85mwkCQCOEjUQX0C3JCSr6fyZIjpEr+znyc9eFHyBp+533Ur4g
eFu2ewJ3ufJiUBmEj1rEQku8W7h9DS2rXl10IiSwUAA=
-----END RSA PRIVATE KEY-----
`)

func TestEncrypt(t *testing.T) {
	priv, err := GenRsaPair(privateKey)
	if err != nil {
		t.Error(err)
	}
	res, err := RsaEncrypt(priv, []byte("Hello World"))
	if err != nil {
		t.Error(err)
	}

	decryptRes, err := RsaDecrypt(priv, res)
	if err != nil {
		t.Error(err)
	}

	if string(decryptRes) != "Hello World" {
		t.Error("not equal")
	}
}
