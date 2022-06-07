package lts

import (
    "encoding/base64"
    "fmt"
)

type LTSBody struct {
    Data string `json:"data"`
}

func (b *LTSBody) GetRawData() string {
    res, err := base64.StdEncoding.DecodeString(b.Data)
    if err != nil {
        return ""
    }
    return string(res)
}

func (b *LTSBody) String() string {
    return fmt.Sprintf(`LTSBody{
                                 data='%v'
                               }`, b.Data)
}
