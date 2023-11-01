package license_model

//go:generate msgp

type Claims struct {
	Iat               int64  `json:"iat" msg:"iat"`
	Iss               string `json:"iss" msg:"iss"`
	Product           string `json:"product" msg:"product"`
	LicenseKeyHash    string `json:"license_key_hash" msg:"license_key_hash"`
	LicenseMaxVersion int    `json:"license_max_version" msg:"license_max_version"`
	LicenseeName      string `json:"licensee_name" msg:"licensee_name"`
	LicenseeEmail     string `json:"licensee_email" msg:"licensee_email"`
}
