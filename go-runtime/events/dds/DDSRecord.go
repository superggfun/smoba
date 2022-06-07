package dds

import (
    "fmt"
)

type DDSRecord struct {
    EventName     string            `json:"event_name"`
    EventVersion  string            `json:"event_version"`
    EventSource   string            `json:"event_source"`
    Region        string            `json:"region"`
    Dds           map[string]string `json:"dds"`
    EventSourceId string            `json:"event_source_id"`
}

func (r *DDSRecord) String() string {
    return fmt.Sprintf(`DDSRecord{
                                 event_name=%v,
                                 event_version=%v,
                                 event_source=%v,
                                 region=%v,
                                 dds=%+v,
                                 event_source_id=%v
                               }`, r.EventName, r.EventVersion, r.EventSource, r.Region, r.Dds, r.EventSourceId)
}