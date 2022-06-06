package account

import (
	"github.com/superggfun/smoba/http"
)

type chatroles struct {
	Result     int    `json:"result"`
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	data2      `json:"data"`
}

type data2 struct {
	Roles []roles `json:"roles"`
}

type roles struct {
	ServerId       string `json:"serverId"`
	RoleId         string `json:"roleId"`
	OriginalRoleId string `json:"originalRoleId"`
	RoleName       string `json:"roleName"`
	RoleJob        string `json:"roleJob"`
}

type dejson struct {
	Result     int    `json:"result"`
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	data       `json:"data"`
}

type data struct {
	UserId   string `json:"userId"`
	Token    string `json:"token"`
	UserName string `json:"userName"`
}

type Account struct {
	*http.Account
	UserId         string
	Token          string
	OriginalRoleId string
	RoleId         string
	ServerId       string
	UserName       string
	RoleName       string
	RoleJob        string
}
