package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"
)

type MariaDBConfig struct {
	Username string
	Host     string
	Password string
	Port     string
	DBName   string
}

type InitMariaDBParams struct {
	Conf   *MariaDBConfig
	Logger *logrus.Entry
}

func InitMariaDB(params *InitMariaDBParams) (db *sql.DB, err error) {
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		params.Conf.Username, params.Conf.Password,
		params.Conf.Host, params.Conf.Port, params.Conf.DBName,
	)

	for i := 10; i > 0; i-- {
		db, err = sql.Open("mysql", dataSource)
		if err == nil {
			break
		}
		params.Logger.Errorf("[InitMariaDB] error init opening db for %s: %+v, retrying in 1 second", dataSource, err)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return
	}

	for i := 10; i > 0; i-- {
		err = db.Ping()
		if err == nil {
			break
		}
		params.Logger.Errorf("[InitMariaDB] error ping db for %s: %+v, retrying in 1 second", dataSource, err)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return
	}

	params.Logger.Infoln("[InitMariaDB] db init successfully")
	return
}
