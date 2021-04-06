package handler

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	. "credit_gin/utils"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//用户绑定信用卡
//银行id: bank_id  信用卡号: credit_card_num
func BindCard(c * gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	var req model.CreditCardReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
	}
	log.Println("bind card request:", req)
	//参数校验
	if req.BankId < 1 || req.CreditCardNum == "" {
		log.Println("request param err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10004, ERR_CODE[10004])
		return
	}
	//通过session获取UserId
	session := sessions.Default(c)
	if session.Get("user") == nil {
		log.Println("user already login err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 20003, ERR_CODE[20003])
		return
	}
	req.UserId = session.Get("user").(int64)

	tx := mysql.WriteDB().Begin()
	//保证每个用户在一个银行只有一张信用卡
	cards, err := mysql.SelectUserCardById(req.UserId, tx)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	flag := false
	for _, card := range cards {
		if card.BankId == req.BankId {
			flag = true
			break
		}
	}
	if flag {
		log.Println("user already has credit card in this bank")
		ErrResponse(c, http.StatusInternalServerError, 30001, ERR_CODE[30001])
		return
	}
	//新增信用卡
	card, err := mysql.AddCreditCard(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	res := model.CreditCardReq{
		Id:            card.Id,
		BankId:        card.BankId,
		UserId:        card.UserId,
		CreditCardNum: card.CreditCardNum,
		Balance:       card.Balance,
		CreateTime:	   card.CreateTime,
		UpdateTime:    card.UpdateTime,
	}
	SucResponse(c, res)
}

//存取款/转账
//操作描述：desc 操作类型：operate_type 收款人id：other_user_id 收款人卡号：other_credit_card_id money: 金额
func OperateCard(c *gin.Context) {
	//1、查询收款人是否存在，收款人卡号是否正确

	//2、检查打款人信用卡余额是否足够

	//3、补全操作信息

	//4、更新双方信用卡余额

	//5、新增双方的操作信息
}