package smn

import (
    "fmt"
)

type SMNTriggerEvent struct {
    Record []SMNRecord `json:"record"`
}

func (e *SMNTriggerEvent) String() string {
    return fmt.Sprintf(`SMNTriggerEvent{
                                 record=%+v
                               }`, e.Record)
}
