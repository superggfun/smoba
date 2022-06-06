package doTask

import (
	"encoding/json"
	"errors"
	"fmt"
)

type friend struct {
	Result     int    `json:"result"`
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

func (m *Account) AddFriend(friendUserId string) (bool, error) {
	data := fmt.Sprintf("friendUserId=%v&token=%v&userId=%v", friendUserId, m.Token, m.UserId)
	bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/user/addfriend", data)
	var friend friend
	err := json.Unmarshal(bodyText, &friend)
	if err != nil {
		return false, err
	} else if friend.ReturnCode == -30029 {
		return false, nil
	} else if friend.ReturnCode == 0 {
		return true, nil
	} else {
		return false, errors.New(friend.ReturnMsg)
	}
}

func (m *Account) DelFriend(friendUserId string) (bool, error) {
	data := fmt.Sprintf("friendUserId=%v&token=%v&userId=%v", friendUserId, m.Token, m.UserId)
	bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/user/delfriend", data)
	var friend friend
	err := json.Unmarshal(bodyText, &friend)
	if err != nil {
		return false, err
	} else if friend.ReturnCode == -30029 {
		return false, nil
	} else if friend.ReturnCode == 0 {
		return true, nil
	} else {
		return false, errors.New(friend.ReturnMsg)
	}
}
