# openssl  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/openssl.svg)](https://github.com/golang-module/openssl/releases)
[![Go Build](https://github.com/golang-module/openssl/actions/workflows/bulid.yml/badge.svg)](https://github.com/golang-module/openssl/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/openssl)](https://goreportcard.com/report/github.com/golang-module/openssl)
[![codecov](https://codecov.io/gh/golang-module/openssl/branch/main/graph/badge.svg)](https://codecov.io/gh/golang-module/openssl)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/openssl)
![License](https://img.shields.io/github/license/golang-module/openssl)

English | [简体中文](README.cn.md)

#### Introduction
A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption

If you think it is helpful, please give me a star

[github.com/golang-module/openssl](https://github.com/golang-module/openssl "github.com/golang-module/openssl")

[gitee.com/go-package/openssl](https://gitee.com/go-package/openssl "gitee.com/go-package/openssl")

#### Installation
```go
// By github
go get -u github.com/golang-module/openssl

import (
    "github.com/golang-module/openssl"
)

// By gitee
go get -u gitee.com/go-package/openssl

import (
    "gitee.com/go-package/openssl"
)
```

#### Encode and encrypt

##### Encode by base32 
```go
// Encode by base32 from string and output string
openssl.Encode.FromString("hello world").ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// Encode by base32 from string and output byte slice
openssl.Encode.FromString("hello world").ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")

// Encode by base32 from byte slice and output string
openssl.Encode.FromBytes([]byte("hello world")).ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// Encode by base32 from byte slice and output byte slice
openssl.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")
```

##### Encode by base64
```go
// Encode by base64 from string and output string
openssl.Encode.FromString("hello world").ByBase64().ToString() // aGVsbG8gd29ybGQ=
// Encode by base64 from string and output byte slice
openssl.Encode.FromString("hello world").ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")

// Encode by base64 from byte slice and output string
openssl.Encode.FromBytes([]byte("hello world")).ByBase64().ToString() // aGVsbG8gd29ybGQ=
// Encode by base64 from byte slice and output byte slice
openssl.Encode.FromBytes([]byte("hello world")).ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")
```

##### Encode by hex
```go
// Encode by hex from string and output string
openssl.Encode.FromString("hello world").ByHex().ToString() // 68656c6c6f20776f726c64=
// Encode by hex from string and output byte slice
openssl.Encode.FromString("hello world").ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")

// Encode by hex from byte slice and output string
openssl.Encode.FromBytes([]byte("hello world")).ByHex().ToString() // 68656c6c6f20776f726c64
// Encode by hex from byte slice and output byte slice
openssl.Encode.FromBytes([]byte("hello world")).ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")
```

##### Encrypt by md5
```go
// Encrypt by md5 from string and output string with hex encoding
openssl.Encrypt.FromString("hello world").ByMd5().ToString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
openssl.Encrypt.FromString("hello world").ByMd5().ToString("hex") // 5eb63bbbe01eeed093cb22bb8f5acdc3
// Encrypt by md5 from string and output string with base32 encoding
openssl.Encrypt.FromString("hello world").ByMd5().ToString("base32") // L23DXO7AD3XNBE6LEK5Y6WWNYM======
// Encrypt by md5 from string and output string with base64 encoding
openssl.Encrypt.FromString("hello world").ByMd5().ToString("base64") // XrY7u+Ae7tCTyyK7j1rNww==

// Encrypt by md5 from string and output byte slice with hex encoding
openssl.Encrypt.FromString("hello world").ByMd5().ToBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
openssl.Encrypt.FromString("hello world").ByMd5().ToBytes("hex") // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// Encrypt by md5 from string and output byte slice with base32 encoding
openssl.Encrypt.FromString("hello world").ByMd5().ToBytes("base32") // []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")
// Encrypt by md5 from string and output byte slice with base64 encoding
openssl.Encrypt.FromString("hello world").ByMd5().ToBytes("base64") // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// Encrypt by md5 from byte slice and output string with hex encoding
openssl.Encrypt.FromBytes("hello world").ByMd5().ToString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
openssl.Encrypt.FromBytes("hello world").ByMd5().ToString("hex") // 5eb63bbbe01eeed093cb22bb8f5acdc3
// Encrypt by md5 from byte slice and output string with base32 encoding
openssl.Encrypt.FromBytes("hello world").ByMd5().ToString("base32") // L23DXO7AD3XNBE6LEK5Y6WWNYM======
// Encrypt by md5 from byte slice and output string with base64 encoding
openssl.Encrypt.FromBytes("hello world").ByMd5().ToString("base64") // XrY7u+Ae7tCTyyK7j1rNww==

// Encrypt by md5 from byte slice and output byte slice with hex encoding
openssl.Encrypt.FromBytes("hello world").ByMd5().ToBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
openssl.Encrypt.FromBytes("hello world").ByMd5().ToBytes("hex") // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// Encrypt by md5 from byte slice and output byte slice with base32 encoding
openssl.Encrypt.FromBytes("hello world").ByMd5().ToBytes("base32") // []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")
// Encrypt by md5 from byte slice and output byte slice with base64 encoding
openssl.Encrypt.FromBytes("hello world").ByMd5().ToBytes("base64") // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// Encrypt by md5 from file and output string with hex encoding
openssl.Encrypt.FromFile("./LICENSE")).ByMd5().ToString() // 014f03f9025ea81a8a0e9734be540c53
openssl.Encrypt.FromFile("./LICENSE")).ByMd5().ToString("hex") // 014f03f9025ea81a8a0e9734be540c53
// Encrypt by md5 from file and output string with base32 encoding
openssl.Encrypt.FromFile("./LICENSE")).ByMd5().ToString("base32") // AFHQH6ICL2UBVCQOS42L4VAMKM======
// Encrypt by md5 from file and output string with base64 encoding
openssl.Encrypt.FromFile("./LICENSE")).ByMd5().ToString("base64") // AU8D+QJeqBqKDpc0vlQMUw==

// Encrypt by md5 from file and output byte slice with hex encoding
openssl.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes() // []byte("014f03f9025ea81a8a0e9734be540c53")
openssl.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes("hex") // []byte("014f03f9025ea81a8a0e9734be540c53")
// Encrypt by md5 from file and output byte slice with base32 encoding

openssl.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes("base32") // []byte("AFHQH6ICL2UBVCQOS42L4VAMKM======")
// Encrypt by md5 from file and output byte slice with base64 encoding
openssl.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes("base64") // []byte("AU8D+QJeqBqKDpc0vlQMUw==")
```

##### Encrypt by sha1
```go
// Encrypt by sha1 from string and output string with hex encoding
openssl.Encrypt.FromString("hello world").BySha1().ToString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
openssl.Encrypt.FromString("hello world").BySha1().ToString("hex") // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// Encrypt by sha1 from string and output string with base32 encoding
openssl.Encrypt.FromString("hello world").BySha1().ToString("base32") // FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN
// Encrypt by sha1 from string and output string with base64 encoding
openssl.Encrypt.FromString("hello world").BySha1().ToString("base64") // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// Encrypt by sha1 from string and output byte slice with hex encoding
openssl.Encrypt.FromString("hello world").BySha1().ToBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
openssl.Encrypt.FromString("hello world").BySha1().ToBytes("hex") // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// Encrypt by sha1 from string and output byte slice with base32 encoding
openssl.Encrypt.FromString("hello world").BySha1().ToBytes("base32") // []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")
// Encrypt by sha1 from string and output byte slice with base64 encoding
openssl.Encrypt.FromString("hello world").BySha1().ToBytes("base64") // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")

// Encrypt by sha1 from byte slice and output string with hex encoding
openssl.Encrypt.FromBytes("hello world").BySha1().ToString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
openssl.Encrypt.FromBytes("hello world").BySha1().ToString("hex") // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// Encrypt by sha1 from byte slice and output string with base32 encoding
openssl.Encrypt.FromBytes("hello world").BySha1().ToString("base32") // FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN
// Encrypt by sha1 from byte slice and output string with base64 encoding
openssl.Encrypt.FromBytes("hello world").BySha1().ToString("base64") // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// Encrypt by sha1 from byte slice and output byte slice with hex encoding
openssl.Encrypt.FromBytes("hello world").BySha1().ToBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
openssl.Encrypt.FromBytes("hello world").BySha1().ToBytes("hex") // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// Encrypt by sha1 from byte slice and output byte slice with base32 encoding
openssl.Encrypt.FromBytes("hello world").BySha1().ToBytes("base32") // []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")
// Encrypt by sha1 from byte slice and output byte slice with base64 encoding
openssl.Encrypt.FromBytes("hello world").BySha1().ToBytes("base64") // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

##### Encrypt by sha224 
```go
// Encrypt by sha224 from string and output string with hex encoding
openssl.Encrypt.FromString("hello world").Sha224().ToString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
openssl.Encrypt.FromString("hello world").Sha224().ToString("hex") // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// Encrypt by sha224 from string and output string with base32 encoding
openssl.Encrypt.FromString("hello world").Sha224().ToString("base32") // F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===
// Encrypt by sha224 from string and output string with base64 encoding
openssl.Encrypt.FromString("hello world").Sha224().ToString("base64") // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// Encrypt by sha224 from string and output byte slice with hex encoding
openssl.Encrypt.FromString("hello world").Sha224().ToBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
openssl.Encrypt.FromString("hello world").Sha224().ToBytes("hex") // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// Encrypt by sha224 from string and output byte slice with base32 encoding
openssl.Encrypt.FromString("hello world").Sha224().ToBytes("base32") // []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")
// Encrypt by sha224 from string and output byte slice with base64 encoding
openssl.Encrypt.FromString("hello world").Sha224().ToBytes("base64") // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")

// Encrypt by sha224 from byte slice and output string with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha224().ToString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
openssl.Encrypt.FromBytes("hello world").Sha224().ToString("hex") // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// Encrypt by sha224 from byte slice and output string with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha224().ToString("base32") // F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===
// Encrypt by sha224 from byte slice and output string with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha224().ToString("base64") // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// Encrypt by sha224 from byte slice and output byte slice with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha224().ToBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
openssl.Encrypt.FromBytes("hello world").Sha224().ToBytes("hex") // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// Encrypt by sha224 from byte slice and output byte slice with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha224().ToBytes("base32") // []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")
// Encrypt by sha224 from byte slice and output byte slice with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha224().ToBytes("base64") // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

##### Encrypt by sha256 
```go
// Encrypt by sha256 from string and output string with hex encoding
openssl.Encrypt.FromString("hello world").Sha256().ToString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
openssl.Encrypt.FromString("hello world").Sha256().ToString("hex") // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by sha256 from string and output string with base32 encoding
openssl.Encrypt.FromString("hello world").Sha256().ToString("base32") // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// Encrypt by sha256 from string and output string with base64 encoding
openssl.Encrypt.FromString("hello world").Sha256().ToString("base64") // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by sha256 from string and output byte slice with hex encoding
openssl.Encrypt.FromString("hello world").Sha256().ToBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
openssl.Encrypt.FromString("hello world").Sha256().ToBytes("hex") // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by sha256 from string and output byte slice with base32 encoding
openssl.Encrypt.FromString("hello world").Sha256().ToBytes("base32") // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// Encrypt by sha256 from string and output byte slice with base64 encoding
openssl.Encrypt.FromString("hello world").Sha256().ToBytes("base64") // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")

// Encrypt by sha256 from byte slice and output string with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha256().ToString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
openssl.Encrypt.FromBytes("hello world").Sha256().ToString("hex") // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by sha256 from byte slice and output string with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha256().ToString("base32") // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// Encrypt by sha256 from byte slice and output string with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha256().ToString("base64") // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by sha256 from byte slice and output byte slice with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha256().ToBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
openssl.Encrypt.FromBytes("hello world").Sha256().ToBytes("hex") // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by sha256 from byte slice and output byte slice with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha256().ToBytes("base32") // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// Encrypt by sha256 from byte slice and output byte slice with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha256().ToBytes("base64") // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by sha384 
```go
// Encrypt by sha384 from string and output string with hex encoding
openssl.Encrypt.FromString("hello world").Sha384().ToString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
openssl.Encrypt.FromString("hello world").Sha384().ToString("hex") // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// Encrypt by sha384 from string and output string with base32 encoding
openssl.Encrypt.FromString("hello world").Sha384().ToString("base32") // 7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===
// Encrypt by sha384 from string and output string with base64 encoding
openssl.Encrypt.FromString("hello world").Sha384().ToString("base64") // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// Encrypt by sha384 from string and output byte slice with hex encoding
openssl.Encrypt.FromString("hello world").Sha384().ToBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
openssl.Encrypt.FromString("hello world").Sha384().ToBytes("hex") // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// Encrypt by sha384 from string and output byte slice with base32 encoding
openssl.Encrypt.FromString("hello world").Sha384().ToBytes("base32") // []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")
// Encrypt by sha384 from string and output byte slice with base64 encoding
openssl.Encrypt.FromString("hello world").Sha384().ToBytes("base64") // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")

// Encrypt by sha384 from byte slice and output string with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha384().ToString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
openssl.Encrypt.FromBytes("hello world").Sha384().ToString("hex") // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// Encrypt by sha384 from byte slice and output string with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha384().ToString("base32") // 7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===
// Encrypt by sha384 from byte slice and output string with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha384().ToString("base64") // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// Encrypt by sha384 from byte slice and output byte slice with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha384().ToBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
openssl.Encrypt.FromBytes("hello world").Sha384().ToBytes("hex") // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// Encrypt by sha384 from byte slice and output byte slice with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha384().ToBytes("base32") // []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")
// Encrypt by sha384 from byte slice and output byte slice with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha384().ToBytes("base64") // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

##### Encrypt by sha512
```go
// Encrypt by sha512 from string and output string with hex encoding
openssl.Encrypt.FromString("hello world").Sha512().ToString() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
openssl.Encrypt.FromString("hello world").Sha512().ToString("hex") // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// Encrypt by sha512 from string and output string with base32 encoding
openssl.Encrypt.FromString("hello world").Sha512().ToString("base32") // GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=
// Encrypt by sha512 from string and output string with base64 encoding
openssl.Encrypt.FromString("hello world").Sha512().ToString("base64") // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// Encrypt by sha512 from string and output byte slice with hex encoding
openssl.Encrypt.FromString("hello world").Sha512().ToBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
openssl.Encrypt.FromString("hello world").Sha512().ToBytes("hex") // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// Encrypt by sha512 from string and output byte slice with base32 encoding
openssl.Encrypt.FromString("hello world").Sha512().ToBytes("base32") // []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")
// Encrypt by sha512 from string and output byte slice with base64 encoding
openssl.Encrypt.FromString("hello world").Sha512().ToBytes("base64") // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")

// Encrypt by sha512 from byte slice and output string with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha512().ToString() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
openssl.Encrypt.FromBytes("hello world").Sha512().ToString("hex") // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// Encrypt by sha512 from byte slice and output string with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha512().ToString("base32") // GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=
// Encrypt by sha512 from byte slice and output string with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha512().ToString("base64") // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// Encrypt by sha512 from byte slice and output byte slice with hex encoding
openssl.Encrypt.FromBytes("hello world").Sha512().ToBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
openssl.Encrypt.FromBytes("hello world").Sha512().ToBytes("hex") // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// Encrypt by sha512 from byte slice and output byte slice with base32 encoding
openssl.Encrypt.FromBytes("hello world").Sha512().ToBytes("base32") // []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")
// Encrypt by sha512 from byte slice and output byte slice with base64 encoding
openssl.Encrypt.FromBytes("hello world").Sha512().ToBytes("base64") // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")
```

#### Decode and decrypt
##### Decode by base32
```go
// Decode by base32 from string and output string
openssl.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToString() // hello world
// Decode by base32 from string and output byte slice
openssl.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToBytes() // []byte("hello world")

// Decode by base32 from byte slice and output string
openssl.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToString() // hello world
// Decode by base32 from byte slice and output byte slice
openssl.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToBytes() // []byte("hello world")
```

##### Decode by base64
```go
// Decode by base64 from string and output string
openssl.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToString() // hello world
// Decode by base64 from string and output byte slice
openssl.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToBytes() // []byte("hello world")

// Decode by base64 from byte slice and output string
openssl.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToString() // hello world
// Decode by base64 from byte slice and output byte slice
openssl.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToBytes() // []byte("hello world")
```

##### Decode by hex
```go
// Decode by hex from string and output string
openssl.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToString() // hello world
// Decode by hex from string and output byte slice
openssl.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToBytes() // []byte("hello world")

// Decode by hex from byte slice and output string
openssl.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToString() // hello world
// Decode by hex from byte slice and output byte slice
openssl.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToBytes() // []byte("hello world")
```

#### Error handling
> If more than one error occurs, only the first error is returned

```go
e := openssl.Encrypy.FromFile("./demo.txt").ByMd5()
if e.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToString())
// Output
invalid file "./demo.txt", please make sure the file exists
```

#### Todo List
 - [ ] Encryption and decryption by aes
 - [ ] Encryption and decryption by des
 - [ ] Encryption and decryption by 3aes
 - [ ] Encryption and decryption by rsa
 - [ ] Encryption and decryption by rc2
 - [ ] Encryption and decryption by rc4
 - [ ] Encryption and decryption by rc5
 - [ ] Encryption and decryption by rc6