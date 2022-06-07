package wxpush

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type data struct {
	Token    string `json:"token"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Template string `json:"template"`
}

type push struct {
	Code int    `json:"code"`
	Msg  string `json:"Msg"`
}

func pushplus(token string, str string) error {
	url := "http://www.pushplus.plus/send"
	data := data{token, "王者营地签到", str, "markdown"}
	s, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	d := strings.NewReader(string(s))
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var push push
	err = json.Unmarshal(bodyText, &push)
	if err != nil {
		fmt.Println(err)
	}
	if push.Code == 200 {
		return nil
	} else {
		return errors.New(push.Msg)
	}
}
