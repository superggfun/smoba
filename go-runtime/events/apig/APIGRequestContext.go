package apig

import (
	"fmt"
)

type APIGRequestContext struct {
	ApiId     string `json:"apiId"`
	RequestId string `json:"requestId"`
	Stage     string `json:"stage"`
	SourceIp  string `json:"sourceIp"`
}

func (rc APIGRequestContext) String() string {
	return fmt.Sprintf(`APIGRequestContext{
        apiId='%s',
        requestId='%s',
        stage='%s',
        sourceIp='%s',
        }`, rc.ApiId, rc.RequestId, rc.Stage, rc.SourceIp)
}
