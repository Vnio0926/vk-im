package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"vk-im/util/errors"
)

var (
	Db   *gorm.DB
	once sync.Once
)

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vnio_im?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = db
	return
}

func GetDb() *gorm.DB {
	//once.Do(func() {
	//	Init()
	//})
	Init()
	return Db
}

func CreateTable(dst interface{}) error {
	db := GetDb()
	err := db.AutoMigrate(dst)
	if err != nil {
		return errors.ErrDbCreateTableFail()
	}
	return err
}
