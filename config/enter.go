package config

type Account struct {
	AccessToken     string `json:"accessToken"`
	OpenId          string `json:"openId"`
	RoleId          string `json:"roleId"`
	Sig             string `json:"sig"`
	Timestamp       string `json:"timestamp"`
	MsdkEncodeParam string `json:"msdkEncodeParam"`
	MsdkToken       string `json:"msdkToken"`
}
