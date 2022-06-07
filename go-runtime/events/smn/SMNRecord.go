package smn

import (
    "fmt"
)

type SMNRecord struct {
    EventVersion string `json:"event_version"`
    EventSubscriptionUrn string `json:"event_subscription_urn"`
    EventSource string `json:"event_source"`
    Smn SMNBody `json:"smn"`
}

func (r *SMNRecord) String() string {
    return fmt.Sprintf(`SMNRecord{
                                 event_version=%v,
                                 event_subscription_urn=%v,
                                 event_source=%v,
                                 smn=%+v
                               }`, r.EventVersion, r.EventSubscriptionUrn, r.EventSource, r.Smn)
}