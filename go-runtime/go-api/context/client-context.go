package context

import (
	"huaweicloud.com/go-runtime/pkg/runtime/common"
)

type RuntimeContext interface {
	GetRequestID() string

	GetRemainingTimeInMilliSeconds() int

	GetAccessKey() string

	GetSecretKey() string

	GetFunctionName() string

	GetUserData(string) string

	GetLogger() common.RuntimeLogger

	GetRunningTimeInSeconds() int

	GetVersion() string

	GetMemorySize() int

	GetCPUNumber() int

	GetProjectID() string

	GetPackage() string

	GetToken() string
	
	GetSecurityToken() string
}
