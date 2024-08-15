package database

import (
	"Go_rest_now_android/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=uk888888 dbname=testgorm sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Tokens{}, &model.User{}, model.Task{})

	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavilable. Wait for %d sec \n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}

	return dbase
}
