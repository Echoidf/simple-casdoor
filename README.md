## 介绍

本项目对casdoor进行了精简，符合OIDC协议，实现了密码登录和短信登录

## 注意事项

fork 本项目后，注意在 `object.jwks.go` 中需要初始化 cert, 启动项目后在 `object.jwks.go` 程序初始化时会查找项目根目录下的两个文件（用于OIDC服务发现的private key 和 certificate），如果没有会在初始化时生成这两个文件

```go
func init() {
	cert = new(Cert)
	permPath, _ := util.GetAbsolutePath("token_jwt_key.pem")
	keyPath, _ := util.GetAbsolutePath("token_jwt_key.key")

	pemData, err := getFileData(permPath)
	if err != nil {
		// 生成 cert
		certStr, keyStr := generateRsaKeys(4096, 20, "fbCert", "fireboom")
		cert.Certificate = []byte(certStr)
		cert.PrivateKey = []byte(keyStr)
		// 生成文件
		go generateCertFile(permPath, cert.Certificate)
		go generateCertFile(keyPath, cert.PrivateKey)
		log.Info("已生成cert文件...")
	} else {
		cert.Certificate = pemData
		key, err := getFileData(keyPath)
		if err != nil {
			log.Error(err.Error())
		}
		cert.PrivateKey = key
	}
}
```


