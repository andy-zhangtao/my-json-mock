package services

import (
	mdb "github.com/andy-zhangtao/my-json-mock/db"
	"github.com/andy-zhangtao/my-json-mock/types/config"
	"github.com/andy-zhangtao/my-json-mock/types/db"
	"github.com/sirupsen/logrus"
)

type MockService struct {
	db *mdb.MysqlClient
}

func NewMockService(c config.RunParams) (mc MockService, err error) {
	db, err := mdb.NewMysql(c)
	if err != nil {
		return mc, err
	}

	logrus.Debug("MockService Init Success")
	return MockService{db: &db}, nil
}

func (ts MockService) FindAllMock() ([]db.MockRequest, error) {
	all, err := ts.db.FindAllMock()
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (ts MockService) FindSpecifyMockWithId(id int) (db.MockRequest, error) {
	t, err := ts.db.FindSpecMockWithId(id)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (ts MockService) AddNewMock(t db.MockRequest) error {
	return ts.db.AddMock(t)
}

func (ts MockService) UpdateMock(t db.MockRequest) error {
	return ts.db.UpdateMock(t)
}

func (ts MockService) DeleteMock(id int) error {
	t, err := ts.db.FindSpecMockWithId(id)
	if err != nil {
		return err
	}

	return ts.db.DeleteMock(t)
}
