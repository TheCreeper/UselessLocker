package browser

import "os/exec"

func openURL(u string) error {
	return exec.Command("cmd", "/c", "start", u).Start()
}
