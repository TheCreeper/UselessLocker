package crypto

import "testing"

var (
	GeneratedKey   []byte
	SampleBytes    []byte = []byte("test")
	EncryptedBytes []byte
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey(AES256)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", key)
	GeneratedKey = key
}

func TestEncryptBytes(t *testing.T) {
	ciphertext, err := EncryptBytes(GeneratedKey, SampleBytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x\n", ciphertext)
	EncryptedBytes = ciphertext
}

func TestDecryptBytes(t *testing.T) {
	b, err := DecryptBytes(GeneratedKey, EncryptedBytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x\n", b)
}
