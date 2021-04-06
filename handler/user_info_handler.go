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
//注册用户
//用户名：user_name  密码：password  手机号：phone
func RegisterUser(c *gin.Context) {
	//从request body读取数据
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.UserInfoReq
	//反序列化
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
	}
	//参数校验
	if req.UserName == "" || req.Password == "" || req.Phone == "" {
		log.Println("request param err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10004, ERR_CODE[10004])
		return
	}
	req.Password = MD5(req.Password)
	db := mysql.WriteDB()
	userId, err := mysql.InsertUserInfo(req, db)
	req.Id = userId
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
	}
	//写session
	session := sessions.Default(c)
	if session.Get("user") == nil {
		session.Set("user", userId)
		session.Save()
	}
	SucResponse(c, req)
}

//登录
//用户名: user_name 	密码: password
func LoginUser(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.UserInfoReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	db := mysql.WriteDB()
	user, err := mysql.SelectUserByUserName(req.UserName, db)
	if user.Password != MD5(req.Password) {
		log.Println("user password err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 20001, ERR_CODE[20001])
		return
	}
	//登录成功写session
	session := sessions.Default(c)
	//判断用户是否已经登录
	log.Println("session: ", session.Get("user"))
	if session.Get("user") != nil && session.Get("user").(int64) == user.Id {
		log.Println("user already login err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 20003, ERR_CODE[20003])
		return
	}
	session.Set("user", user.Id)
	session.Save()
	log.Println("user login success: ", user)
	SucResponse(c, user)
}

//根据id查询用户全部信息
func GetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user") == nil {
		log.Println("get session user err")
		ErrResponse(c, http.StatusInternalServerError,20002, ERR_CODE[20002])
	}
	userId := session.Get("user")
	db := mysql.WriteDB()
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()
	//获取用户基础信息
	userInfo, err := mysql.SelectUserInfoById(userId.(int64), tx)
	if err != nil {
		db.Rollback()
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	res := model.UserResp{
		UserId:     userInfo.Id,
		UserName:   userInfo.Name,
		Phone:      userInfo.Phone,
	}
	//获取用户详细信息
	userDetail, err := mysql.SelectUserDetailById(userId.(int64), tx)
	if err != nil {
		db.Rollback()
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	res.Gender = userDetail.Gender
	res.Birthday = userDetail.Birthday
	res.UserDesc = userDetail.UserDesc
	res.UserPic = userDetail.UserPic

	//TODO: 查询用户的信用卡和存款信息

	SucResponse(c, res)
}

//修改个人信息
//性别: gender 生日: birthday 头像: user_pic 个人简介: user_desc
func UpdateUserDetail(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
		return
	}
	var req model.UserDetailReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("update user detail: ", req)
	//从session获取当前用户id
	session := sessions.Default(c)
	if session.Get("user") == nil {
		log.Println("user not login err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 20004, ERR_CODE[20004])
		return
	}
	userId := session.Get("user").(int64)
	req.UserId = userId
	err = mysql.UpdateUserDetail(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	detail, err := mysql.SelectUserDetailById(req.UserId, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	SucResponse(c, detail)
}