package models

type HttpMockRequest struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Mock   string `json:"mock"`
	Enable bool   `json:"enable"`
	Remark string `json:"remark"`
}
