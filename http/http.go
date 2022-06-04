package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Do interface {
	DoTask(url string, data string) []byte
	DoGift() []byte
}

func (m *Account) DoTask(url string, data string) []byte {
	client := &http.Client{}
	d := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText

}

func (m *Account) DoGift(url string, data string, userId string) []byte {
	client := &http.Client{}
	d := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("timestamp", m.Timestamp)
	req.Header.Add("userId", "460014701")
	req.Header.Add("algorithm", "v2")
	req.Header.Add("openid", m.OpenId)
	req.Header.Add("encode", "2")
	req.Header.Add("roleId", m.RoleId)
	req.Header.Add("source", "smoba_zhushou")
	req.Header.Add("msdkToken", m.MsdkToken)
	req.Header.Add("gameOpenid", "8C7A5CE953A1E8A10B1CD2029A8DCFEB")
	req.Header.Add("msdkEncodeParam", m.MsdkEncodeParam)
	req.Header.Add("gameId", "20001")
	req.Header.Add("sig", m.Sig)
	req.Header.Add("appid", "1105200115")
	req.Header.Add("version", "3.1.96a")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText

}
