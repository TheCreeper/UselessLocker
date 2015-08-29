package store

import "testing"

func TestOpen(t *testing.T) {
	s, err := OpenFile("testdata/dumb/dumb")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	file, err := s.Load("helloworld.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(file))
}
