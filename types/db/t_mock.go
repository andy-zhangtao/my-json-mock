package db

type MockRequest struct {
	Id int `json:"id"`

	Name string `json:"name"`

	Mock string `json:"mock"`

	Enable bool `json:"enable"`

	Remark string `json:"remark"`
}

func (t MockRequest) TableName() string {
	return DB_T_MOCK
}
