package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Config .
type Config struct {
	Address   string `yaml:"addr"`
	Password  string `yaml:"pass"`
	MaxActive int    `yaml:"active"`
	MaxIdle   int    `yaml:"idle"`
	DBNumber  int    `yaml:"db_number"`

	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

// NewRedisPool .
func NewRedisPool(cfg *Config) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Address)
			if err != nil {
				return nil, err
			}

			if len(cfg.Password) > 0 {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", cfg.DBNumber); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
