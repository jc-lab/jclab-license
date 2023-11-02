package license_jws

import "crypto/ed25519"

type publicKey struct {
	kid      string
	key      ed25519.PublicKey
	notAfter int64
}

var publicKeys map[string]*publicKey

func init() {
	publicKeys = map[string]*publicKey{}
	addPublicKey := func(raw []byte, notAfter int64) {
		item := &publicKey{
			kid:      MakeKid(raw),
			notAfter: notAfter,
			key:      raw,
		}
		publicKeys[item.kid] = item
	}
	addPublicKey([]byte{
		0xd4, 0x99, 0xef, 0x98, 0x1e, 0xff, 0x90, 0x8a,
		0xca, 0xbe, 0x2a, 0xd7, 0x14, 0xd7, 0x23, 0xfc,
		0xba, 0x21, 0x5f, 0x2c, 0x14, 0x32, 0xa1, 0x04,
		0xd7, 0xcb, 0xe6, 0x7d, 0xa6, 0x09, 0xc8, 0x26,
	}, -1)
}
