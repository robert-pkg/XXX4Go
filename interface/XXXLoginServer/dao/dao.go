package dao

import (
	"errors"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	redisHelper "github.com/robert-pkg/micro-go/cache/redis"
	"github.com/robert-pkg/micro-go/db/mysql"
	"github.com/robert-pkg/micro-go/log"
)

// Dao dao
type Dao struct {
	gormDB    *gorm.DB
	redisPool *redis.Pool
}

// New new a dao.
func New(dbConfig *mysql.Config, redisConfig *redisHelper.Config) (*Dao, error) {

	if dbConfig == nil {
		return nil, errors.New("no db config")
	}

	if redisConfig == nil {
		return nil, errors.New("no redis config")
	}

	gormDB, err := mysql.InitDb(dbConfig)
	if err != nil {
		log.Error("err", "err", err)
		return nil, err
	}

	redisPool := redisHelper.NewRedisPool(redisConfig)

	d := &Dao{
		gormDB:    gormDB,
		redisPool: redisPool,
	}
	return d, nil
}
