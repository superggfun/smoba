package doTask

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type like struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

func (m *Account) AddedLikes() ([]string, error) {
	l := make([]string, 0, 10)
	u, err := m.GetNews()
	if err != nil {
		return nil, err
	}
	for _, value := range u {
		data := fmt.Sprintf("iInfoId=%v&like=1&token=%v&userId=%v", value, m.Token, m.UserId)
		bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/user/addlike", data)
		var like like
		err = json.Unmarshal(bodyText, &like)
		if err != nil {
			return nil, err
		} else if like.ReturnCode == 0 {
			l = append(l, value)
		} else {
			return l, errors.New(like.ReturnMsg)
		}
		time.Sleep(950 * time.Millisecond)
	}
	return l, nil
}
