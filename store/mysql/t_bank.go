package mysql

import (
	"credit_gin/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Bank struct {
	Id			int64 		`gorm:"column:id"`
	BankName	string		`gorm:"column:bank_name"`
	BankPic		string		`gorm:"column:bank_pic"`
	BankCardPic	string		`gorm:"column:bank_card_pic"`
	BankAddress	string		`gorm:"column:bank_address"`
	Desc		string		`gorm:"column:desc"`
	Province	string		`gorm:"column:province"`
	Town		string		`gorm:"column:town"`
	Country		string		`gorm:"column:country"`
	CreateTime	int64		`gorm:"column:create_time"`
	UpdateTime	int64		`gorm:"column:update_time"`
}

func (*Bank) TableName() string{
	return "bank"
}

//InsertBank: 新增银行
func InsertBank(bankReq model.BankReq, db *gorm.DB) error {
	bankModel := Bank{
		BankName:    bankReq.BankName,
		BankPic:	 bankReq.BankPic,
		BankCardPic: bankReq.BankCardPic,
		BankAddress: bankReq.BankAddress,
		Desc:		 bankReq.Desc,
		Province:    bankReq.Province,
		Town:        bankReq.Town,
		Country:     bankReq.Country,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}
	if err := db.Create(&bankModel).Error; err != nil {
		log.Println("insert bank info err: ", err)
		return err
	}
	log.Println("insert bank info success")
	return nil
}

//UpdateBank: 更新银行信息
func UpdateBank(bankReq model.BankReq, db * gorm.DB) error{
	bankModel := Bank{
		Id:			 bankReq.Id,
		BankName:    bankReq.BankName,
		BankPic:	 bankReq.BankPic,
		BankCardPic: bankReq.BankCardPic,
		BankAddress: bankReq.BankAddress,
		Desc:		 bankReq.Desc,
		Province:    bankReq.Province,
		Town:        bankReq.Town,
		Country:     bankReq.Country,
		UpdateTime:  time.Now().Unix(),
	}
	var bank Bank
	if err := db.Model(&Bank{}).Where("id = ?", bankReq.Id).First(&bank).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			bankModel.CreateTime = time.Now().Unix()
			log.Println("[db] create bank: ", bankModel)
			db.Create(&bankModel)
		}
	}else {
		log.Println("[db] update bank")
		db.Model(&Bank{}).Where("id = ?", bankReq.Id).Update(&bankModel)
	}
	return nil
}

//DeleteBank: 根据id删除银行信息
func DeleteBank(bankId int64, db *gorm.DB) error {
	bank := Bank{
		Id: bankId,
	}
	if err := db.Delete(&bank).Error; err != nil {
		log.Println("delete bank err: ", err)
		return err
	}
	return nil
}

//SelectBankInfoById: 根据id查询银行信息
func SelectBankInfoById(bankId int64, db *gorm.DB) (*Bank, error) {
	bank := &Bank{}
	if err := db.Where("id = ?", bankId).Find(&bank).Error; err != nil {
		log.Println("[db] select bank by id err: ", err)
		return nil, err
	}
	log.Println("[db] select bank by id success")
	return bank, nil
}

//SelectBankByPage: 分页查询所以银行信息 limit=-1 && offset = 0查询全部
func SelectBankInfoByPage(limit int, offset int) ([]*Bank, error) {
	banks := make([]*Bank, 0, 0)
	if err := db.Limit(limit).Offset(offset).Find(&banks).Error; err != nil {
		log.Println("select bank by page err: ", err)
		return banks, err
	}
	log.Println("select bank by page success")
	return banks, nil
}









