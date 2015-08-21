package useless

import (
	"os/user"
	"testing"
)

func TestGetFileList(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		return
	}
	t.Log(user.HomeDir)

	files, err := GetFileList(user.HomeDir, 10485760)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(files)
}
