package db

type MockRequest struct {
	Id     int    `json:"id"`
	User   string `json:"user"`
	Mid    string `json:"mid"`
	Method string `json:"method"`
	Path   string `json:"path"`
	Name   string `json:"name"`
	Mock   string `json:"mock"`
	Enable bool   `json:"enable"`
	Params string `json:"params"`
	Remark string `json:"remark"`
}

func (t MockRequest) TableName() string {
	return DB_T_MOCK
}
