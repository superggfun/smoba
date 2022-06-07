package runtime

import (
    "huaweicloud.com/go-runtime/pkg/runtime/fnhandler"
    "log"
    "net/http"
    "net/rpc"
    "os"
)

// Valid function signatures:
// 	func Handler(payload []byte, ctx context.RuntimeContext) (interface{}, error)
func Register(handler interface{}) {
    wrappedHandler := fnhandler.NewHandler(handler)
    RegisterHandler(wrappedHandler)
}

func RegisterHandler(handler fnhandler.IRequestHandler) {
    err := rpc.Register(fnhandler.NewFunction(handler))
    if err != nil {
        log.Fatal("failed to register handler function")
    }
    rpc.HandleHTTP()
    runtimeAddr := os.Getenv("RUNTIME_API_ADDR")
    err = http.ListenAndServe(runtimeAddr, nil)
    if err != nil {
        log.Fatal(err)
    }
}