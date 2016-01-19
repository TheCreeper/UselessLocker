package unique

import (
	"crypto/sha256"
	"io/ioutil"
	"os"
)

func MachineID() (id []byte, err error) {
	// Try using the systemd generated machine id
	id, err = ioutil.ReadFile("/etc/machine-id")
	if err == nil {
		return
	}

	// Try using the dbus generated machine id
	id, err = ioutil.ReadFile("/var/lib/dbus/machine-id")
	if err == nil {
		return
	}

	// If all else fails then simply hash the hostname
	name, err := os.Hostname()
	if err != nil {
		return
	}
	return
}
