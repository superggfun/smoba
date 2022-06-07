package dis

import (
    "fmt"
)

type DISTriggerEvent struct {
    ShardID string
    Message DISMessage
    Tag string
    StreamName string
}

func (e *DISTriggerEvent) String() string {
    return fmt.Sprintf(`DISTriggerEvent{
                                  ShardID=%v,
                                  Message=%+v,
                                  Tag=%v,
                                  StreamName=%v
                               }`, e.ShardID, e.Message, e.Tag, e.StreamName)
}