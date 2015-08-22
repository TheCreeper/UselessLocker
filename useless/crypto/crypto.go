package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"github.com/TheCreeper/UselessLocker/useless/crypto/pkcs7"
	"github.com/TheCreeper/UselessLocker/useless/crypto/pwgen"
)

const (
	AES128 = 16
	AES192 = 24
	AES256 = 32
)

var (
	ErrKeySize        = errors.New("useless/crypto: Key size should be either 16, 24, or 32")
	ErrKeyBlockSize   = errors.New("useless/crypto: key size is not a multiple of the block size")
	ErrCipherTextSize = errors.New("useless/crypto: ciphertext is not a multiple of the block size")
)

func GenerateKey(size int) ([]byte, error) {
	// Check key size
	if (size != AES128) && (size != AES192) && (size != AES256) {
		return nil, ErrKeySize
	}
	return pwgen.Generate(size)
}

func EncryptBytes(key, b []byte) (ciphertext []byte, err error) {
	if len(key)%aes.BlockSize != 0 {
		return nil, ErrKeyBlockSize
	}

	cb, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	b, err = pkcs7.Pad(b, aes.BlockSize)
	if err != nil {
		return
	}

	ciphertext = make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(b); err != nil {
		return
	}

	enc := cipher.NewCBCEncrypter(cb, iv)
	enc.CryptBlocks(ciphertext[aes.BlockSize:], b)
	return
}

func DecryptBytes(key, ciphertext []byte) (b []byte, err error) {
	if len(key)%aes.BlockSize != 0 {
		return nil, ErrKeyBlockSize
	}

	if len(b)%aes.BlockSize != 0 {
		return nil, ErrCipherTextSize
	}

	cb, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	ciphertext, err = pkcs7.UnPad(ciphertext, aes.BlockSize)
	if err != nil {
		return
	}

	dec := cipher.NewCBCDecrypter(cb, ciphertext[:aes.BlockSize])
	dec.CryptBlocks(b, ciphertext[aes.BlockSize:])
	return
}
