package doTask

import (
	"encoding/json"
)

const NewsURL = "https://kohcamp.qq.com/info/listinfov2"

type NewsResponse struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	Data       struct {
		List []NewsItem `json:"list"`
	} `json:"data"`
}

type NewsItem struct {
	InfoContent struct {
		InfoId string `json:"infoId"`
		User   struct {
			User struct {
				User struct {
					UserId string `json:"userId"`
				} `json:"user"`
			} `json:"user"`
		} `json:"user"`
	} `json:"infoContent"`
}

func (m *Account) GetNews() (map[string]string, error) {
	u := make(map[string]string)
	data := `{"page": 0, "channelId": 25818}`

	bodyText, err := m.DoGift(NewsURL, data, m.UserId, m.OriginalRoleId)
	if err != nil {
		return nil, err
	}

	var newsResp NewsResponse
	if err = json.Unmarshal(bodyText, &newsResp); err != nil {
		return nil, err
	}

	for _, item := range newsResp.Data.List {
		infoID := item.InfoContent.InfoId
		userID := item.InfoContent.User.User.User.UserId
		if infoID != "" {
			u[userID] = infoID
		}
	}
	return u, nil
}
