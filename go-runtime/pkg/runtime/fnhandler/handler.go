package fnhandler

import (
    "fmt"
    "net/http"
    "os"
    "reflect"
    "runtime"
    "strings"

    "huaweicloud.com/go-runtime/go-api/context"
)

var (
    runtimeRoot = os.Getenv("RUNTIME_ROOT")
    realRuntimeRoot = ""
    functionCodePath = os.Getenv("RUNTIME_CODE_ROOT")
    goRoot = runtime.GOROOT()
    goRootSrcPath = fmt.Sprintf("%s/src/", goRoot)
)

func init() {
    realPath, err := os.Readlink(runtimeRoot)
    if err != nil {
        realRuntimeRoot = realPath
    }
}

type FunctionLoadFailedError struct {
    StatusCode   int
    ErrorMessage string
}

func (e *FunctionLoadFailedError) Error() string {
    return e.ErrorMessage
}

func hideAbsolutePath(path string) string {
    result := path
    if strings.Contains(path, goRootSrcPath) {
        result = strings.Replace(result, goRootSrcPath, "", -1)
    } else {
        index := strings.Index(result, "/src/")
        if index >= 0 {
            result = result[index + 5:]
        }
    }

    if len(functionCodePath) > 0 {
        result = strings.Replace(result, functionCodePath, "", -1)
    }

    if len(runtimeRoot) > 0 {
        result = strings.Replace(result, runtimeRoot, "", -1)
    }

    if len(realRuntimeRoot) > 0 {
        result = strings.Replace(result, realRuntimeRoot, "", -1)
    }

    index := strings.Index(result, "go-runtime")
    if index != -1 {
        result = result[index:]
    }

    return result
}

func formatFuncName(name string) string {
    result := name
    index := strings.LastIndex(result, "/")
    if index != -1 {
        result = result[index+1:]
    }
    index = strings.Index(result, ".")
    if index != -1 {
        result = result[index+1:]
    }
    return result
}

// HandlerFunc implements Handler.
type HandlerFunc func([]byte, context.RuntimeContext) (interface{}, error)

type IRequestHandler interface {
    Handle(payload []byte, ctx context.RuntimeContext) (interface{}, error)
}

type RequestHandler struct {
    handlerFunc HandlerFunc
}

func (handler *RequestHandler) Handle(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
    return handler.handlerFunc(payload, ctx)
}

func NewHandler(handlerFunc interface{}) IRequestHandler {
    if handlerFunc == nil {
        return &FaultRequestHandler{err: &FunctionLoadFailedError{
            StatusCode:   http.StatusBadRequest,
            ErrorMessage: "handler is nil",
        }}
    }
    handlerType := reflect.TypeOf(handlerFunc)
    handler := reflect.ValueOf(handlerFunc)
    if handlerType.Kind() != reflect.Func {
        return &FaultRequestHandler{err: &FunctionLoadFailedError{
            StatusCode:   http.StatusBadRequest,
            ErrorMessage: fmt.Sprintf("handler kind %s is not %s", handlerType.Kind(), reflect.Func),
        }}
    }

    return &RequestHandler{handlerFunc: HandlerFunc(func(payload []byte, ctx context.RuntimeContext) (interface{}, error){
        var args []reflect.Value
        args = append(args, reflect.ValueOf(payload))
        args = append(args, reflect.ValueOf(ctx))
        response := handler.Call(args)
        val := response[0].Interface()
        if val == nil {
            return nil,  fmt.Errorf("response is empty")
        }
        var err error
        if errVal, ok := response[1].Interface().(error);ok {
            err = errVal
        }
        return val, err
    })}
}

type FaultRequestHandler struct {
    err *FunctionLoadFailedError
}

func (handler *FaultRequestHandler) Handle(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
    return nil, handler.err
}

