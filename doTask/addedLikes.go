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

const (
	likeURL = "https://ssl.kohsocialapp.qq.com:10001/user/addlike"
)

// PerformLike performs the "like" action and returns any encountered error.
func (m *Account) PerformLike(value string) error {
	data := fmt.Sprintf("iInfoId=%v&like=1&token=%v&userId=%v", value, m.Token, m.UserId)
	bodyText, err := m.DoTask(likeURL, data)
	if err != nil {
		return err
	}
	var likeResp like
	if err := json.Unmarshal(bodyText, &likeResp); err != nil {
		return err
	}
	if likeResp.ReturnCode != 0 {
		return errors.New(likeResp.ReturnMsg)
	}
	return nil
}

func (m *Account) AddedLikes() ([]string, error) {
	liked := make([]string, 0, 10)
	u, err := m.GetNews()
	if err != nil {
		return nil, err
	}
	for _, value := range u {
		if err := m.PerformLike(value); err != nil {
			// Append the value to the liked list even if there's an error.
			// You may want to adjust this depending on your requirements.
			liked = append(liked, value)
			return liked, err
		}
		liked = append(liked, value)
		time.Sleep(950 * time.Millisecond)
	}
	return liked, nil
}
