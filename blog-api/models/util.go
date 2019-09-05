package models

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/astaxie/beego/logs"
	"github.com/gofrs/uuid"
)

/*
Message 返回格式
*/
type Message struct {
	Code   int         `json:"code"`
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

/*
UID 随机生成id
*/
func UID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		logs.Info(err)
	}
	return uid.String()
}

/*
Encrypt 加密密码
*/
func Encrypt(text string) (string, error) {
	key := []byte{0xBA, 0x47, 0x2F, 0x02, 0xC8, 0x92, 0x1F, 0x7D,
		0x2A, 0x3D, 0x8F, 0x06, 0x41, 0x9B, 0x6F, 0x2D,
		0xBA, 0x36, 0x6F, 0x07, 0xC7, 0x52, 0x1F, 0x7D,
		0x4A, 0x5D, 0x4F, 0x06, 0x45, 0x8B, 0x3F, 0x4D,
	}
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

/*
Decrypt 解密密码
*/
func Decrypt(encrypted string) (string, error) {
	key := []byte{0xBA, 0x47, 0x2F, 0x02, 0xC8, 0x92, 0x1F, 0x7D,
		0x2A, 0x3D, 0x8F, 0x06, 0x41, 0x9B, 0x6F, 0x2D,
		0xBA, 0x36, 0x6F, 0x07, 0xC7, 0x52, 0x1F, 0x7D,
		0x4A, 0x5D, 0x4F, 0x06, 0x45, 0x8B, 0x3F, 0x4D,
	}
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
