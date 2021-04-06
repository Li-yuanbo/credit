package test

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	"fmt"
	"testing"
)

func TestAddCreditCard(t *testing.T) {
	req := model.CreditCardReq{
		BankId:        2,
		UserId:        2,
		CreditCardNum: "4396 7923 3333",
	}
	mysql.AddCreditCard(req, mysql.WriteDB())
}

func TestUpdateCreditCard(t *testing.T) {
	req := model.CreditCardReq{
		Id:			   1,
		Balance:       98.94,
	}
	_ = mysql.UpdateCreditCard(req, mysql.WriteDB())
}

func TestDeleteCreditCard(t *testing.T) {
	_ = mysql.DeleteCreditCard(3, mysql.WriteDB())
}

func TestSelectCardById(t *testing.T) {
	creditCard, _ := mysql.SelectCardById(1, mysql.WriteDB())
	fmt.Println(creditCard)
}

func TestSelectCardByPage(t *testing.T) {
	creditCards, _ := mysql.SelectCardByPage(-1, 0, mysql.WriteDB())
	for _, card := range creditCards {
		fmt.Println(card)
	}
}