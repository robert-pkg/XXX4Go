package dao

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/robert-pkg/micro-go/log"
)

// GetUserID 获取用户ID
func (dao *Dao) GetUserID(mobile string, autoCreate bool) (userID int64, err error) {

	if userID, err = dao.getUserID(mobile); err != nil {
		return
	}

	if userID > 0 {
		// 已经存在，直接返回
		return
	}

	// 不存在，则插入
	userID, err = dao.insertUser(mobile)
	if err != nil {
		if strings.Contains(err.Error(), "1062: Duplicate entry") {
			// 既然重复插入了，那数据一定存在
			return dao.getUserID(mobile)
		}

		log.Error("err", "err", err)
		// 数据库错误
		return
	}

	// 成功
	return
}

func (dao *Dao) getUserID(mobile string) (userID int64, err error) {

	var item struct {
		UserID int64 `gorm:"column:user_id"`
	}

	err = dao.gormDB.Table("user").Select("user_id").
		Where("mobile=?", mobile).
		Where("is_deleted=1").First(&item).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error("err", "err", err)
			return
		}

		return 0, nil
	}

	return item.UserID, nil
}

func (dao *Dao) insertUser(mobile string) (userID int64, err error) {

	sql := "insert into user (mobile) values (?)"

	db := dao.gormDB.DB()
	result, err := db.Exec(sql, mobile)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// SaveUserToken .
func (dao *Dao) SaveUserToken(userID int64, deviceType int, token string, expireTimeStamp int64) (err error) {

	sql := "insert into user_token (user_id,device_type,token,expire_ts) values (?,?,?,?)"

	db := dao.gormDB.DB()
	_, err = db.Exec(sql, userID, deviceType, token, expireTimeStamp)
	if err != nil {
		log.Error("err", "err", err)
		return
	}

	return
}

// GetUserToken .
func (dao *Dao) GetUserToken(userID int64, deviceType int, token string) (exist bool, expireTimeStamp int64, err error) {

	var item struct {
		ExpireTS int64 `gorm:"column:expire_ts"`
	}

	err = dao.gormDB.Table("user_token").Select("expire_ts").
		Where("user_id=?", userID).
		Where("device_type=?", deviceType).
		Where("token=?", token).
		Where("is_deleted=1").
		Order("expire_ts desc").
		First(&item).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error("err", "err", err)
			return
		}

		exist = false
		err = nil
		return
	}

	exist = true
	expireTimeStamp = item.ExpireTS
	return
}
