package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(PublicKeyStr))
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	fmt.Println(PrivateKeyStr)
	//解密
	block, _ := pem.Decode([]byte(PrivateKeyStr))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	data,err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return data,nil
}

var PrivateKeyStr string
var PublicKeyStr string
func GenRsaKey(bits int) error {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
		priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	PrivateKeyStr = string(pem.EncodeToMemory(priBlock))
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	PublicKeyStr = string(pem.EncodeToMemory(publicBlock))

	if err != nil {
		return err
	}
	return nil
}