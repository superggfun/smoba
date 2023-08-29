# smoba
* 王者营地自动签到每周浏览点赞任务
* 支持多账号
* 微信推送
* [注意]华为云函数9月1日开始收费，正在转移新战场[LeanCLoud](https://www.leancloud.cn/)，新的[厂库地址](https://github.com/superggfun/smoba_LeanCloud)，目前正在测试。

## 使用手册
[使用手册](https://gosmoba.vercel.app/)

## 修复
修复华为云函数显示 open config.json: no such file or directory & unexpected end of JSON input

说明：之前是可以正常运行的，突然就报这个错误了，可能后面华为云改了读取静态文件的位置。

原因：在咨询过华为云客服之后，原来的config.json路径改成code/config.json路径就能使用，微信推送读取markdown文件同理。读取json配置错误导致后面解析json文件错误。

*如需自行编译使用，请删除源码中读取文件路径前面的"code/"，否则就会出现上面找不到配置文件错误。*

> 函数工作目录权限说明
函数可以读取代码目录下的文件，函数工作目录在入口文件的上一级，例如用户上传了文件夹backend，需要读取与入口文件同级目录的文件test.conf，可以用相对路径“code/backend/test.conf”，或者使用绝对路径（相关目录为RUNTIME_CODE_ROOT环境变量对应的值）。如果需要写文件（如创建新文件或者下载文件等），可以在/tmp目录下进行或者使用函数提供的挂载文件系统功能。
说明：
若容器回收，文件的读写就会失效。
函数目前不支持持久化。
您好您这里尝试添加一下相对路径测试一下。

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
