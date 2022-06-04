package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func (m *Account) GetToken() {
	data := fmt.Sprintf(`accessToken=%v&loginType=openSdk&openId=%v&cSystem=android&gameId=20001`, m.AccessToken, m.OpenId)
	bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/user/login", data)
	var j dejson
	err := json.Unmarshal(bodyText, &j)
	if err != nil {
		log.Panic(err)
	} else if j.Result != 0 {
		log.Panic(errors.New("login error"))
	}
	m.Token = j.Token
	m.UserId = j.UserId
}
