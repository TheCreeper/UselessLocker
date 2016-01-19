package store

import "testing"

func TestOpen(t *testing.T) {
	fs, err := OpenFile("testdata/dumb/dumb")
	if err != nil {
		t.Fatal(err)
	}
	//	defer s.Close()

	b, err := fs.ReadFile("/helloworld.txt")
	if err != nil {
		t.Fatal(err)
	}
	println(string(b))
}
