// Copyright (C) 2023 JC-Lab. All rights reserved.

package license_jws

import (
	"bytes"
	"compress/zlib"
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"io"
)

var ErrInvalid = errors.New("invalid format")

func Verify[D any](input string) (*JWS[D], error) {
	raw := []byte(input)

	n := bytes.Index(raw, dotByte)
	if n < 0 {
		return nil, ErrInvalid
	}
	headerB64 := raw[0:n]
	m := bytes.Index(raw[n+1:], dotByte)
	if m < 0 {
		return nil, ErrInvalid
	}
	compressedBodyB64 := raw[n+1 : n+1+m]
	signedPart := raw[:n+1+m]
	signatureB64 := raw[n+1+m+1:]

	header, err := encoding.DecodeString(string(headerB64))
	if m < 0 {
		return nil, ErrInvalid
	}
	compressedBody, err := encoding.DecodeString(string(compressedBodyB64))
	if m < 0 {
		return nil, ErrInvalid
	}
	signature, err := encoding.DecodeString(string(signatureB64))
	if m < 0 {
		return nil, ErrInvalid
	}

	jws := &JWS[D]{}
	if err = json.Unmarshal(header, &jws.Header); err != nil {
		return nil, ErrInvalid
	}

	jwk, ok := publicKeys[jws.Header.Kid]
	if !ok {
		return nil, ErrInvalid
	}

	if !ed25519.Verify(jwk.key, signedPart, signature) {
		return nil, ErrInvalid
	}

	r, err := zlib.NewReader(bytes.NewReader(compressedBody))
	if err != nil {
		return nil, ErrInvalid
	}
	bodyRaw, err := io.ReadAll(r)
	if err != nil {
		return nil, ErrInvalid
	}

	if err = json.Unmarshal(bodyRaw, &jws.Claims); err != nil {
		return nil, ErrInvalid
	}

	return jws, nil
}
