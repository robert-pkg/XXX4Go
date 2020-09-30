package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/robert-pkg/XXX4Go/common/db"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/conf"
	"github.com/robert-pkg/micro-go/log"
)

// Dao dao
type Dao struct {
	gormDB *gorm.DB
}

// New new a dao.
func New(c *conf.Config) (*Dao, error) {

	if c.DBConfig == nil {
		return nil, errors.New("no db config")
	}

	gormDB, err := db.InitDb(c.DBConfig)
	if err != nil {
		log.Error("err", "err", err)
		return nil, err
	}

	d := &Dao{
		gormDB: gormDB,
	}
	return d, nil
}
