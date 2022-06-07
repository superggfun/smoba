package cts

import (
    "fmt"
)

type CTS struct {
    Time string `json:"time"`
    User User `json:"user"`
    Request map[string]string `json:"request"`
    Response map[string]string `json:"response"`
    Code int `json:"code"`
    ServiceType string `json:"service_type"`
    ResourceType string `json:"resource_type"`
    ResourceName string `json:"resource_name"`
    ResourceId string `json:"resource_id"`
    TraceName string `json:"trace_name"`
    TraceType string `json:"trace_type"`
    RecordTime string `json:"record_time"`
    TraceId string `json:"trace_id"`
    TraceStatus string `json:"trace_status"`
}

func (cts *CTS) String() string {
    return fmt.Sprintf(`CTS{
                                  time='%v',
                                  user=%+v,
                                  request=%v,
                                  response=%v,
                                  code=%v,
                                  service_type='%v',
                                  resource_type='%v',
                                  resource_name='%v',
                                  resource_id='%v',
                                  trace_name='%v',
                                  trace_type='%v',
                                  record_time='%v',
                                  trace_id='%v',
                                  trace_status='%v'
                               }`, cts.Time, cts.User, cts.Request, cts.Response, cts.Code, cts.ServiceType, cts.ResourceType,
                                   cts.ResourceName,cts.ResourceId, cts.TraceName, cts.TraceType, cts.RecordTime, cts.TraceId,
                                   cts.TraceStatus)
}