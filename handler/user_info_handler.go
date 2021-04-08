package handler

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	. "credit_gin/utils"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)
//注册用户
//用户名：user_name  密码：password  手机号：phone 身份证：id_card 性别：gender 生日：birthday 个人简介：desc
func RegisterUser(c *gin.Context) {
	//从request body读取数据
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.RegisterUserReq
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
	var res model.UserInfoModel
	req.Password = MD5(req.Password)
	db := mysql.WriteDB()
	userId, err := mysql.InsertUserInfo(req, db)
	res.Id = userId
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	//写session
	session := sessions.Default(c)
	if session.Get("user") == nil {
		session.Set("user", userId)
		session.Save()
	}
	SucResponse(c, req)
}

//判断用户是否为admin权限
func CheckAdmin(id int64) (bool, error) {
	flag := false
	user, err := mysql.SelectUserInfoById(id, mysql.WriteDB())
	if err != nil {
		return flag, err
	}
	if user.Level == 1 {
		flag = true
	}
	return flag, nil
}

//删除用户
//用户id： id
func DeleteUser(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.DeleteUserReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	//判断用户权限是否为admin
	flag, err := CheckAdmin(req.Id)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	if !flag {
		ErrResponse(c, http.StatusInternalServerError, 20008, ERR_CODE[20008])
		return
	}
	//删除用户
	err = mysql.DeleteUserInfoById(req, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	//删除session
	session := sessions.Default(c)
	if session.Get("user").(int64) == req.Id {
		session.Delete("user")
		session.Save()
	}
	SucResponse(c, nil)
}

//根据id获取用户信息
func GetUserInfo(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.GetUserInfoReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	user, err := mysql.SelectUserInfoById(req.Id, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	res := model.UserInfoModel{
		Id:         user.Id,
		UserName:   user.Name,
		Password:   user.Password,
		Phone:      user.Phone,
		IdCard:     user.IdCard,
		Level:      user.Level,
		Gender:     user.Gender,
		Birthday:   user.Birthday,
		Desc:       user.Desc,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
	SucResponse(c, res)
}

//分页获取用户信息
//limit offset cur_page
func GetUserInfos(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.GetUserInfosReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	users, err := mysql.SelectUserInfos(req.Limit, req.Offset, mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	//获取用户总数
	total, err := mysql.GetUserTotalCount(mysql.WriteDB())
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, 10003, ERR_CODE[10003])
		return
	}
	var res model.GetUserInfosResp
	if total % req.Limit == 0 {
		res.TotalPage = total / req.Limit
	}else {
		res.TotalPage = total / req.Limit + 1
	}
	res.CurPage = req.CurPage
	for _, user := range users {
		u := &model.UserInfoModel{
			Id:         user.Id,
			UserName:   user.Name,
			Password:   user.Password,
			Phone:      user.Phone,
			IdCard:     user.IdCard,
			Level:      user.Level,
			Gender:     user.Gender,
			Birthday:   user.Birthday,
			Desc:       user.Desc,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		}
		res.Users = append(res.Users, u)
	}
	SucResponse(c, res)
}

//登录
//用户名: user_name 	密码: password 权限：level
func LoginUser(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.LoginUserReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	db := mysql.WriteDB()
	user, err := mysql.SelectUserByUserName(req.UserName, db)
	if gorm.IsRecordNotFoundError(err) {
		log.Println("username err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 20007, ERR_CODE[20007])
		return
	}
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
	res := model.UserInfoModel{
		Id:         user.Id,
		UserName:   user.Name,
		Password:   user.Password,
		Phone:      user.Phone,
		IdCard:     user.IdCard,
		Level:		user.Level,
		Gender:     user.Gender,
		Birthday:   user.Birthday,
		Desc:       user.Desc,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
	SucResponse(c, res)
}

func ExitLoginUser(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError,10001, ERR_CODE[10001])
	}
	var req model.ExitLoginUserReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	session := sessions.Default(c)
	if session.Get("user").(int64) == req.Id {
		session.Delete("user")
		session.Save()
	}else {
		ErrResponse(c, http.StatusInternalServerError, 20004, ERR_CODE[20004])
		return
	}
	SucResponse(c, nil)
}