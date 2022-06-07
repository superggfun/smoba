package apig

import (
    "encoding/base64"
)

type APIGTriggerResponse struct {
    Body string `json:"body"`
    Headers map[string]string `json:"headers"`
    StatusCode int `json:"statusCode"`
    IsBase64Encoded bool `json:"isBase64Encoded"`
}

func (r *APIGTriggerResponse) SetBase64EncodedBody(body string) {
    r.Body = base64.StdEncoding.EncodeToString([]byte(body))
}


