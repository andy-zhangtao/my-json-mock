package models

type HttpMockRequest struct {
	Id     int                   `json:"id,omitempty"`
	User   string                `json:"user"`
	Name   string                `json:"name"`
	Method string                `json:"method"`
	Mock   string                `json:"mock"`
	Path   string                `json:"path"`
	Enable bool                  `json:"enable"`
	Params HttpMockParamsRequest `json:"params"`
	Remark string                `json:"remark,omitempty"`
}

type HttpMockParamsRequest struct {
	ContentType string `json:"content_type"`
}
