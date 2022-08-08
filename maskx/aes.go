package maskx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/panyaxbo/libs/logx"
	"github.com/pkg/errors"
)

type AESer interface {
	Encrypt(raw string) string
	Decrypt(cipherText string) (string, error)
}

type AES struct {
	aesgcm cipher.AEAD
	nonce  []byte
}

func NewAES(key, nonce []byte) *AES {
	block, err := aes.NewCipher(key)
	if err != nil {
		logx.Panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		logx.Panic(err)
	}

	return &AES{
		aesgcm: aesgcm,
		nonce:  nonce,
	}
}

func (a *AES) Encrypt(raw string) string {
	cipherText := a.aesgcm.Seal(nil, a.nonce, []byte(raw), nil)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func (a *AES) Decrypt(cipherText string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", errors.WithStack(err)
	}

	plainText, err := a.aesgcm.Open(nil, a.nonce, encrypted, nil)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(plainText), nil
}
