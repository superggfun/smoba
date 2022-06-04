package doTask

import (
	"github.com/superggfun/smoba/account"
	"github.com/superggfun/smoba/config"
	"github.com/superggfun/smoba/http"
)

type Account struct {
	*account.Account
}

func Input(a config.Account) *Account {
	var m Account
	m.Account = new(account.Account)
	m.PassWord = new(account.PassWord)
	m.Account.Account = new(http.Account)

	m.AccessToken = a.AccessToken
	m.MsdkEncodeParam = a.MsdkEncodeParam
	m.MsdkToken = a.MsdkToken
	m.OpenId = a.OpenId
	m.RoleId = a.RoleId
	m.Sig = a.Sig
	m.Timestamp = a.Timestamp

	return &m
}
