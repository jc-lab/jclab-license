package license_jws

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
)

func MakeKid(publicKey ed25519.PublicKey) string {
	h := sha256.New()
	h.Write(publicKey[0:32])
	digest := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(digest[0:8])
}
