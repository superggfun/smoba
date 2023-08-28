package doGift

import (
	"encoding/json"
	"errors"
	"fmt"
)

type sign struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	signData   `json:"data"`
}

type signData struct {
	Good  string `json:"good"`
	Bad   string `json:"bad"`
	Lunar string `json:"lunar"`
}

func (m *Account) Sign() (signData, error) {
	data := fmt.Sprintf(`{"cSystem":"android","h5Get":1,"roleId":"%v"}`, m.RoleId)
	bodyText, err := m.DoGift("https://kohcamp.qq.com/operation/action/signin", data, m.UserId, m.OriginalRoleId)
	var sign sign
	if err != nil {
		return sign.signData, err
	}
	err = json.Unmarshal(bodyText, &sign)
	if err != nil {
		return sign.signData, err
	} else if sign.ReturnCode != 0 {
		return sign.signData, errors.New(sign.ReturnMsg)
	} else {
		return sign.signData, nil
	}
}
