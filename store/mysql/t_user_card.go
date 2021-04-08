package mysql

import (
	"credit_gin/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type CreditCard struct {
	Id				int64 		`gorm:"column:id"`
	BankId			int64		`gorm:"column:bank_id"`
	UserId			int64		`gorm:"column:user_id"`
	CreditCardNum	string		`gorm:"column:credit_card_num"`
	Balance			float64		`gorm:"column:balance"`
	CreateTime		int64		`gorm:"column:create_time"`
	UpdateTime		int64		`gorm:"column:update_time"`
}

func (*CreditCard)	TableName() string {
	return "user_card"
}

//AddCreditCard
func AddCreditCard(req model.CreditCardReq, db *gorm.DB) (*CreditCard, error) {
	creditCardModel := CreditCard{
		BankId:        req.BankId,
		UserId:        req.UserId,
		CreditCardNum: req.CreditCardNum,
		CreateTime:    time.Now().Unix(),
		UpdateTime:    time.Now().Unix(),
	}
	if err := db.Create(&creditCardModel).Error; err != nil {
		log.Println("add credit card err: ", err)
		return nil, err
	}
	log.Println("add credit card success")
	return &creditCardModel, nil
}

//UpdateCreditCard: 更新信用卡信息
func UpdateCreditCard(req model.CreditCardReq, db *gorm.DB) error {
	creditCardModel := CreditCard{
		Balance:       req.Balance,
		UpdateTime:    time.Now().Unix(),
	}
	if err := db.Model(&CreditCard{}).Where("credit_card_num = ?", req.CreditCardNum).Update(creditCardModel).Error; err != nil {
		log.Println("update credit card err: ", err)
		return err
	}
	log.Println("update credit card success")
	return nil
}

//DeleteCreditCard：删除信用卡信息
func DeleteCreditCard(card model.CreditCardReq, db *gorm.DB) error {
	creditCardModel := CreditCard{
		Id:            card.Id,
		UserId:		   card.UserId,
	}
	var c CreditCard
	if err := db.Where(&creditCardModel).First(&c).Error; err != nil{
		if gorm.IsRecordNotFoundError(err) {
			log.Println("not have card err: ", err, ". card id is ", card.Id, ". user id is ", card.UserId)
			return err
		}
	}
	if err := db.Delete(creditCardModel).Error; err != nil {
		log.Println("delete credit card err: ", err, ". card id is ", card.Id, ". user id is ", card.UserId)
		return err
	}
	log.Println("delete credit card success")
	return nil
}

//SelectCardById：根据id查询信用卡信息
func SelectCardById(cardId int, db *gorm.DB) (*CreditCard, error) {
	creditCard := &CreditCard{}
	if err := db.Model(&CreditCard{}).Where("id = ?", cardId).Find(&creditCard).Error; err != nil {
		log.Println("select credit card err: ", err, ". card id is ",cardId)
		return nil, err
	}
	log.Println("select credit card success")
	return creditCard, nil
}

//通过银行id和卡号查询信用卡
func SelectCardByBankAndCardId(bankId int64, userId int64, db *gorm.DB) (*CreditCard, error) {
	var card CreditCard
	if err := db.Model(&CreditCard{}).Where("bank_id = ? and user_id = ?", bankId, userId).First(&card).Error; err != nil {
		log.Println("[db] find card by bank and card id err: ", err)
		return nil, err
	}
	log.Println("[db] find card by bank and card id success")
	return &card, nil
}

//SelectCardByPage：分页查询信用卡信息    limit=-1 && offset=0查询所有
func SelectCardByPage(limit int, offset int, db *gorm.DB) ([]*CreditCard, error) {
	creditCards := make([]*CreditCard, 0, 0)
	if err := db.Model(&CreditCard{}).Limit(limit).Offset(offset).Find(&creditCards).Error; err != nil {
		log.Println("select credit card err: ", err, ". card num is ",len(creditCards))
		return nil, err
	}
	log.Println("select credit card success")
	return creditCards, nil
}

//SelectUserCardById：根据user id查询信用卡信息
func SelectUserCardById(userId int64, db *gorm.DB) ([]*CreditCard, error) {
	creditCard := make([]*CreditCard, 0, 0)
	if err := db.Model(&CreditCard{}).Where("user_id = ?", userId).Find(&creditCard).Error; err != nil {
		log.Println("select credit card err: ", err, ". user id is ",userId)
		return nil, err
	}
	log.Println("select credit card success")
	return creditCard, nil
}