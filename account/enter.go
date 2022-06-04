package account

import (
	"github.com/superggfun/smoba/http"
)

type PassWord struct {
	UserId string
	Token  string
}

type dejson struct {
	Result int `json:"result"`
	data   `json:"data"`
}

type data struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type Account struct {
	*http.Account
	*PassWord
}
