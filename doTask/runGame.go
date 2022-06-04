package doTask

import (
	"encoding/json"
	"errors"
	"fmt"
)

type run struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

// 运行游戏
func (m *Account) RunGame() error {
	data := fmt.Sprintf("type=2&token=%v&userId=%v", m.Token, m.UserId)
	bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata", data)
	var run run
	err := json.Unmarshal(bodyText, &run)
	if err != nil {
		return err
	} else if run.ReturnCode != 0 {
		return errors.New(run.ReturnMsg)
	}
	return nil
}
