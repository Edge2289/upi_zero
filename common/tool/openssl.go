package tool

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// RSASign 私钥签名
func RSASign(data []byte, privateKeyStr string) (string, error) {
	// 1、选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	// 2、读取私钥文件，解析出私钥对象
	privateKey, err := readParsePrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	// 3、RSA数字签名（参数是随机数、私钥对象、哈希类型、签名文件的哈希串，生成bash64编码）
	bytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, myhash, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// readParsePrivateKey 读取私钥文件，解析出私钥对象
func readParsePrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 解析DER编码的私钥，生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// RSAVerify 公钥验证
func RSAVerify(data []byte, base64Sig, publicKeyStr string) error {
	// 1、对base64编码的签名内容进行解码，返回签名字节
	bytes, err := base64.StdEncoding.DecodeString(base64Sig)
	if err != nil {
		return err
	}
	// 2、选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	// 3、读取公钥文件，解析出公钥对象
	publicKey, err := readParsePublicKey(publicKeyStr)
	if err != nil {
		return err
	}
	// 4、RSA验证数字签名（参数是公钥对象、哈希类型、签名文件的哈希串、签名后的字节）
	return rsa.VerifyPKCS1v15(publicKey, myhash, hashed, bytes)
}

// readParsePublicKey 读取公钥文件，解析公钥对象
func readParsePublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	// 解码公钥字节，生成加密对象
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	// 解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}
