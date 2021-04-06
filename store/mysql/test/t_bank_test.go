package test

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	"fmt"
	"log"
	"testing"
)

func TestInsertBank(t *testing.T) {
	bankReq := model.BankReq{
		BankName:    "中国银行",
		BankPic:	 "http://bank-pic.cn",
		BankCardPic: "http://bank_card_pic.cn",
		BankAddress: "惠泉路",
		Desc:		 "中国银行简介",
		Province:    "陕西省",
		Town:        "榆林市",
		Country:     "神木县",
	}
	err := mysql.InsertBank(bankReq, mysql.WriteDB())
	if err != nil {
		log.Println("insert bank info err: ", err)
		return
	}
	log.Println("insert bank info success")
}

func TestUpdateBank(t *testing.T){
	bankReq := model.BankReq{
		Id:          12,
		BankName:    "工商银行",
		Town:        "西安",
		Country:     "莲湖区",
	}
	_ = mysql.UpdateBank(bankReq, mysql.WriteDB())
}

func TestDeleteBank(t *testing.T) {
	mysql.DeleteBank(13, mysql.WriteDB())
}

func TestSelectBankInfoById(t *testing.T) {
	bank, _ := mysql.SelectBankInfoById(14, mysql.WriteDB())
	fmt.Println(bank)
}

func TestSelectBankInfoByPage(t *testing.T) {
	banks, _ := mysql.SelectBankInfoByPage(-1, 0)
	for _, bank := range banks {
		fmt.Println(bank)
	}
}