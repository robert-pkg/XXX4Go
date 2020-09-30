package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/robert-pkg/micro-go/log"
)

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
