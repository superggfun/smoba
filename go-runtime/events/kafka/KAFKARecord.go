package kafka

import (
    "fmt"
)

type KAFKARecord struct {
    Messages []string `json:"messages"`
    TopicId string `json:"topic_id"`
}

func (r *KAFKARecord) String() string {
    return fmt.Sprintf(`KAFKARecord{
                                 messages:%+v,
                                 topic_id:%s
                               }`, r.Messages, r.TopicId)
}