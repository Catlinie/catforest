package dao

import (
	"catlinie_test01/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Dao struct {
	Engine *gorm.DB
}

func New(db *gorm.DB) *Dao {
	return &Dao{
		Engine: db,
	}
}

func NewDBEngine(databaseSetting setting.DatabaseSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	s = fmt.Sprintf(s,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.CharSet,
		databaseSetting.ParseTime,
	)
	log.Println(s)
	// Open会初始化Config里面所有非bool类型的属性
	// mysql.Open也只会做两件事，解析出DSN对象并且初始化Dialector的两个属性，分别是DSN、DSNConfig
	db, err := gorm.Open(mysql.Open(s))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d *Dao) beginTrans(method func() error) error {
	tx := d.Engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := method(); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
