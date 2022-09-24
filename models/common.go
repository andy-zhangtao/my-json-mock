package models

type CommonResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Token  string      `json:"token,omitempty"`
}

type CommonMockResponse struct {
	Code   int
	Result MockResponse
	Error  error
}

type MockResponse struct {
	ContentType string
	Response    string
}
