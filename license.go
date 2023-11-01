// Copyright (C) 2023 JC-Lab. All rights reserved.

package jclab_license

import (
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
	if jws.Claims.Product != product || ((jws.Claims.LicenseMaxVersion > 0) && jws.Claims.LicenseMaxVersion < version) {
		return nil, ErrLicense
	}
	return &jws.Claims, nil
}
