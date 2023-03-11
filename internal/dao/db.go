package dao

import (
	"catlinie_test01/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
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
	db, err := gorm.Open(databaseSetting.DBType, s)
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
