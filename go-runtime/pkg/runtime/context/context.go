package rtcontext

type ContextEnv struct {
    rtProjectID   string
    rtFcName      string
    rtFcVersion   string
    rtPackage     string
    rtMemory      int
    rtCPU         int
    rtTimeout     int
    rtHanlder     string
    rtUserData    map[string]string
    rtInitializerTimeout     int
    rtInitializerHanlder     string
}

type ContextHTTP struct {
    requestID     string
    accesskey     string
    secretKey     string
    authToken     string
    fcStartTime   int64
    rtRemainTime  int
    securityToken string
}

type ContextProvider struct {
    ctxEnv      *ContextEnv
    ctxHTTPHead *ContextHTTP
}

type userFunctionLog struct {
    requestID string
}

func (l *userFunctionLog) getRequestId() string {
    return l.requestID
}

func (l *userFunctionLog) setRequestId(requestId string) {
    l.requestID = requestId
}


