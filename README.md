# smoba
* 王者营地自动签到每周浏览点赞任务
* 支持多账号
* 微信推送
* 华为云函数，免费自动每天刷

## 使用手册
[使用手册](https://gosmoba.vercel.app/)

## 修复
修复华为云函数显示 open config.json: no such file or directory & unexpected end of JSON input
说明：之前是可以正常运行的，突然就报这个错误了，可能后面华为云改了读取静态文件的位置。
原因：在咨询过华为云客服之后，原来的config.json路径改成code/config.json路径就能使用，微信推送读取markdown文件同理。读取json配置错误导致后面解析json文件错误。
*如需自行编译使用，请删除源码中读取文件路径前面的"code/"，否则就会出现上面找不到配置文件错误。

## 编译使用
*如果不使用华为云函数请在函数入口main下直接使用run()函数*
```go
func main() {
	run()
}
```
*编译(Linux) 华为云函数*
```Terminal
go mod tidy
set GOOS=linux
set GOARCH=amd64
go build -o handler main.go
zip main.zip handler config.go static
```

*编译(Windows)*
```cmd
go mod tidy
set GOOS=windows
set GOARCH=386
go build -o handler main.go
```

## 免责声明
* 本厂库内容仅用于学习研究，禁止用于任何商业或非法用途，违反者作者概不负责。
* 如果任何单位或个人认为该厂库内容可能涉嫌侵犯其权利。则应及时通知并提供身份证明，所有权证明，我将在收到认证文件后删除厂库内容。
