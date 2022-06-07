package smn

import (
    "fmt"
)

type SMNBody struct {
    TopicUrn string `json:"topic_urn"`
    TimeStamp string `json:"timestamp"`
    MessageAttributes map[string]string `json:"message_attributes"`
    Message string `json:"message"`
    Type string `json:"type"`
    MessageId string `json:"message_id"`
    Subject string `json:"subject"`
}

func (b *SMNBody) String() string {
    return fmt.Sprintf(`SMNBody{
                                 topic_urn=%v,
                                 timestamp=%v,
                                 message_attributes=%v,
                                 message=%v,
                                 type=%v,
                                 message_id=%v,
                                 subject=%v
                               }`, b.TopicUrn, b.TimeStamp, b.MessageAttributes, b.Message, b.Type, b.MessageId, b.Subject)
}