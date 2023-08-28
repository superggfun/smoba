package doTask

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	addFriendURL = "https://ssl.kohsocialapp.qq.com:10001/user/addfriend"
	delFriendURL = "https://ssl.kohsocialapp.qq.com:10001/user/delfriend"
)

type friend struct {
	Result     int    `json:"result"`
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

func (m *Account) AddFriend(friendUserId string) (bool, error) {
	data := fmt.Sprintf("friendUserId=%v&token=%v&userId=%v", friendUserId, m.Token, m.UserId)
	bodyText, err := m.DoTask(addFriendURL, data)
	if err != nil {
		return false, err
	}

	var friendResponse friend
	if err = json.Unmarshal(bodyText, &friendResponse); err != nil {
		return false, err
	}

	switch friendResponse.ReturnCode {
	case 0:
		return true, nil
	case -30029:
		return false, nil
	default:
		return false, errors.New(friendResponse.ReturnMsg)
	}
}

func (m *Account) DelFriend(friendUserId string) (bool, error) {
	data := fmt.Sprintf("friendUserId=%v&token=%v&userId=%v", friendUserId, m.Token, m.UserId)
	bodyText, err := m.DoTask(delFriendURL, data)
	if err != nil {
		return false, err
	}

	var friendResponse friend
	if err = json.Unmarshal(bodyText, &friendResponse); err != nil {
		return false, err
	}

	switch friendResponse.ReturnCode {
	case 0:
		return true, nil
	case -30029:
		return false, nil
	default:
		return false, errors.New(friendResponse.ReturnMsg)
	}
}
