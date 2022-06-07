package kafka

import (
    "fmt"
)

type KAFKATriggerEvent struct {
    InstanceId   string        `json:"instance_id"`
    Records      []KAFKARecord `json:"records"`
    TriggerType  string        `json:"trigger_type"`
    Region       string        `json:"region"`
    EventTime    int64         `json:"event_time"`
    EventVersion string        `json:"event_version"`
}

func (e *KAFKATriggerEvent) String() string {
    return fmt.Sprintf(`KAFKATriggerEvent{
                                 instance_id=%v,
                                 records=%+v,
                                 trigger_type=%v,
                                 region=%v,
                                 event_time=%v,
                                 event_version=%v
                               }`, e.InstanceId, e.Records, e.TriggerType, e.Region, e.EventTime, e.EventVersion)
}