package config

type Account struct {
	AccessToken     string `json:"accessToken"`
	OpenId          string `json:"openId"`
	Sig             string `json:"sig"`
	Timestamp       string `json:"timestamp"`
	MsdkEncodeParam string `json:"msdkEncodeParam"`
	MsdkToken       string `json:"msdkToken"`
}

type Wxpush struct {
	Pushplus string `json:"pushplus"`
}

type Config struct {
	Account []Account
	Wxpush  `json:"wxpush"`
}
