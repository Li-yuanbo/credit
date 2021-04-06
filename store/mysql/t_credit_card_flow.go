package mysql

import (
	"credit_gin/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type CreditCardFlow struct {
	Id 						int 		`gorm:"column:id"`
	CreditCardId			int			`gorm:"column:credit_card_id"`
	OperateUserId			int 		`gorm:"column:operate_user_id"`
	OperateFlow				string		`gorm:"column:operate_flow"` //操作记录
	Desc					string		`gorm:"column:desc"`  //操作备注
	OperateType				int			`gorm:"column:operate_type"` //0-存款 1-取款 2-转账 3-被转账
	OtherUserId				int			`gorm:"column:other_user_id"`//收款人id
	OtherCreditCard			int			`gorm:"column:other_credit_card_id"` //收款人信用卡号
	CreateTime				int64		`gorm:"column:create_time"`
	UpdateTime				int64		`gorm:"column:update_time"`
}

func (*CreditCardFlow) TableName() string {
	return "credit_card_flow"
}

//AddCreditCardFlow: 新增转账/存款/取款记录
func AddCreditCardFlow(req model.CreditCardFlowReq, db *gorm.DB) error{
	flowModel := CreditCardFlow{
		CreditCardId:  		req.CreditCardId,
		OperateUserId: 		req.OperateUserId,
		OperateFlow:   		req.OperateFlow,
		Desc:         		req.Desc,
		OperateType:   		req.OperateType,
		OtherUserId:   		req.OtherUserId,
		OtherCreditCard: 	req.OtherCreditCardId,
		CreateTime:    		time.Now().Unix(),
		UpdateTime:    		time.Now().Unix(),
	}
	if err := db.Create(&flowModel).Error; err != nil {
		log.Println("add operate flow err: ", err)
		return err
	}
	log.Println("add operate flow success, operate flow: ", flowModel)
	return nil
}

//SelectFlowByUserId: 根据user id查询操作记录
func SelectFlowByUserId(userId int, db *gorm.DB) ([]*CreditCardFlow, error) {
	creditCardFlows := make([]*CreditCardFlow, 0, 0)
	if err := db.Model(&CreditCardFlow{}).Where("operate_user_id = ?", userId).Find(&creditCardFlows).Error; err != nil {
		log.Println("select user operate err: ", err)
		return nil, err
	}
	log.Println("select user operate success")
	return creditCardFlows, nil
}

//SelectFlowByCardId: 根据card id查询操作记录
func SelectFlowByCardId(cardId int, db *gorm.DB) ([]*CreditCardFlow, error) {
	creditCardFlows := make([]*CreditCardFlow, 0, 0)
	if err := db.Model(&CreditCardFlow{}).Where("credit_card_id = ?", cardId).Find(&creditCardFlows).Error; err != nil {
		log.Println("select card operate err: ", err)
		return nil, err
	}
	log.Println("select card operate success")
	return creditCardFlows, nil
}