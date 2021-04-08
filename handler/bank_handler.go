package handler

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	. "credit_gin/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//新增或更新银行信息  不存在则新增，存在则修改
//id	bank_name	bank_pic	bank_card_pic	desc	province	town	country
func UpdateBank(c *gin.Context) {
	//读取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.BankReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("add bank request:", req)
	//db操作
	tx := mysql.WriteDB().Begin()
	defer tx.RollbackUnlessCommitted()
	err = mysql.UpdateBank(req, tx)
	if err != nil {
		tx.Rollback()
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	bank, err := mysql.SelectBankInfoById(req.Id, tx)
	if err != nil {
		tx.Rollback()
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	tx.Commit()
	res := model.BankModel{
		Id:          bank.Id,
		BankName:    bank.BankName,
		BankPic:     bank.BankPic,
		BankCardPic: bank.BankCardPic,
		BankAddress: bank.BankAddress,
		Desc:        bank.Desc,
		Province:    bank.Province,
		Town:        bank.Town,
		Country:     bank.Country,
		CreateTime:  bank.CreateTime,
		UpdateTime:  bank.UpdateTime,
	}
	SucResponse(c, res)
}

//删除银行
//银行id：id
func DeleteBank(c *gin.Context) {
	//获取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var bank model.BankModel
	err = json.Unmarshal(body, &bank)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("delete bank id:", bank.Id)
	//参数校验
	if bank.Id < 1 {
		log.Println("bank id err: ", bank.Id)
		ErrResponse(c, http.StatusInternalServerError, 10004, ERR_CODE[10004])
		return
	}
	err = mysql.DeleteBank(bank.Id, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	SucResponse(c, nil)
}

//获取银行信息
//bank_name
func GetBankInfo(c *gin.Context) {
	//获取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.GetBankInfoReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("request param: ", req)
	bank, err := mysql.SelectBankInfoByName(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	res := model.BankModel{
		Id:          bank.Id,
		BankName:    bank.BankName,
		BankPic:     bank.BankPic,
		BankCardPic: bank.BankCardPic,
		BankAddress: bank.BankAddress,
		Desc:        bank.Desc,
		Province:    bank.Province,
		Town:        bank.Town,
		Country:     bank.Country,
		CreateTime:  bank.CreateTime,
		UpdateTime:  bank.UpdateTime,
	}
	SucResponse(c, res)
}

//分页获取全部银行信息
//page limit
func GetBankInfos(c *gin.Context) {
	//获取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.GetBankInfosReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("request param: ", req)

	banks, err := mysql.SelectBankInfoByPage(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	var banksModel []*model.BankModel
	for _, b := range banks {
		bModel := &model.BankModel{
			Id:          b.Id,
			BankName:    b.BankName,
			BankPic:     b.BankPic,
			BankCardPic: b.BankCardPic,
			BankAddress: b.BankAddress,
			Desc:        b.Desc,
			Province:    b.Province,
			Town:        b.Town,
			Country:     b.Country,
			CreateTime:  b.CreateTime,
			UpdateTime:  b.UpdateTime,
		}
		banksModel = append(banksModel, bModel)
	}
	res := model.GetBankInfosResp{
		Banks:   banksModel,
	}
	//获取当前页数和总页数
	//1、获取银行总数
	totalNum, err := mysql.SelectBankTotalCount(mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	//总页数
	if totalNum % req.Limit == 0 {
		res.TotalPage = totalNum / req.Limit
	}else {
		res.TotalPage = totalNum / req.Limit + 1
	}
	res.CurPage = req.CurPage + 1
	SucResponse(c, res)
}

//银行发布文章
//bank_id	news_title	news_content
func PublishNews(c *gin.Context) {
	//获取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.PublishNewsReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	news, err := mysql.InsertBankNews(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	SucResponse(c, news)
}












