## 介绍

本项目对casdoor进行了精简，符合OIDC协议，实现了密码登录和短信登录

## 如何启动

fork 本项目后，注意在 `object.jwks.go` 中需要初始化 cert:

```go
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
```

此处依赖的两个文件需自己生成（private key 和 CERTIFICATE），且需满足 RSA 算法

如有需要可参考以下代码进行生成：

```go
import (
"crypto/rand"
"crypto/rsa"
"crypto/x509"
"crypto/x509/pkix"
"encoding/pem"
"math/big"
"time"
)


func generateRsaKeys(bitSize int, expireInYears int, commonName string, organization string) (string, string) {
	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	// Encode private key to PKCS#1 ASN.1 PEM.
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	tml := x509.Certificate{
		// you can add any attr that you need
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(expireInYears, 0, 0),
		// you have to generate a different serial number each execution
		SerialNumber: big.NewInt(123456),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{organization},
		},
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	return string(certPem), string(privateKeyPem)
}
```

