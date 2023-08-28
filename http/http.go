package http

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func (m *Account) sendRequest(url string, data string, headers map[string]string) ([]byte, error) {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyText, nil
}

func (m *Account) DoTask(url string, data string) ([]byte, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"NOENCRYPT":    "1",
	}

	return m.sendRequest(url, data, headers)
}

func (m *Account) DoGift(url string, data string, userId string, roleId string) ([]byte, error) {
	headers := map[string]string{
		"Content-Type":    "application/json",
		"timestamp":       m.Timestamp,
		"userId":          userId,
		"algorithm":       "v2",
		"openid":          m.OpenId,
		"encode":          "2",
		"roleId":          roleId,
		"source":          "smoba_zhushou",
		"msdkToken":       m.MsdkToken,
		"msdkEncodeParam": m.MsdkEncodeParam,
		"gameId":          "20001",
		"sig":             m.Sig,
		"appid":           "1105200115",
		"version":         "3.1.96a",
		"NOENCRYPT":       "1",
	}

	return m.sendRequest(url, data, headers)
}
