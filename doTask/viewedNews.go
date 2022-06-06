package doTask

import (
	"encoding/json"
	"fmt"
)

type view struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

func (m *Account) ViewedNews() ([]string, error) {
	l := make([]string, 0, 10)
	u, err := m.GetNews()
	if err != nil {
		return nil, err
	}
	for _, value := range u {
		data := fmt.Sprintf("iInfoId=%v&gameId=20001&token=%v&userId=%v", value, m.Token, m.UserId)
		bodyText := m.DoTask("https://ssl.kohsocialapp.qq.com:10001/game/detailinfov3", data)
		var view view
		err = json.Unmarshal(bodyText, &view)
		if err != nil {
			return nil, err
		} else if view.ReturnCode == 0 {
			l = append(l, value)
		}
	}
	return l, nil
}
