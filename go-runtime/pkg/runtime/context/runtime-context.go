package rtcontext

import (
    "fmt"
    "huaweicloud.com/go-runtime/pkg/runtime/common"
    "sync"
    "time"
)

var (
    once       sync.Once
    contextobj *ContextEnv

    funcLogger = &userFunctionLog{}
)

func (ctxProvider ContextProvider) GetRemainingTimeInMilliSeconds() int {
    currentTime := getCurrentTime()
    usedTime := int(currentTime - ctxProvider.ctxHTTPHead.fcStartTime)
    timeout := ctxProvider.ctxEnv.rtTimeout * 1000
    if usedTime < timeout {
        ctxProvider.ctxHTTPHead.rtRemainTime = timeout - usedTime
    } else {
        ctxProvider.ctxHTTPHead.rtRemainTime = 0
    }
    return ctxProvider.ctxHTTPHead.rtRemainTime
}

func (ctxProvider ContextProvider) GetFunctionName() string {
    return ctxProvider.ctxEnv.rtFcName
}

func (ctxProvider ContextProvider) GetRunningTimeInSeconds() int {
    return ctxProvider.ctxEnv.rtTimeout
}

func (ctxProvider ContextProvider) GetVersion() string {
    return ctxProvider.ctxEnv.rtFcVersion
}

func (ctxProvider ContextProvider) GetMemorySize() int {
    return ctxProvider.ctxEnv.rtMemory
}

func (ctxProvider ContextProvider) GetCPUNumber() int {
    return ctxProvider.ctxEnv.rtCPU
}

func (ctxProvider ContextProvider) GetUserData(key string) string {
    if ctxProvider.ctxEnv.rtUserData != nil {
        return ctxProvider.ctxEnv.rtUserData[key]
    }
    return ""
}

func (ctxProvider ContextProvider) GetLogger() common.RuntimeLogger {
    if funcLogger.getRequestId() != ctxProvider.ctxHTTPHead.requestID {
        funcLogger.setRequestId(ctxProvider.ctxHTTPHead.requestID)
    }

    return funcLogger
}

func (ctxProvider ContextProvider) GetProjectID() string {
    return ctxProvider.ctxEnv.rtProjectID
}

func (ctxProvider ContextProvider) GetPackage() string {
    return ctxProvider.ctxEnv.rtPackage
}

func (ctxProvider ContextProvider) GetHandler() string {
    return ctxProvider.ctxEnv.rtHanlder
}

func (ctxProvider ContextProvider) GetInitializerHandler() string {
    return ctxProvider.ctxEnv.rtInitializerHanlder
}

func (ctxProvider ContextProvider) GetAccessKey() string {
    return ctxProvider.ctxHTTPHead.accesskey
}

func (ctxProvider ContextProvider) GetSecretKey() string {
    return ctxProvider.ctxHTTPHead.secretKey
}

func (ctxProvider ContextProvider) GetToken() string {
    return ctxProvider.ctxHTTPHead.authToken
}

func (ctxProvider ContextProvider) GetRequestID() string {
    return ctxProvider.ctxHTTPHead.requestID
}

func (ctxProvider ContextProvider) GetSecurityToken() string {
    return ctxProvider.ctxHTTPHead.securityToken
}

func (logger *userFunctionLog) Logf(format string, args ...interface{}) {
    logTimeFormat := "2006-01-02 15:04:05.999-07:00"
    content := fmt.Sprintf(format, args...)
    myFormat := fmt.Sprintf("%s %s %s", time.Now().Format(logTimeFormat), logger.requestID, content)
    fmt.Println(myFormat)
}

func GetContextProvider(ctxEnv *ContextEnv, ctxHTTPHead *ContextHTTP) ContextProvider {
    return ContextProvider{ctxEnv, ctxHTTPHead}
}

func GetContextEnvInstance() *ContextEnv {

    once.Do(func() {
        contextobj = new(ContextEnv)
    })
    return contextobj
}

func GetContextHTTPHeadInstance(req *common.InvokeRequest) *ContextHTTP {
    contextHTTPHead := new(ContextHTTP)
    requestID := req.Header.Get("X-CFF-Request-Id")
    if requestID != "" {
        contextHTTPHead.requestID = requestID
    }
    accesskey := req.Header.Get("X-CFF-Access-Key")
    if accesskey != "" {
        contextHTTPHead.accesskey = accesskey
    }
    secretKey := req.Header.Get("X-CFF-Secret-Key")
    if secretKey != "" {
        contextHTTPHead.secretKey = secretKey
    }
    authToken := req.Header.Get("X-CFF-Auth-Token")
    if authToken != "" {
        contextHTTPHead.authToken = authToken
    }
    contextHTTPHead.fcStartTime = getCurrentTime()

    securityToken := req.Header.Get("X-CFF-Security-Token")
    if securityToken != "" {
        contextHTTPHead.securityToken = securityToken
    }
    return contextHTTPHead
}
