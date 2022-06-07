package dds

import (
    "fmt"
)

type DDSTriggerEvent struct {
    Records []DDSRecord `json:"records"`
}

func (e *DDSTriggerEvent) String() string {
    return fmt.Sprintf(`DDSTriggerEvent{
                                  records=%+v
                               }`, e.Records)
}
