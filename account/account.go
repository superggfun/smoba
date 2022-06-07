package account

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (m *Account) GetToken() error {
	data := fmt.Sprintf(`accessToken=%v&loginType=openSdk&openId=%v&cSystem=android&gameId=20001`, m.AccessToken, m.OpenId)
	bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/user/login", data)
	var j dejson
	err := json.Unmarshal(bodyText, &j)
	if j.ReturnCode != 0 {
		return errors.New(j.ReturnMsg)
	} else if err != nil {
		return err
	}
	m.Token = j.Token
	m.UserId = j.UserId
	m.UserName = j.UserName

	data = fmt.Sprintf("gameId=20001&token=%v&userId=%v", m.Token, m.UserId)
	bodyText = m.DoTask("https://ssl.kohsocialapp.qq.com:10001/game/chatroles", data)
	var chatroles chatroles
	err = json.Unmarshal(bodyText, &chatroles)
	if chatroles.ReturnCode != 0 {
		return errors.New(chatroles.ReturnMsg)
	} else if err != nil {
		return err
	}
	m.OriginalRoleId = chatroles.Roles[0].OriginalRoleId
	m.RoleId = chatroles.Roles[0].RoleId
	m.ServerId = chatroles.Roles[0].ServerId
	m.RoleName = chatroles.Roles[0].RoleName
	m.RoleJob = chatroles.Roles[0].RoleJob
	return nil
}
