package mysql

import (
	"credit_gin/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type CreditCardFlow struct {
	Id 					int64 	`gorm:"column:id"`
	UserId				int64	`gorm:"column:user_id"`
	BankId				int64	`gorm:"column:bank_id"`
	CreditCardNum		string	`gorm:"column:credit_card_num"`
	OperateDesc			string	`gorm:"column:operate_desc"`
	OperateFlow			string	`gorm:"column:operate_flow"`
	OperateType			int64	`gorm:"column:operate_type"`  //0-存款 1-取款 2-转账 3-被转账
	Money				float64	`gorm:"column:money"`
	MoneyType			int64	`gorm:"column:money_type"`    //0-收入 1-支出
	AfterBalance		float64	`gorm:"column:after_balance"`
	CreateTime			int64	`gorm:"column:create_time"`
	UpdateTime			int64	`gorm:"column:update_time"`
}

func (*CreditCardFlow) TableName() string {
	return "credit_card_flow"
}

func InsertOperateFlow(req model.CreditCardFlowModel, db *gorm.DB) (*CreditCardFlow, error) {
	model := CreditCardFlow{
		UserId:        req.UserId,
		BankId:        req.BankId,
		CreditCardNum: req.CreditCardNum,
		OperateDesc:   req.OperateDesc,
		OperateFlow:   req.OperateFlow,
		OperateType:   req.OperateType,
		Money:         req.Money,
		MoneyType:     req.MoneyType,
		AfterBalance:  req.AfterBalance,
		CreateTime:    time.Now().Unix(),
		UpdateTime:    time.Now().Unix(),
	}
	if err := db.Create(&model).Error; err != nil {
		log.Println("[db] insert card operate flow err: ", err)
		return nil, err
	}
	log.Println("[db] insert card operate flow success")
	return &model, nil
}