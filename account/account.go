package account

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (m *Account) sendRequestAndDecode(url string, data string, v interface{}) error {
	bodyText, err := m.DoTask(url, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyText, v)
	if err != nil {
		return err
	}

	switch j := v.(type) {
	case *dejson:
		if j.ReturnCode != 0 {
			return errors.New(j.ReturnMsg)
		}
	case *chatroles:
		if j.ReturnCode != 0 {
			return errors.New(j.ReturnMsg)
		}
	}

	return nil
}

func (m *Account) GetToken() error {
	data := fmt.Sprintf(`accessToken=%v&loginType=openSdk&openId=%v&cSystem=android&gameId=20001`, m.AccessToken, m.OpenId)
	var j dejson
	err := m.sendRequestAndDecode("https://ssl.kohsocialapp.qq.com:10001/user/login", data, &j)
	if err != nil {
		return err
	}

	m.Token = j.Token
	m.UserId = j.UserId
	m.UserName = j.UserName

	data = fmt.Sprintf("gameId=20001&token=%v&userId=%v", m.Token, m.UserId)
	var chatroles chatroles
	err = m.sendRequestAndDecode("https://ssl.kohsocialapp.qq.com:10001/game/chatroles", data, &chatroles)
	if err != nil {
		return err
	}

	m.OriginalRoleId = chatroles.Roles[0].OriginalRoleId
	m.RoleId = chatroles.Roles[0].RoleId
	m.ServerId = chatroles.Roles[0].ServerId
	m.RoleName = chatroles.Roles[0].RoleName
	m.RoleJob = chatroles.Roles[0].RoleJob
	return nil
}
