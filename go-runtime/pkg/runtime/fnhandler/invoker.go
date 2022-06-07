package fnhandler

import (
    "bytes"
    "encoding/json"
    "fmt"
    "huaweicloud.com/go-runtime/pkg/runtime/common"
    "huaweicloud.com/go-runtime/pkg/runtime/context"
    "log"
    "net/http"
    "os"
    "runtime"
    "strconv"
    "strings"
)

const (
    maxFunctionInvokeStackDepth = 32
    headerCFFInvokeType = "X-CFF-Invoke-Type"
    invokeTypeSync = "sync"
    invokeTypeAsync = "async"
    errmsgBadRequestParameters = "bad request parameters"
)

var (
    maxResponseBodySize = 0
    EmptyStringBytes = []byte("")
    contextObj = rtcontext.GetContextEnvInstance()
)

func init() {
    maxResponseBodySizeEnvValue := os.ExpandEnv("$RUNTIME_MAX_RESP_BODY_SIZE")
    if len(maxResponseBodySizeEnvValue) > 0 {
        value, err := strconv.Atoi(maxResponseBodySizeEnvValue)
        if err != nil {
            log.Println("env 'RUNTIME_MAX_RESP_BODY_SIZE'(%s) invalid.", maxResponseBodySizeEnvValue)
        } else {
            maxResponseBodySize = value
        }
    }

    contextObj.InitiliazeContext()
}

type stack struct {
    File string   `json:"file"`
    Lineno int    `json:"lineno"`
    FunctionName string `json:"function"`
}

type InvokeErrorMessage struct {
    ErrorMessage string   `json:"errorMessage"`
    ErrorType    string   `json:"errorType,omitempty"`
    StackTrace   []string `json:"stackTrace,omitempty"`
}

type InvokeError struct {
    ErrorCode int `json:"error_code"`
    ErrorMsg string `json:"error_msg"`
}

func (e *InvokeError) Error() string {
    return e.ErrorMsg
}

func makeErrorMessage(errMessage, errType string, stacks []*stack) string {
    var stackTraces []string
    if len(stacks) > 0 {
        stackTraces = make([]string, 0)
        for _, stack := range stacks {
            stackTraces = append(stackTraces, fmt.Sprintf("%s()", stack.FunctionName))
            stackTraces = append(stackTraces, fmt.Sprintf("    %s:%d", stack.File, stack.Lineno))
        }
    }

    m := &InvokeErrorMessage{
        ErrorMessage: errMessage,
        ErrorType: errType,
        StackTrace: stackTraces,
    }

    data, err := json.MarshalIndent(m, "", "    ")
    if err != nil {
        log.Println("marshal error message failed.")
        return ""
    }

    return string(data)
}

func makePanicMessage(message string, stackTrace []*stack) (string) {
    return makeErrorMessage(message, "panic", stackTrace)
}

type Function struct {
    handler IRequestHandler
}

func NewFunction(handler IRequestHandler) *Function {
    return &Function{handler: handler}
}

func (fn *Function) Invoke(req *common.InvokeRequest, resp *common.InvokeResponse) (funcErr error) {
    invokeType := req.Header.Get(headerCFFInvokeType)
    if len(invokeType) == 0 {
        invokeType = invokeTypeSync
    }

    contextHTTPHeaderObj := rtcontext.GetContextHTTPHeadInstance(req)
    contextProvider := rtcontext.GetContextProvider(contextObj, contextHTTPHeaderObj)

    var err error
    var payload []byte
    if req.Payload != nil {
        payload = req.Payload
    } else {
        return &InvokeError{
            ErrorCode: http.StatusBadRequest,
            ErrorMsg:  errmsgBadRequestParameters,
        }
    }

    defer func() {
        if e := recover(); e != nil {
            contextProvider.GetLogger().Logf("invoke function  failed for function code panic, error=%+v.", e)
            var panicBuffer bytes.Buffer

            invokeErr := &InvokeError{
                ErrorCode: 555,
            }
            stackTrace := make([]*stack, 0)
            stackCount := 0
            for skip := 1; ; skip++ {
                pc, filePath, lineno, ok:= runtime.Caller(skip)
                if !ok {
                    break
                }

                if strings.HasSuffix(filePath, ".s") {
                    continue
                }

                p := runtime.FuncForPC(pc)
                if p == nil {
                    break
                }

                funcName := p.Name()
                if strings.HasSuffix(funcName, "go-runtime/pkg/runtime/fnhandler.NewHandler.func1") {
                    break
                }
                if !strings.HasPrefix(funcName, "runtime.") && !strings.HasPrefix(funcName, "reflect.") {
                    filePath = hideAbsolutePath(filePath)
                    funcName = formatFuncName(funcName)
                    panicBuffer.WriteString(fmt.Sprintf("%s()\n    %s:%d\n", funcName, filePath, lineno))
                    stackTrace = append(stackTrace, &stack{FunctionName: funcName, File: filePath, Lineno: lineno})
                    stackCount += 1
                    if stackCount >= maxFunctionInvokeStackDepth {
                        break
                    }
                }
            }
            
            if invokeType != invokeTypeAsync {
                invokeErr.ErrorMsg = makePanicMessage(fmt.Sprintf("%+v", e), stackTrace)
            }
            if len(panicBuffer.String()) > 0 {
                fmt.Print(panicBuffer.String())
            }
            funcErr = invokeErr
        }
    }()
    
    var invokeErr error
    invokeResult, err := fn.handler.Handle(payload, contextProvider)
    if err != nil {
        errorMessage := hideAbsolutePath(err.Error())
        invokeErr = &InvokeError{
            ErrorCode: http.StatusInternalServerError,
            ErrorMsg:  makeErrorMessage(errorMessage, "FunctionReturnError", nil),
        }
        contextProvider.GetLogger().Logf("%s", errorMessage)
        return invokeErr
    }

    if invokeType == invokeTypeAsync {
        resp.StatusCode = http.StatusOK
        resp.Payload = EmptyStringBytes
        return nil
    }
    finalResult, _ := transformInvokeResultToBytes(invokeResult)

    if maxResponseBodySize > 0 && len(finalResult) > maxResponseBodySize {
        errorMessage := fmt.Sprintf("Response body size '%d' larger than max value '%d'.", len(finalResult), maxResponseBodySize)
        invokeErr =  &InvokeError{
            ErrorCode: http.StatusInsufficientStorage,
            ErrorMsg:  makeErrorMessage(errorMessage, "FunctionResponseTooLarge", nil),
        }
        contextProvider.GetLogger().Logf("%s", errorMessage)
        return invokeErr
    }
    resp.StatusCode = http.StatusOK
    resp.Payload = finalResult
    return nil
}

func (fn *Function) HealthCheck(req *common.HealthCheckRequest, resp *common.HealthCheckResponse) error {
    *resp = common.HealthCheckResponse{}
    return nil
}

func transformInvokeResultToBytes(invokeResult interface{}) ([]byte, bool) {
    bJsonResult := false
    var result []byte
    switch v := invokeResult.(type) {
    case nil:
        result = EmptyStringBytes
    case string:
        result = []byte(v)
    case []byte:
        result = v
    default:
        data, err := json.Marshal(v)
        if err != nil {
            log.Println("marshal invoke result failed.")
            result = EmptyStringBytes
        } else {
            bJsonResult = true
            result = data
        }
    }

    return result, bJsonResult
}
