package services

import "github.com/andy-zhangtao/my-json-mock/types/db"

type IMockService interface {
	FindAllMock() ([]db.MockRequest, error)
	FindSpecifyMockWithId(id int) (db.MockRequest, error)
	AddNewMock(t db.MockRequest) error
	UpdateMock(t db.MockRequest) error
	DeleteMock(id int) error
}
