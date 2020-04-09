package v3

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 微信RSA签名与验签
// 参考：https://blog.csdn.net/heng4719/article/details/103216633

//商户私钥签名
func RsaSignWithSha256(data, priKey string) (string, error) {
	keyBytes := []byte(priKey)
	dataBytes := []byte(data)
	h := sha256.New()
	h.Write(dataBytes)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("RsaSignWithSha256 rsa key error")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

//平台公钥验签
func RsaVeryWithSha256(data, signature, pubKey string) (bool, error) {
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return false, errors.New("RsaVeryWithSha256 rsa key error")
	}
	oldSign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPKCS1v15(pk.(*rsa.PublicKey), crypto.SHA256, hashed[:], oldSign)
	if err != nil {
		return false, err
	}
	return true, nil
}
