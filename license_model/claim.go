package license_model

type Claims struct {
	Iat               int64  `json:"iat"`
	Iss               string `json:"iss"`
	Product           string `json:"product"`
	LicenseKeyHash    string `json:"license_key_hash"`
	LicenseMaxVersion int    `json:"license_max_version"`
	LicenseeName      string `json:"licensee_name"`
	LicenseeEmail     string `json:"licensee_email"`
}
