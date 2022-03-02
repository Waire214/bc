package helper

import "github.com/google/uuid"

type ErrorBody struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
	Source  string      `json:"source"`
}
type ErrorResponse struct {
	ErrorReference uuid.UUID   `json:"error_reference"`
	TimeStamp      string      `json:"timestamp"`
	Errors         []ErrorBody `json:"errors"`
}
type LogStruct struct {
	TimeStamp       string `json:"@timestamp"`
	Version         string `json:"version"`
	Level           string `json:"level"`
	LevelValue      int    `json:"level_value"`
	StatusCode      string `json:"statuscode"`
	Message         string `json:"message"`
	LoggerName      string `json:"logger_name"`
	AppName         string `json:"app_name"`
	Path            string `json:"path"`
	Method          string `json:"method"`
	CorrelationId   string `json:"barafiri-correlation-id"`
	UserAgent       string `json:"user-agent"`
	ResponseTime    string `json:"barafiri-response-time"`
	ApplicationHost string `json:"barafiri-application-host"`
	RemoteIP        string `json:"remote_ip"`
}
