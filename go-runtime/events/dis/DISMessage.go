package dis

import (
    "fmt"
)

type DISMessage struct {
    NextPatitionCursor string `json:"next_patition_cursor"`
    Records []DISRecord `json:"records"`
    MillisBehindLatest string `json:"millisBehindLatest"`
}

func (d *DISMessage) String() string {
    return fmt.Sprintf(`DISMessage{
                                 next_patition_cursor=%v,
                                 records=%+v,
                                 millisBehindLatest=%v
                               }`, d.NextPatitionCursor, d.Records, d.MillisBehindLatest)
}