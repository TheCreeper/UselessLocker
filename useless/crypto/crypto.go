package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/TheCreeper/UselessLocker/useless/crypto/pkcs7"
	"github.com/TheCreeper/UselessLocker/useless/crypto/pwgen"
)

// AES key sizes
const (
	AES128 = 16
	AES192 = 24
	AES256 = 32
)

// Errors
var (
	ErrKeySize        = errors.New("useless/crypto: Key size should be either 16, 24, or 32")
	ErrKeyBlockSize   = errors.New("useless/crypto: key size is not a multiple of the block size")
	ErrCipherTextSize = errors.New("useless/crypto: ciphertext is not a multiple of the block size")
)

// GenerateKey will attempt to generate a secure aes key of specified size. The aes key can only be
// 16, 24, or 32 bytes in length.
func GenerateKey(size int) ([]byte, error) {
	// Check key size
	if (size != AES128) && (size != AES192) && (size != AES256) {
		return nil, ErrKeySize
	}
	return pwgen.Generate(size)
}

// EncryptBytes will encrypted a byte slice using the provided key. Padding will be added using
// the pkcs7 padding scheme.
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
	if _, err = rand.Read(iv); err != nil {
		return
	}

	enc := cipher.NewCBCEncrypter(cb, iv)
	enc.CryptBlocks(ciphertext[aes.BlockSize:], b)
	return
}

// DecryptBytes will attempt to decrypt the ciphertext using the provided AES key. The ciphertext
// will be unpadded using the pkcs7 padding scheme.
// The AES key and ciphertext should be a multiple of the aes block size (16).
func DecryptBytes(key, ciphertext []byte) (b []byte, err error) {
	if len(key)%aes.BlockSize != 0 {
		return nil, ErrKeyBlockSize
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, ErrCipherTextSize
	}

	cb, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Get size of ciphertext not including the iv.
	b = make([]byte, len(ciphertext[aes.BlockSize:]))

	dec := cipher.NewCBCDecrypter(cb, ciphertext[:aes.BlockSize])
	dec.CryptBlocks(b, ciphertext[aes.BlockSize:])

	b, err = pkcs7.UnPad(b, aes.BlockSize)
	if err != nil {
		return
	}
	return
}

// EncryptKey can be used to encrypt a AES key with the provided RSA public key. The ciphertext
// can only be decrypted using the RSA private key of the public key.
func EncryptKey(pubBytes, key []byte) (ciphertext []byte, err error) {
	block, _ := pem.Decode(pubBytes)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	// Convert pkix to rsa public key
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, key, nil)
}
