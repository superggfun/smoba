package cts

import (
    "fmt"
)

type CTSTriggerEvent struct {
    Cts CTS  `json:"cts"`
}

func (e *CTSTriggerEvent) String() string {
    return fmt.Sprintf(`CTSTriggerEvent{
                                  cts=%+v
                               }`, e.Cts)
}

