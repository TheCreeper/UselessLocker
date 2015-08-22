package pwgen

import (
	"crypto/rand"
	"math/big"
)

const StdChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"

func Generate(size int) (password []byte, err error) {
	for i := 0; i < size; i++ {
		c, err := RandChar()
		if err != nil {
			return nil, err
		}
		password = append(password, c)
	}
	return
}

func RandChar() (c byte, err error) {
	max := big.NewInt(int64(len(StdChars)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return
	}
	c = StdChars[n.Int64()]
	return
}
