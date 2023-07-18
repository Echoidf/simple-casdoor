package object

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/square/go-jose.v2"
	"io"
	"os"
)

type Cert struct {
	Certificate []byte `json:"certificate"`
	PrivateKey  []byte `json:"privateKey"`
}

var cert *Cert

func init() {
	cert = new(Cert)
	pemData, err := getFileData("../token_jwt_key.pem")
	if err != nil {
		log.Error(err.Error())
	}
	cert.Certificate = pemData

	key, err := getFileData("../token_jwt_key.key")
	if err != nil {
		log.Error(err.Error())
	}
	cert.PrivateKey = key
}

func GetJsonWebKeySet() (jose.JSONWebKeySet, error) {
	jwks := jose.JSONWebKeySet{}

	certPemBlock := cert.Certificate
	certDerBlock, _ := pem.Decode(certPemBlock)
	x509Cert, _ := x509.ParseCertificate(certDerBlock.Bytes)

	var jwk jose.JSONWebKey
	jwk.Key = x509Cert.PublicKey
	jwk.Certificates = []*x509.Certificate{x509Cert}
	jwk.KeyID = "fireboom"
	jwk.Algorithm = "RS256"
	jwk.Use = "sig"

	jwks.Keys = append(jwks.Keys, jwk)

	return jwks, nil
}

func getFileData(filePath string) ([]byte, error) {
	// 打开文本文件进行读取
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return nil, err
	}
	defer file.Close()

	// 读取文件的所有数据
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件时出错:", err)
		return nil, err
	}

	// 将数据转换为字节数组
	return data, nil
}
