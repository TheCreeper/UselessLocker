package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"github.com/TheCreeper/UselessLocker/useless/crypto/pkcs7"
)

const (
	AES128 = 16
	AES192 = 24
	AES256 = 32
)

func GenerateAESKey(size int) ([]byte, error) {
	// Check key size
	if (size != AES128) && (size != AES192) && (size != AES256) {
		return nil, errors.New("useless/crypto: Key size should be either 16, 24, or 32")
	}
	return GeneratePassword(size)
}

type AES struct{ cipher.Block }

func NewAES(key []byte) (AES, error) {
	if len(key)%aes.BlockSize != 0 {
		return AES{}, errors.New("useless/crypto: key is not a multiple of the block size")
	}

	cb, err := aes.NewCipher(key)
	if err != nil {
		return AES{}, err
	}
	return AES{cb}, nil
}

func (a AES) EncryptBytes(b []byte) (ciphertext []byte, err error) {
	b, err = pkcs7.Pad(b, aes.BlockSize)
	if err != nil {
		return
	}

	ciphertext = make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(b); err != nil {
		return
	}

	enc := cipher.NewCBCEncrypter(a, iv)
	enc.CryptBlocks(ciphertext[aes.BlockSize:], b)
	return
}

func (a AES) DecryptBytes(ciphertext []byte) (b []byte, err error) {
	if len(b)%aes.BlockSize != 0 {
		return nil, errors.New("useless/crypto: encrypted bytes is not a multiple of the block size")
	}

	ciphertext, err = pkcs7.UnPad(ciphertext, aes.BlockSize)
	if err != nil {
		return
	}

	dec := cipher.NewCBCDecrypter(a, ciphertext[:aes.BlockSize])
	dec.CryptBlocks(b, ciphertext[aes.BlockSize:])
	return
}
