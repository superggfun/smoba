package timer

import (
	"fmt"
)

type TimerTriggerEvent struct {
	Version     string `json:"version"`
	Time        string `json:"time"`
	TriggerName string `json:"trigger_name"`
	TriggerType string `json:"trigger_type"`
	UserEvent   string `json:"user_event"`
}

func (e *TimerTriggerEvent) String() string {
	return fmt.Sprintf(
		`TimerTriggerEvent{
        version=%v,
        time=%v,
        trigger_name=%v,
        trigger_type=%v,
        user_event=%v
        }`, e.Version, e.Time, e.TriggerName, e.TriggerType, e.UserEvent)
}
