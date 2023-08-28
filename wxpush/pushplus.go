package wxpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type data struct {
	Token    string `json:"token"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Template string `json:"template"`
}

type pushResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var httpClient = &http.Client{}

func pushplus(token string, str string) error {
	url := "http://www.pushplus.plus/send"
	data := data{Token: token, Title: "王者营地签到", Content: str, Template: "markdown"}

	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var push pushResponse
	err = json.Unmarshal(bodyText, &push)
	if err != nil {
		return err
	}

	if push.Code != 200 {
		return errors.New(push.Msg)
	}

	return nil
}
