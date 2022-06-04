package doTask

import (
	"encoding/json"
	"errors"
	"fmt"
)

type share struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

// 分享游戏
func (m *Account) ShareGame() error {
	data := fmt.Sprintf("type=1&token=%v&userId=%v", m.Token, m.UserId)
	bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata", data)
	var share share
	err := json.Unmarshal(bodyText, &share)
	if err != nil {
		return err
	} else if share.ReturnCode != 0 {
		return errors.New(share.ReturnMsg)
	}
	return nil
}
