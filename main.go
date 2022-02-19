package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tidwall/gjson"
)

type Account struct {
	AccessToken     string `json:"accessToken"`
	OpenId          string `json:"openId"`
	RoleId          string `json:"roleId"`
	Sig             string `json:"sig"`
	Timestamp       string `json:"timestamp"`
	MsdkEncodeParam string `json:"msdkEncodeParam"`
	MsdkToken       string `json:"msdkToken"`
}

type Token struct {
	UserId string
	Token  string
}

type DefineEvent struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func readFile() *[]Account {
	b, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Print(err)
	}

	account := []Account{}
	err = json.Unmarshal(b, &account)
	if err != nil {
		fmt.Print(err)
	}
	return &account
}

// Post请求
func postHttp(url string, data string) string {
	client := &http.Client{}
	d := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "131")
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
	return string(bodyText)

}

// 获取游戏账号信息
func (m *Account) getToken() *Token {
	data := fmt.Sprintf(`accessToken=%v&loginType=openSdk&openId=%v&cSystem=android&gameId=20001`, m.AccessToken, m.OpenId)
	bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/user/login", data)
	a := Token{
		UserId: gjson.Get(bodyText, "data.userId").String(),
		Token:  gjson.Get(bodyText, "data.token").String(),
	}
	return &a
}

// 获取资讯
func mainPage() *[]string {
	client := &http.Client{}
	var data = strings.NewReader(`{"page": 0, "channelId": 25818}`)
	req, err := http.NewRequest("POST", "https://kohcamp.qq.com/info/listinfov2", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, _ := ioutil.ReadAll(resp.Body)
	list := gjson.Get(string(bodyText), "data.list").String()
	num := gjson.Get(list, "#").Int()
	m := []string{}
	for i := 0; i < int(num); i++ {
		data := gjson.Get(list, fmt.Sprintf("%v.infoContent.infoId", i)).String()
		m = append(m, data)
	}
	return &m
}

// 浏览资讯
func (a *Token) read(iInfoIds []string, c *chan int) {
	for _, iInfoId := range iInfoIds {
		data := fmt.Sprintf("iInfoId=%v&gameId=20001&token=%v&userId=%v", iInfoId, a.Token, a.UserId)
		bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/game/detailinfov3", data)
		returnCode := gjson.Get(bodyText, "returnCode").Int()
		if returnCode == 0 {
			fmt.Printf("浏览%v成功\n", iInfoId)
		} else {
			fmt.Println(gjson.Get(bodyText, "returnMsg").String())
		}
	}
	*c <- 1
}

// 喜欢资讯
func (a *Token) like(iInfoIds []string, c *chan int) {
	for _, iInfoId := range iInfoIds {
		data := fmt.Sprintf("iInfoId=%v&like=1&token=%v&userId=%v", iInfoId, a.Token, a.UserId)
		bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/user/addlike", data)
		returnCode := gjson.Get(bodyText, "returnCode").Int()
		if returnCode == 0 {
			fmt.Printf("点赞%v成功\n", iInfoId)
		} else {
			fmt.Println(gjson.Get(bodyText, "returnMsg").String())
		}
	}
	*c <- 1
}

// 运行游戏
func (a *Token) runGame(c *chan int) {
	data := fmt.Sprintf("type=2&token=%v&userId=%v", a.Token, a.UserId)
	bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata", data)
	returnCode := gjson.Get(string(bodyText), "returnCode").Int()
	if returnCode == 0 {
		fmt.Println("启动游戏成功")
	} else {
		fmt.Println(gjson.Get(bodyText, "returnMsg").String())
	}
	*c <- 1
}

// 分享游戏
func (a *Token) shareGame(c *chan int) {
	data := fmt.Sprintf("type=1&token=%v&userId=%v", a.Token, a.UserId)
	bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata", data)
	returnCode := gjson.Get(string(bodyText), "returnCode").Int()
	if returnCode == 0 {
		fmt.Println("分享成功")
	} else {
		fmt.Println(gjson.Get(bodyText, "returnMsg").String())
	}
	*c <- 1
}

// 签到
func (m *Account) sign(UserId string, c *chan int) {
	data := fmt.Sprintf("serverId=1146&roleId=%v&userId=%v&gameOpenid=A&areaId=1&platid=1&algorithm=v2&appid=1105200115&encode=2&openid=%v&sig=%v&source=smoba_zhushou&timestamp=%v&msdkEncodeParam=%v&cSystem=android&msdkToken=%v", m.RoleId, UserId, m.OpenId, m.Sig, m.Timestamp, m.MsdkEncodeParam, m.MsdkToken)
	bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/play/h5sign", data)
	returnCode := gjson.Get(string(bodyText), "returnCode").Int()
	if returnCode == 0 {
		fmt.Println("签到成功")
	} else {
		fmt.Println(gjson.Get(string(bodyText), "returnMsg").String())
	}
	*c <- 1
}

// 领取奖励
func (m *Account) taskgift(userId string, c *chan int) {
	// 2019072200001 启动游戏
	// 2019071900008 支持点赞
	// 2019071900007 浏览资讯
	// 2019071900006 关注作者
	// 2019071900005 获得胜利
	// 2019071900004 进行游戏
	// 2019071900003 分享内容
	taskId := [7]int64{2019072200001, 2019071900008, 2019071900007, 2019071900006, 2019071900005, 2019071900004, 2019071900003}
	for i := 0; i < 7; i++ {
		data := fmt.Sprintf("serverId=1146&roleId=%v&userId=%v&gameOpenid=A&areaId=1&platid=1&algorithm=v2&appid=1105200115&encode=2&openid=%v&sig=%v&source=smoba_zhushou&timestamp=%v&msdkEncodeParam=%v&cSystem=android&msdkToken=%v&taskId=%v", m.RoleId, userId, m.OpenId, m.Sig, m.Timestamp, m.MsdkEncodeParam, m.MsdkToken, taskId[i])
		bodyText := postHttp("https://ssl.kohsocialapp.qq.com:10001/play/h5taskgetgift", data)
		returnCode := gjson.Get(string(bodyText), "returnCode").Int()
		if returnCode == 0 {
			fmt.Println("领取成功")
		} else {
			fmt.Println(gjson.Get(string(bodyText), "returnMsg").String())
		}
	}
	*c <- 1
}

func entrance(ctx context.Context, event DefineEvent) (string, error) {
	for _, v := range *readFile() {
		c := make(chan int, 6)
		a := *v.getToken()
		paper := *mainPage()
		go a.like(paper, &c)
		go a.read(paper, &c)
		go a.runGame(&c)
		go a.shareGame(&c)
		go v.sign(a.UserId, &c)
		go v.taskgift(a.UserId, &c)
		for i := 0; i < 6; i++ {
			<-c
		}
	}
	return "成功", nil
}

func main() {
	cloudfunction.Start(entrance)
}
