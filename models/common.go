package models

type CommonResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Token  string      `json:"token,omitempty"`
}
