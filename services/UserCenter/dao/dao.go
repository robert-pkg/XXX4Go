package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/robert-pkg/micro-go/db/mysql"
	"github.com/robert-pkg/micro-go/log"
)

// Dao dao
type Dao struct {
	gormDB *gorm.DB
}

// New new a dao.
func New(dbConfig *mysql.Config) (*Dao, error) {

	if dbConfig == nil {
		return nil, errors.New("no db config")
	}

	gormDB, err := mysql.InitDb(dbConfig)
	if err != nil {
		log.Error("err", "err", err)
		return nil, err
	}

	d := &Dao{
		gormDB: gormDB,
	}
	return d, nil
}
