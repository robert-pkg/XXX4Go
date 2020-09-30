package dao

import (
	"errors"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	redisHelper "github.com/robert-pkg/XXX4Go/common/cache/redis"
	"github.com/robert-pkg/XXX4Go/common/db"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/conf"
	"github.com/robert-pkg/micro-go/log"
)

// Dao dao
type Dao struct {
	gormDB    *gorm.DB
	redisPool *redis.Pool
}

// New new a dao.
func New(c *conf.Config) (*Dao, error) {

	if c.DBConfig == nil {
		return nil, errors.New("no db config")
	}

	if c.RedisConfig == nil {
		return nil, errors.New("no redis config")
	}

	gormDB, err := db.InitDb(c.DBConfig)
	if err != nil {
		log.Error("err", "err", err)
		return nil, err
	}

	redisPool := redisHelper.NewRedisPool(c.RedisConfig)

	d := &Dao{
		gormDB:    gormDB,
		redisPool: redisPool,
	}
	return d, nil
}
