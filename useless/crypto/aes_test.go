package crypto

import "testing"

var (
	GeneratedKey   []byte
	SampleBytes    []byte = []byte("test")
	EncryptedBytes []byte
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateAESKey(AES256)
	if err != nil {
		t.Fatal(err)
	}
	GeneratedKey = key
}

func TestEncryptBytes(t *testing.T) {
	a, err := NewAES(GeneratedKey)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext, err := a.EncryptBytes(SampleBytes)
	if err != nil {
		t.Fatal(err)
	}
	EncryptedBytes = ciphertext

	t.Logf("%x\n", ciphertext)
}

func TestDecryptBytes(t *testing.T) {
	a, err := NewAES(GeneratedKey)
	if err != nil {
		t.Fatal(err)
	}

	b, err := a.DecryptBytes(EncryptedBytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x\n", b)
}
