package mysql

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const(
	USERNAME = "root"
	PASSWORD = "mysql"
	URL = "127.0.0.1:3306"
	DATABASE = "credit_card"
)

var db *gorm.DB

func init(){
	connStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, URL, DATABASE)
	var err error
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		log.Println("connect database err: ", err)
		return
	}
}

func WriteDB() *gorm.DB{
	return db
}