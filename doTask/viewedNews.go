package doTask

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseURL = "https://ssl.kohsocialapp.qq.com:10001"
)

type TaskResponse struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

func (m *Account) performTaskWithData(taskData, endpoint string) ([]string, error) {
	l := make([]string, 0, 10)

	responseBody, err := m.DoTask(BaseURL+endpoint, taskData)

	if err != nil {
		return nil, err
	}

	var response TaskResponse
	err = json.Unmarshal(responseBody, &response)

	if err != nil {
		return nil, err
	} else if response.ReturnCode != http.StatusOK {
		return nil, fmt.Errorf(response.ReturnMsg)
	}

	l = append(l, taskData)

	return l, nil
}

func (m *Account) ViewedNews() ([]string, error) {
	u, err := m.GetNews()
	if err != nil {
		return nil, err
	}
	var results []string
	for _, value := range u {
		data := fmt.Sprintf("iInfoId=%v&gameId=20001&token=%v&userId=%v", value, m.Token, m.UserId)
		result, err := m.performTaskWithData(data, "/game/detailinfov3")
		if err != nil {
			return nil, err
		}
		results = append(results, result...)
	}
	return results, nil
}
