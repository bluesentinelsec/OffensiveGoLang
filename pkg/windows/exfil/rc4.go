package exfil

import (
	"crypto/rc4"
)

// CryptRC4 uses a key to RC4 encrypt data
func CryptRC4(key, data []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	c.XORKeyStream(dst, data)
	return dst, err
}
