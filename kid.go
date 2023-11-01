package jclab_license

import (
	"crypto/ed25519"
	"github.com/jc-lab/jclab-license/internal/license_jws"
)

func MakeKid(key ed25519.PublicKey) string {
	return license_jws.MakeKid(key)
}
