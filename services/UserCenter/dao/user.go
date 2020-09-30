package dao

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/robert-pkg/micro-go/log"
)

// GetUserID 获取用户ID
func (dao *Dao) GetUserID(mobile string, autoCreate bool) (userID int64, isNew bool, err error) {

	if userID, err = dao.getUserID(mobile); err != nil {
		return
	}

	if userID > 0 {
		// 已经存在，直接返回
		isNew = false
		return
	}

	// 不存在，则插入
	if userID, err = dao.insertUser(mobile); err != nil {
		if strings.Contains(err.Error(), "1062: Duplicate entry") {
			// 既然重复插入了，那数据一定存在, 在查一次
			if userID, err = dao.getUserID(mobile); err != nil {
				return
			}

			if userID > 0 {
				// 已经存在，直接返回
				isNew = false
				return
			}

			err = errors.New("查询数据出错")
		}

		log.Error("err", "err", err)
		// 数据库错误
		return
	}

	// 成功
	isNew = true
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
