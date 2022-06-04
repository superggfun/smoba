package doTask

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

// 获取任务列表
func (m *Account) TaskList() ([]task, error) {
	data := fmt.Sprintf(`{"cSystem":"android","h5Get":1,"serverId":"1073","roleId":%v}
	`, m.RoleId)
	list := m.DoGift("https://kohcamp.qq.com/operation/action/tasklist", data, m.UserId)
	var tasks tasklist
	var d []task
	err := json.Unmarshal(list, &tasks)
	if err != nil {
		return d, err
	} else if tasks.ReturnCode != 0 {
		return d, errors.New("get tasklist error")
	}
	return tasks.Data.Task, nil
}
