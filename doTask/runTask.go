package doTask

import "errors"

func (m *Account) RunTask(TaskId string) ([]string, error) {
	switch TaskId {
	case "2019071900006": //关注作者
		any, err := m.Subscribe()
		return any, err

	case "2019071900008": //支持点赞
		any, err := m.AddedLikes()
		return any, err

	case "2019072200001": //启动游戏
		err := m.RunGame()
		return nil, err

	case "2019071900003": //分享内容
		err := m.ShareGame()
		return nil, err

	case "2019071900007": //浏览资讯
		any, err := m.ViewedNews()
		return any, err

	case "2019071900005": //获得胜利
		return nil, errors.New("未完成")

	case "2019071900004": //进行游戏
		return nil, errors.New("未完成")

	default:
		return nil, nil
	}
}
