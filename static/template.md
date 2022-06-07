## 账户信息
* ID: {{.UserId}}
* 昵称: {{.UserName}}
* 名称: {{.RoleName}}
* 段位: {{.RoleJob}}

## 完成任务 {{range $i, $v := .DoTask}}
* {{$v}}{{end}}

## 签到 
* {{.Sign}}{{if eq .SignB true}}
* 事宜：{{.Good}}
* 禁忌：{{.Bad}}
* 农历: {{.Lunar}}
{{end}}

## 领取奖励 {{range $i, $v := .DoGift}}
* {{$v}}{{end}}

## 时间
* {{.Time}}
