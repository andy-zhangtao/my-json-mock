package db

import (
	"database/sql"
	"github.com/andy-zhangtao/my-json-mock/types/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	gm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlClient struct {
	cli *sql.DB
	gdb *gorm.DB
}

// NewMysql 初始化Mysql客户端
func NewMysql(c config.RunParams) (ms MysqlClient, err error) {
	logrus.Debugf("Mysql Init Use [%s]", c.Mysql.Dsn)
	db, err := sql.Open("mysql", c.Mysql.Dsn)
	if err != nil {
		err = errors.WithMessage(err, "mysql init")
		return
	}

	logrus.Debug("mysql open success. Then try pong ...")
	err = db.Ping()
	if err != nil {
		err = errors.WithMessage(err, "mysql pong")
		return
	}

	_db, err := gorm.Open(gm.Open(c.Mysql.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return ms, errors.WithMessage(err, "Gorm open mysql")
	}

	logrus.Debug("mysql pong success ...")
	return MysqlClient{cli: db, gdb: _db}, nil
}
