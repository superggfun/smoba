package doTask

import (
	"encoding/json"
)

type gNews struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	data       `json:"data"`
}

type data struct {
	List []list `json:"list"`
}

type list struct {
	infoContent `json:"infoContent"`
}

type infoContent struct {
	InfoId string `json:"infoId"`
	user1  `json:"user"`
}

type user1 struct {
	user2 `json:"user"`
}

type user2 struct {
	user3 `json:"user"`
}

type user3 struct {
	UserId string `json:"userId"`
}

func (m *Account) GetNews() (map[string]string, error) {
	u := make(map[string]string)
	data := `{"page": 0, "channelId": 25818}`
	bodyText := m.DoGift("https://kohcamp.qq.com/info/listinfov2", data, m.UserId)
	var gNews gNews
	err := json.Unmarshal(bodyText, &gNews)
	if err != nil {
		return nil, err
	} else {
		for _, v := range gNews.List {
			if v.InfoId != "" {
				u[v.UserId] = v.InfoId
			}
		}
	}
	return u, nil
}
