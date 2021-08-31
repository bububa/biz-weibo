package model

import "fmt"

// Response api response interface
type Response interface {
	IsError() bool
	Error() string
}

// BaseResponse api error
type BaseResponse struct {
	TraceID string `json:"trace_id,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// IsError check if the response is error
func (r BaseResponse) IsError() bool {
	return r.Code > 0
}

// Error implement error interface
func (r BaseResponse) Error() string {
	return fmt.Sprintf("%s:[%d]%s", r.TraceID, r.Code, r.Message)
}
