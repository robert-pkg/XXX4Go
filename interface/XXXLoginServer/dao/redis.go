package dao

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/robert-pkg/micro-go/log"
)

func getUserTokenKeyForRedis(userID int64, deviceType int, token string) string {
	return fmt.Sprintf("token:%d:%d:%s", userID, deviceType, token)
}

// DeleteUserTokenByRedis .
func (dao *Dao) DeleteUserTokenByRedis(userID int64, deviceType int, token string) error {

	key := getUserTokenKeyForRedis(userID, deviceType, token)
	conn := dao.redisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("del", key); err != nil {
		log.Error("err", "err", err)
		return err
	}

	return nil
}

// SetUserToken2Redis .
func (dao *Dao) SetUserToken2Redis(userID int64, deviceType int, token string, expireTimeStamp int64, createOnly bool) error {

	key := getUserTokenKeyForRedis(userID, deviceType, token)
	conn := dao.redisPool.Get()
	defer conn.Close()

	strExpireTimeStamp := fmt.Sprintf("%d", expireTimeStamp)
	if createOnly {
		conn.Do("SET", key, strExpireTimeStamp, "EX", 86400, "NX") // 1天有效
		return nil
	}

	// 3天有效
	if _, err := conn.Do("SET", key, strExpireTimeStamp, "EX", 86400*3); err != nil {
		log.Error("err", "err", err)
		return err
	}

	return nil
}

// GetUserTokenFromRedis .
func (dao *Dao) GetUserTokenFromRedis(userID int64, deviceType int, token string) (exist bool, expireTimeStamp int64, err error) {

	key := getUserTokenKeyForRedis(userID, deviceType, token)
	conn := dao.redisPool.Get()
	defer conn.Close()

	if expireTimeStamp, err = redis.Int64(conn.Do("GET", key)); err != nil {
		if err != redis.ErrNil {
			return false, 0, err
		}

		return false, 0, nil
	}

	return true, expireTimeStamp, nil
}

func getVerifyCodeKeyForRedis(mobile string, verifyCode string) string {
	return fmt.Sprintf("vcode:%s:%s", mobile, verifyCode)
}

// GetVerifyCodeFromRedis .
func (dao *Dao) GetVerifyCodeFromRedis(mobile string, verifyCode string) (exist bool, expireTimeStamp int64, err error) {

	key := getVerifyCodeKeyForRedis(mobile, verifyCode)
	conn := dao.redisPool.Get()
	defer conn.Close()

	if expireTimeStamp, err = redis.Int64(conn.Do("GET", key)); err != nil {
		if err != redis.ErrNil {
			return false, 0, err
		}

		return false, 0, nil
	}

	return true, expireTimeStamp, nil

}

// SetVerifyCode2Redis .
func (dao *Dao) SetVerifyCode2Redis(mobile string, verifyCode string, expireTimeStamp int64) error {

	key := getVerifyCodeKeyForRedis(mobile, verifyCode)
	conn := dao.redisPool.Get()
	defer conn.Close()

	strExpireTimeStamp := fmt.Sprintf("%d", expireTimeStamp)

	// 5分钟有效
	if _, err := conn.Do("SET", key, strExpireTimeStamp, "EX", 60*5); err != nil {
		log.Error("err", "err", err)
		return err
	}

	return nil
}
