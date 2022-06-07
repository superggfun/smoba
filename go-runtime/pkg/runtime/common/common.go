package common

import (
    "net/http"
)

type RuntimeLogger interface {
    Logf(format string, args ...interface{})
}

type InvokeResponse struct {
    StatusCode int
    Payload []byte
}

type InvokeRequest struct {
    Payload []byte
    Header http.Header
}

type HealthCheckRequest struct {
}


type HealthCheckResponse struct {
}
