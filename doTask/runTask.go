package doTask

import (
	"errors"
	"fmt"
)

func (m *Account) RunTask(TaskId string) ([]string, error) {
	var any []string
	var err error
	fmt.Println(TaskId)
	switch TaskId {
	case "2019071900006": //关注作者
		any, err = m.Subscribe()

	case "2023091500002": //支持点赞
		any, err = m.AddedLikes()

	case "2019072200001": //启动游戏
		err = m.RunGame()

	case "2024010800004": //分享内容
		err = m.ShareGame()
	case "2024010800001": //浏览资讯
		any, err = m.ViewedNews()
	case "2022101100011": //浏览战绩
		err = m.ViewRecords()
	case "2022061700004": //获得胜利
		err = errors.New("未获得胜利")
	case "2019071900004": //进行游戏
		err = errors.New("未进行游戏")
	default:
		err = errors.New("未知任务")
	}

	return any, err
}
