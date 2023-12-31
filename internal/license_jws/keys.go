// Copyright (C) 2023 JC-Lab. All rights reserved.

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
		0xda, 0xc6, 0x67, 0xc3, 0x9d, 0x88, 0x73, 0xd1,
		0x87, 0x77, 0xab, 0xdc, 0x34, 0x5f, 0x3b, 0xbb,
		0x9b, 0x7e, 0x44, 0x50, 0x73, 0x59, 0x22, 0xea,
		0x2b, 0x1d, 0x2f, 0x28, 0x48, 0xb4, 0x3a, 0xc1,
	}, -1)
}
