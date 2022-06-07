package dis

import (
    "fmt"
)

type DISRecord struct {
    PartitionKey string `json:"partition_key"`
    Data string `json:"data"`
    SequenceNumber string `json:"sequence_number"`
}

func (r *DISRecord) String() string {
    return fmt.Sprintf(`DISRecord{
                                 partition_key=%v,
                                 data=%v,
                                 sequence_number=%v
                               }`, r.PartitionKey, r.Data, r.SequenceNumber)
}