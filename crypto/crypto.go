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
	"io/ioutil"

	"github.com/TheCreeper/UselessLocker/crypto/pkcs7"
	"github.com/TheCreeper/UselessLocker/crypto/pwgen"
)

// Errors
var (
	ErrCipherTextSize = errors.New("useless/crypto: ciphertext is not a multiple of the block size")
	ErrKeyBlockSize   = errors.New("useless/crypto: key size is not a multiple of the block size")
	ErrKeySize        = errors.New("useless/crypto: Key size should be either 16, 24, or 32")
)

// AES key sizes
const (
	AES128 = 16
	AES192 = 24
	AES256 = 32
)

// GenerateKey will attempt to generate a secure aes key of specified size.
// The aes key can only be 16, 24, or 32 bytes in length.
func GenerateKey(size int) ([]byte, error) {
	// Check key size
	if (size != AES128) && (size != AES192) && (size != AES256) {
		return nil, ErrKeySize
	}
	return pwgen.Generate(size)
}

// EncryptBytes will encrypt a byte slice using the provided key. Padding
// will be added using the pkcs7 padding scheme.
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

	// It's important that we always use a random iv.
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil {
		return
	}

	enc := cipher.NewCBCEncrypter(cb, iv)
	enc.CryptBlocks(ciphertext[aes.BlockSize:], b)
	return
}

// DecryptBytes will attempt to decrypt the ciphertext using the provided
// AES key. The ciphertext will be unpadded using the pkcs7 padding scheme.
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

// EncryptFile will attempt to copy the contents of the specified file into
// memory and then encrypt it using the provided key. The orginal file contents
// is over written with the encrypted bytes in memory.
func EncryptFile(key []byte, filename string) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	ciphertext, err := EncryptBytes(key, b)
	if err != nil {
		return
	}
	return ioutil.WriteFile(filename, ciphertext, 0644)
}

// DecryptFile will attempt to copy the contents of the specified file into
// memory and then decrypt it using the provided key. The orginal file contents
// is over written with the decrypted bytes in memory.
func DecryptFile(key []byte, filename string) (err error) {
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	b, err := DecryptBytes(key, ciphertext)
	if err != nil {
		return
	}

	// Copy the decrypted file contents from memory to disk.
	// Overwrite the file contents.
	return ioutil.WriteFile(filename, b, 0644)
}

// EncryptKey can be used to encrypt a AES key with the provided RSA
// public key. The ciphertext can only be decrypted using the RSA private key
// of the public key.
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

/*func DecryptKey(privBytes, ciphertext []byte) (err error) {
	privKeyBlock, _ := pem.Decode()
	return
}*/
