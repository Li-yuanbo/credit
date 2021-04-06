package mysql

import (
	"credit_gin/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

//银行的文章表
type BankNews struct {
	Id			int64	`gorm:"column:id"`
	BankId		int64	`gorm:"column:bank_id"`
	NewsTitle	string	`gorm:"column:news_title"`
	NewsContent	string	`gorm:"column:news_content"`
	CreateTime	int64	`gorm:"column:create_time"`
	UpdateTime	int64	`gorm:"column:update_time"`
}

func (*BankNews)TableName() string {
	return "bank_news"
}

func InsertBankNews(req model.PublishNewsReq, db *gorm.DB) (*BankNews, error){
	news := BankNews{
		BankId:      req.BankId,
		NewsTitle:   req.NewsTitle,
		NewsContent: req.NewsContent,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}
	if err := db.Create(&news).Error; err != nil {
		log.Println("create news failed, err: ", err)
		return nil, err
	}
	log.Println("create news success")
	return &news, nil
}