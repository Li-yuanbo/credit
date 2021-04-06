package test

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	"fmt"
	"testing"
)

func TestAddCreditCardFlow(t *testing.T) {
	req := model.CreditCardFlowReq{
		CreditCardId:      1,
		OperateUserId:     1,
		OperateFlow:       "李四存钱1000",
		Desc:              "存钱",
		OperateType:       0,
	}
	mysql.AddCreditCardFlow(req, mysql.WriteDB())
}

func TestSelectFlowByUserId(t *testing.T) {
	res, _ := mysql.SelectFlowByUserId(1, mysql.WriteDB())
	for _, r := range res {
		fmt.Println(r)
	}
}

func TestSelectFlowByCardId(t *testing.T) {
	res, _ := mysql.SelectFlowByCardId(1, mysql.WriteDB())
	for _, r := range res {
		fmt.Println(r)
	}
}