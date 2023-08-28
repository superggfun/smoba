package main

import (
	"fmt"
	"log"
	"time"

	"github.com/superggfun/smoba/config"
	"github.com/superggfun/smoba/doGift"
	"github.com/superggfun/smoba/doTask"
	"github.com/superggfun/smoba/wxpush"
	//"huaweicloud.com/go-runtime/go-api/context"
	"github.com/leancloud/go-sdk/leancloud"
)

func run() {
	log.Println("开始运行")
	cfg, err := config.ReadConfigFile("config.json") // code/config.json
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return
	}
	for i, v := range cfg.Account {
		var markdown wxpush.Markdown
		a := doTask.Input(v)
		err := a.GetToken()
		log.Printf("----------账号%v----------\n", i+1)
		log.Printf("ID:%v\n", a.UserId)
		log.Printf("昵称:%v\n", a.UserName)
		log.Printf("账号:%v\n", a.RoleName)

		if err != nil {
			log.Println(err)
			markdown.Err = err
			err := wxpush.PushE(markdown)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("推送成功")
			}
			continue
		} else {
			markdown.UserId = a.UserId
			markdown.UserName = a.UserName
			markdown.RoleName = a.RoleName
			markdown.RoleJob = a.RoleJob
			l := make([]string, 0, 7)

			list, err := a.TaskList()
			if err != nil {
				log.Println(err)
				markdown.Err = err
				err := wxpush.PushE(markdown)
				if err != nil {
					log.Println(err)
				} else {
					log.Println("推送成功")
				}
				continue
			} else {
				for _, value := range list {
					if value.FinishStatus == 0 {
						list, err := a.RunTask(value.TaskId)
						if err != nil {
							l = append(l, fmt.Sprintf("[失败]%v: %v", value.Title, err))
							log.Printf("[失败]%v: %v\n", value.Title, err)
						} else if list != nil {
							l = append(l, fmt.Sprintf("[成功]%v(%v)", value.Title, list))
							log.Printf("[成功]%v(%v)\n", value.Title, list)
						} else {
							l = append(l, fmt.Sprintf("[成功]%v\n", value.Title))
							log.Printf("[成功]%v\n", value.Title)
						}
					}

				}
			}
			if len(l) == 0 {
				l = append(l, "已全部完成")
				log.Println("已全部完成")
			}
			markdown.DoTask = l
		}

		b := doGift.Input(v)
		b.GetToken()
		signData, err := b.Sign()
		if err != nil {
			log.Printf("[失败]%v\n", err)
			markdown.SignB = false
			markdown.Sign = fmt.Sprintf("[失败]%v", err)
		} else {
			log.Println("[成功]签到")
			markdown.SignB = true
			markdown.Sign = "[成功]签到"
			markdown.Good = signData.Good
			markdown.Bad = signData.Bad
			markdown.Lunar = signData.Lunar
		}

		list, err := b.TaskList()
		if err != nil {
			log.Println(err)
			markdown.Err = err
			err := wxpush.PushE(markdown)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("推送成功")
			}
			continue
		}
		m := make(map[string]struct{}, 7)
		s := make([]string, 0, 7)
		for _, value := range list {
			if value.PackageStatus == 0 {
				m[value.TaskId] = struct{}{}
				s = append(s, value.TaskId)
			}
		}
		err = b.RunGift(s)
		if err != nil {
			log.Println(err)
		}
		list, err = b.TaskList()
		if err != nil {
			log.Println(err)
			markdown.Err = err
			err := wxpush.PushE(markdown)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("推送成功")
			}
			continue
		}
		l := make([]string, 0, 7)
		for _, value := range list {
			if value.PackageStatus == 0 {
				if _, ok := m[value.TaskId]; ok {
					log.Printf("[失败]%v\n", value.Title)
					l = append(l, fmt.Sprintf("[失败]%v", value.Title))
				}
			} else {
				if _, ok := m[value.TaskId]; ok {
					log.Printf("[成功]%v\n", value.Title)
					l = append(l, fmt.Sprintf("[成功]%v", value.Title))
				}
			}
		}
		if len(l) == 0 {
			l = append(l, "已全部领取")
			log.Println("已全部领取")
		}
		markdown.DoGift = l
		markdown.Time = time.Now().Format("2006-01-02 15:04:05")
		err = wxpush.Push(markdown)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("推送成功")
		}
	}

}

func executeTask(req *leancloud.FunctionRequest) (interface{}, error) {
    // 这里添加您的实际定时任务逻辑
	run()
    return map[string]string{
        "status": "success",
    }, nil
}

func main() {
    // 初始化 LeanCloud 客户端
    client := leancloud.NewEnvClient()
    leancloud.Engine.Init(client)

    // 定义云函数
    leancloud.Engine.Define("executeTaskFunction", executeTask)

    // 在此，您不需要显式地启动任何服务监听，因为 LeanCloud 云函数平台会处理请求并调用您定义的函数。
    // 如果您在本地测试，您可能需要添加某种方式来启动 HTTP 服务并监听请求，但在 LeanCloud 云函数中，这是不必要的。
}
/*
func main() {
	//run()
	client := leancloud.NewEnvClient()
	leancloud.Engine.Init(client)
	//  runtime.Register(ApigTest)
}
/*
func ApigTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	run()
	return "执行完毕", nil
}
*/
