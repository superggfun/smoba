package account

import (
	"encoding/json"
	"errors"
	"fmt"
)

type tasklist struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	Data       `json:"data"`
}

type Data struct {
	Task []task `json:"taskList"`
}

type task struct {
	Title         string  `json:"title"`
	Desc          string  `json:"desc"`
	Gifts         []gifts `json:"gifts"`
	Func          string  `json:"func"`
	TaskId        string  `json:"taskId"`
	FinishStatus  int8    `json:"finishStatus"`
	PackageStatus int8    `json:"packageStatus"`
}

type gifts struct {
	PackageId int    `json:"packageId"`
	Name      string `json:"name"`
}

// 发送请求并解析响应
func (m *Account) sendGiftRequestAndDecode(url string, data string, v interface{}) error {
	response, err := m.DoGift(url, data, m.UserId, m.OriginalRoleId)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, v)
	if err != nil {
		return err
	}

	switch j := v.(type) {
	case *tasklist:
		if j.ReturnCode != 0 {
			return errors.New(j.ReturnMsg)
		}
	}

	return nil
}

// 获取任务列表
func (m *Account) TaskList() ([]task, error) {
	data := fmt.Sprintf(`{"cSystem":"android","h5Get":1,"serverId":"%v","roleId":%v}
	`, m.ServerId, m.RoleId)
	var tasks tasklist
	err := m.sendGiftRequestAndDecode("https://kohcamp.qq.com/operation/action/tasklist", data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks.Data.Task, nil
}
