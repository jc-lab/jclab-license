package license_jws

import (
	"encoding/base64"
)

type Header struct {
	Alg string `json:"alg"`
}

type JWS[D any] struct {
	Header Header
	Claims D
}

var dotByte = []byte{'.'}
var encoding = base64.RawURLEncoding
