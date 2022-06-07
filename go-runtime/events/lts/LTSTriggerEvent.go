package lts

import (
    "fmt"
)

type LTSTriggerEvent struct {
    Lts LTSBody `json:"lts"`
}

func (e *LTSTriggerEvent) String() string {
    return fmt.Sprintf(`LTSTriggerEvent{
                                 lts=%+v
                               }`, e.Lts)
}