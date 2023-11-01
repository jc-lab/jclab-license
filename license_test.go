package jclab_license

import (
	"crypto"
	"encoding/base64"
	"github.com/jc-lab/jclab-license/license_model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testLicenseCode = "eyJhbGciOiJFZERTQSIsImtpZCI6Ilo5LVJ1d0RaVVBvIn0.eJwszE9LhEAcxvG3Es9ZD0YONrcQKUq6GBFe4uc04dT8w9-4q7vse1_EPT4PfD9nGEqQhXisKvFwX1YZDDMkXuu8pQEZ4hR-ZpUgweSi1chgjUpr1JAIUXsO86Ru979eR0jESH-iH5fT6GwuFm-7r2J9775rf3x-4UBNP_w25dtnfOK9c7QcIPNiG-zJbfaH5nTXGqU97zprR8ZCApdrAAAA___ZJjtb.h7wAValQU7rvdgCPaVWysgjx18g5LY2A1hnKQCAIYPLDQrmzDVlsZdLqWjOGIRhl3RsPdEGRcKNvV8_6lK9nDw"
const testLicenseKey = "4c6ec158-7410-402f-bae7-904125c0901f"

func TestVerify_success(t *testing.T) {
	claims, err := Verify("sample", 1, testLicenseCode)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "JC-Lab", claims.Iss)
	assert.Equal(t, "sample", claims.Product)
	assert.Equal(t, license_model.LicenseTypeOpenSource, claims.LicenseType)
	assert.Equal(t, licenseKeyHash(testLicenseKey), claims.LicenseKeyHash)
	assert.Equal(t, -1, claims.LicenseMaxVersion)
	assert.Equal(t, "Test License", claims.LicenseeName)
	assert.Equal(t, "", claims.LicenseeEmail)
}

func licenseKeyHash(licenseKey string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(licenseKey))
	return base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
}
