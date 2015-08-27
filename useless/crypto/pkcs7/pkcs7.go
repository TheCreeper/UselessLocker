// PKCS7 implements PKCS#7 padding, described at http://tools.ietf.org/html/rfc5652#section-6.3.
package pkcs7

import (
	"bytes"
	"errors"
)

var (
	ErrBlockSize     = errors.New("useless/crypto/pkcs7: block size is out of bounds")
	ErrByteBlockSize = errors.New("useless/crypto/pkcs7: byte length is not a multiple of the block size")
)

// Pad adds padding, with each padded byte being the total number of bytes added. Byte slice
// should always be padded even if its a multiple of the block size. This makes it easier
// to unpad bytes if we can safely assume they are always padded.
//
// Example for a blocksize of 8:
// -> [DD DD DD DD 04 04 04 04]
func Pad(b []byte, size int) (out []byte, err error) {
	if size < 1 || size >= 256 {
		return nil, ErrBlockSize
	}

	padding := size - len(b)%size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(b, padtext...), nil
}

// UnPad removes padding.
func UnPad(b []byte, size int) (out []byte, err error) {
	if size < 1 || size >= 256 {
		return nil, ErrBlockSize
	}

	if len(b)%size != 0 {
		return nil, ErrByteBlockSize
	}

	// Get total number of bytes
	blen := len(b)

	// Get the total number of bytes added to the slice
	padlen := int(b[blen-1])

	// Remove padding by only returning the non-padded bytes at the beginning of the slice.
	return b[:(blen - padlen)], nil
}
