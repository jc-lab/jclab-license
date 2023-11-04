package jclab_license

import (
	"github.com/jc-lab/jclab-license/license_model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testLicenseCode = "eyJhbGciOiJFZERTQSIsImtpZCI6Ilo5LVJ1d0RaVVBvIn0.eJwkzM1qg0AUxfFXKWetUC0Kzq6ItLTSjSUEQwhXM8FJ5guvSTQh7x6My3Pg_7tD0QARpVn2_hFncRJAMUPgJw9LahDA927PEBswGa8ltgG0aofJSwg4Ly27c99KvO6TnDoIeE_HtO7GW2d0mI5WV-to-qt2ub1-fbOjom4ORfK78p-8dIbGC0QYzYMtmdn-lzy8laqVlhedDSkNATyeAQAA__-CfzrK.aJrYSNA0VJsX9ONwCwhDwjHrt-BTdDdzKjg6wvTQ9-VavvjPjIGmjl3vMqeQJAVCjLe4YQTZuU6-VRfvJW5CCQ"
const testLicenseKey = "4c6ec158-7410-402f-bae7-904125c0901f"

func TestVerify_success(t *testing.T) {
	claims, err := Verify("sample", 1, testLicenseCode)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "JC-Lab", claims.Iss)
	assert.Equal(t, []string{"sample"}, claims.Products)
	assert.Equal(t, license_model.LicenseTypeOpenSource, claims.LicenseType)
	assert.Equal(t, LicenseKeyHash(testLicenseKey), claims.LicenseKeyHash)
	assert.Equal(t, -1, claims.LicenseMaxVersion)
	assert.Equal(t, "Test License", claims.LicenseeName)
	assert.Equal(t, "", claims.LicenseeMail)
}
