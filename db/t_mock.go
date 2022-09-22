package db

import (
	"github.com/andy-zhangtao/my-json-mock/types/db"
)

func (mc *MysqlClient) AddMock(t db.MockRequest) error {
	return mc.gdb.Create(&t).Error
}

func (mc *MysqlClient) UpdateMock(t db.MockRequest) error {
	return mc.gdb.Model(&t).Where("id = ?", t.Id).Updates(&t).Error
}

func (mc *MysqlClient) FindSpecMockWithId(id int) (db.MockRequest, error) {
	result := db.MockRequest{}
	err := mc.gdb.Find(&db.MockRequest{Id: id}).First(&result).Error
	return result, err
}

func (mc *MysqlClient) FindAllMock() ([]db.MockRequest, error) {
	var ts []db.MockRequest
	err := mc.gdb.Find(&ts).Error
	return ts, err
}

func (mc *MysqlClient) DeleteMock(t db.MockRequest) error {
	return mc.gdb.Delete(&t).Error
}
