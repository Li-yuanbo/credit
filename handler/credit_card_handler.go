package handler

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	. "credit_gin/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

//解绑信用卡
//id
func UnBindCard(c *gin.Context) {
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
	if req.Id < 1 {
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
	err = mysql.DeleteCreditCard(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	SucResponse(c, nil)
}

//存取款: 数据库只存个人的操作流
//银行id：bank_id  信用卡号：credit_card_num  操作备注：operate_desc 操作类型：operate_type  金额：money
func OperateBalance(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	var req model.OperateBalanceReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	//判断用户登录
	session := sessions.Default(c)
	if session.Get("user") == nil{
		log.Println("get user session err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 20004, ERR_CODE[20004])
		return
	}
	req.UserId = session.Get("user").(int64)
	tx := mysql.WriteDB().Begin()
	defer tx.RollbackUnlessCommitted()

	//根据card_num获取卡余额, 取款则判断余额是否足够
	card, err := mysql.SelectCardByBankAndCardId(req.BankId, req.CreditCardNum, tx)
	if gorm.IsRecordNotFoundError(err) {
		ErrResponse(c, http.StatusInternalServerError, 30002, ERR_CODE[30002])
		return
	}else if err != nil{
		ErrResponse(c, http.StatusInternalServerError, 30002, ERR_CODE[30002])
		return
	}
	if req.OperateType == 1 && card.Balance < req.Money {
		log.Println("card balance too less")
		ErrResponse(c, http.StatusInternalServerError, 30003, ERR_CODE[30003])
		return
	}
	operateModel := model.CreditCardFlowModel{
		UserId:        req.UserId,
		BankId:        req.BankId,
		CreditCardNum: req.CreditCardNum,
		OperateDesc:   req.OperateDesc,
		OperateType:   req.OperateType,
		Money:         req.Money,
	}
	//修改信用卡余额
	var updateMode model.CreditCardReq
	//流水类型 （0-支出、1-收入）
	if req.OperateType == 1 {
		operateModel.MoneyType = 0
		operateModel.AfterBalance = card.Balance - req.Money
		updateMode.Balance = card.Balance - req.Money
	}else {
		operateModel.MoneyType = 0
		operateModel.AfterBalance = card.Balance + req.Money
		updateMode.Balance = card.Balance + req.Money
	}
	//更新信用卡余额
	updateMode.CreditCardNum = req.CreditCardNum
	err = mysql.UpdateCreditCard(updateMode, tx)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	//组成操作流水
	operateModel.OperateFlow = fmt.Sprintf("卡号为%s的信用卡进行%s操作，操作金额为%.2f，%s后，卡内余额为%.2f",
		req.CreditCardNum, OPERATE_TYPE[req.OperateType], req.Money, OPERATE_TYPE[req.OperateType], operateModel.AfterBalance)
	cardFlow, err := mysql.InsertOperateFlow(operateModel, tx)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	tx.Commit()
	res := model.CreditCardFlowModel{
		Id:            cardFlow.Id,
		UserId:        cardFlow.UserId,
		BankId:        cardFlow.BankId,
		CreditCardNum: cardFlow.CreditCardNum,
		OperateDesc:   cardFlow.OperateDesc,
		OperateFlow:   cardFlow.OperateFlow,
		OperateType:   cardFlow.OperateType,
		Money:         cardFlow.Money,
		MoneyType:     cardFlow.MoneyType,
		AfterBalance:  cardFlow.AfterBalance,
		CreateTime:    cardFlow.CreateTime,
		UpdateTime:    cardFlow.UpdateTime,
	}
	SucResponse(c, res)
}