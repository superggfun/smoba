package main

import (
	"fmt"

	"github.com/superggfun/smoba/config"
	"github.com/superggfun/smoba/doTask"
)

/*
type Token struct {
	UserId string
	Token  string
}

type tasklist struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	data       `json:"data"`
}

type data struct {
	Task []task `json:"taskList"`
}

type task struct {
	Title         string  `json:"title"`
	Desc          string  `json:"desc"`
	Gifts         []gifts `json:"gifts"`
	Func          string  `json:"func"`
	TaskId        string  `json:"taskId"`
	FinishStatus  bool    `json:"finishStatus"`
	PackageStatus bool    `json:"packageStatus"`
}

type gifts struct {
	PackageId int    `json:"packageId"`
	Name      string `json:"name"`
}

// Post请求
func doTask(url string, data string) string {
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
	return string(bodyText)

}

// Post请求
func (m *Account) doGift(url string, data string) []byte {
	client := &http.Client{}
	d := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("timestamp", " 1654165720885")
	req.Header.Add("userId", "460014701")
	req.Header.Add("algorithm", "v2")
	req.Header.Add("openid", "EC0D60B5B571457C1FDA4041B9FEFA27")
	req.Header.Add("encode", "2")
	req.Header.Add("roleId", "34879425")
	req.Header.Add("source", "smoba_zhushou")
	req.Header.Add("msdkToken", "dquzCwHq")
	req.Header.Add("gameOpenid", "8C7A5CE953A1E8A10B1CD2029A8DCFEB")
	req.Header.Add("msdkEncodeParam", "3234183CBFCDC80F8B7EED69710E3EC52182EFD2429A0087254C51AF6A316C0544E1B190A480D1471D9070349FAAC6993A8B867316CEBBEF4E1A69636DB6C296F1A06690ACE4C3790F7E1B5676A29DA2A406DDBB653F0BDE5DD03247FC9AE281C53CDB13D4A31132F12F9D04CD8B97FC86FDA6022CB3AFE6DC128615B2E39F20BEC8F6737DFB6C415A3687F5270187F17BEA955CD746D8E0E0139C91E4D732702D98D98E55E11338CCF8D9A7C769571D")
	req.Header.Add("gameId", "20001")
	req.Header.Add("sig", "f829f30ea7e75d394589834890d4f288")
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

// 获取游戏账号信息
func (m *Account) getToken() {
	data := fmt.Sprintf(`accessToken=%v&loginType=openSdk&openId=%v&cSystem=android&gameId=20001`, m.AccessToken, m.OpenId)
	bodyText := doTask("https://ssl.kohsocialapp.qq.com:10001/user/login", data)
	fmt.Println(string(bodyText))

}

/*
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
func (m *Account) viewedNews(iInfoIds []string, a Token) {
	for _, iInfoId := range iInfoIds {
		data := fmt.Sprintf("iInfoId=%v&gameId=20001&token=%v&userId=%v", iInfoId, a.Token, a.UserId)
		bodyText := m.doTask("https://ssl.kohsocialapp.qq.com:10001/game/detailinfov3", data)
		returnCode := gjson.Get(bodyText, "returnCode").Int()
		if returnCode == 0 {
			fmt.Printf("浏览%v成功\n", iInfoId)
		} else {
			fmt.Println(gjson.Get(bodyText, "returnMsg").String())
		}
	}

}

//支持点赞
func (m *Account) addedLikes(iInfoIds []string, a Token) {
	for _, iInfoId := range iInfoIds {
		data := fmt.Sprintf("iInfoId=%v&like=1&token=%v&userId=%v", iInfoId, a.Token, a.UserId)
		bodyText := m.doTask("https://ssl.kohsocialapp.qq.com:10001/user/addlike", data)
		returnCode := gjson.Get(bodyText, "returnCode").Int()
		if returnCode == 0 {
			fmt.Printf("点赞%v成功\n", iInfoId)
		} else {
			fmt.Println(gjson.Get(bodyText, "returnMsg").String())
		}
		time.Sleep(930 * time.Millisecond)
	}
}

// 运行游戏
func (m *Account) runGame(a Token) {
	data := fmt.Sprintf("type=2&token=%v&userId=%v", a.Token, a.UserId)
	bodyText := m.doTask("https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata", data)
	returnCode := gjson.Get(string(bodyText), "returnCode").Int()
	if returnCode == 0 {
		fmt.Println("启动游戏成功")
	} else {
		fmt.Println(gjson.Get(bodyText, "returnMsg").String())
	}

}

// 分享游戏
func (m *Account) shareGame(a Token) {
	data := fmt.Sprintf("type=1&token=%v&userId=%v", a.Token, a.UserId)
	bodyText := m.doTask("https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata", data)
	returnCode := gjson.Get(string(bodyText), "returnCode").Int()
	if returnCode == 0 {
		fmt.Println("分享成功")
	} else {
		fmt.Println(gjson.Get(bodyText, "returnMsg").String())
	}

}

// 签到
func (m *Account) sign(UserId string) {
	data := `{"cSystem":"android","h5Get":1,"roleId":"1404665258"}`
	bodyText := m.doGift("https://kohcamp.qq.com/operation/action/signin", data)
	returnCode := gjson.Get(string(bodyText), "returnCode").Int()
	if returnCode == 0 {
		fmt.Println("签到成功")
	} else {
		fmt.Println(gjson.Get(string(bodyText), "returnMsg").String())
	}
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
		bodyText := m.doGift("https://ssl.kohsocialapp.qq.com:10001/play/h5taskgetgift", data)
		returnCode := gjson.Get(string(bodyText), "returnCode").Int()
		if returnCode == 0 {
			fmt.Println("领取成功")
		} else {
			fmt.Println(gjson.Get(string(bodyText), "returnMsg").String())
		}
	}
	*c <- 1
}

func (m *Account) getTask() {
	data := `{"cSystem":"android","h5Get":1,"serverId":"1073","roleId":"1404665258"}`
	bodyText := m.doGift("https://kohcamp.qq.com/operation/action/tasklist", data)
	var tasks tasklist
	err := json.Unmarshal(bodyText, &tasks)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tasks)

}
*/

func main() {

	for _, v := range *config.ReadFile() {
		a := doTask.Input(v)
		a.GetToken()
		list, err := a.TaskList()
		if err != nil {
			fmt.Println(err)
		}
		for _, value := range list {
			if value.FinishStatus == 0 {
				fmt.Println(value.TaskId)
			}
		}

		/*
			a.ShareGame()
			a.RunGame()
			//fmt.Println(a.AddedLikes())
			fmt.Println(a.Subscribe())
		*/

	}

	/*
		for _, v := range *readFile() {
			v.getToken()
			//fmt.Println(a.UserId, a.Token)
			//v.read(*mainPage(), a)
			//v.like(*mainPage(), a)
			//v.shareGame(a)
			//v.runGame(a)

			//v.sign(a.UserId)
			//a.shareGame()
			//v.getTask()

		}
	*/

}
