// Copyright (C) 2023 JC-Lab. All rights reserved.

package license_model

//go:generate msgp

type LicenseType string

const (
	LicenseTypeOpenSource LicenseType = "opensource"
	LicenseTypeCommercial LicenseType = "commercial"
)

type Claims struct {
	Iat               int64       `json:"iat" msg:"iat"`
	Iss               string      `json:"iss" msg:"iss"`
	Product           string      `json:"product" msg:"product"`
	LicenseType       LicenseType `json:"lictype" msg:"lictype"`
	LicenseKeyHash    string      `json:"lickeyh" msg:"lickeyh"`
	LicenseMaxVersion int         `json:"licmaxv" msg:"licmaxv"`
	LicenseeName      string      `json:"lisname" msg:"lisname"`
	LicenseeEmail     string      `json:"lisemail" msg:"lisemail"`
}
