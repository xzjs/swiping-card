package lib

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbonce sync.Once

func DB() *gorm.DB {
	dbonce.Do(func() {
		conf := Conf().DB
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Pass, conf.Host, conf.Port, conf.DBName)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}
