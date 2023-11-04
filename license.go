// Copyright (C) 2023 JC-Lab. All rights reserved.

package jclab_license

import (
	"crypto"
	"encoding/base64"
	"errors"
	"github.com/jc-lab/jclab-license/internal/license_jws"
	"github.com/jc-lab/jclab-license/license_model"
)

var ErrLicense = errors.New("invalid license")

func Verify(product string, version int, input string) (*license_model.Claims, error) {
	jws, err := license_jws.Verify[license_model.Claims](input)
	if err != nil {
		return nil, err
	}
	var productMatched bool = false
	for _, s := range jws.Claims.Products {
		if s == product {
			productMatched = true
		}
	}

	if !productMatched || ((jws.Claims.LicenseMaxVersion > 0) && jws.Claims.LicenseMaxVersion < version) {
		return nil, ErrLicense
	}
	return &jws.Claims, nil
}

func LicenseKeyHash(licenseKey string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(licenseKey))
	return base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
}
