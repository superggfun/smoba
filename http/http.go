package http

import (
	"crypto/tls"
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
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	d := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("NOENCRYPT", "1")
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

func (m *Account) DoGift(url string, data string, userId string, roleId string) []byte {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	d := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("timestamp", m.Timestamp)
	req.Header.Add("userId", userId)
	req.Header.Add("algorithm", "v2")
	req.Header.Add("openid", m.OpenId)
	req.Header.Add("encode", "2")
	req.Header.Add("roleId", roleId)
	req.Header.Add("source", "smoba_zhushou")
	req.Header.Add("msdkToken", m.MsdkToken)
	req.Header.Add("msdkEncodeParam", m.MsdkEncodeParam)
	req.Header.Add("gameId", "20001")
	req.Header.Add("sig", m.Sig)
	req.Header.Add("appid", "1105200115")
	req.Header.Add("version", "3.1.96a")
	req.Header.Add("NOENCRYPT", "1")
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
