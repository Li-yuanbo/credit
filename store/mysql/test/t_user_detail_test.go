package test

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	"fmt"
	"testing"
	"time"
)

func TestAddUserDetail(t *testing.T) {
	req := model.UserDetailReq{
		UserId:     12,
		Gender:     0,
		Birthday:   "2021-10-19",
		UserPic:    "http://user-pic.cn",
		UserDesc:   "个人简介3",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	mysql.AddUserDetail(req, mysql.WriteDB())
}

func TestUpdateUserDetail(t *testing.T) {
	req := model.UserDetailReq{
		UserId:		2,
		Birthday:   "2021-03-03",
		UserPic:    "http://user-pic.com",
		UserDesc:   "个人简介",
	}
	mysql.UpdateUserDetail(req, mysql.WriteDB())
}

func TestSelectUserDetailById(t *testing.T) {
	detail, _ := mysql.SelectUserDetailById(1, mysql.WriteDB())
	fmt.Println(detail)
}

func TestSelectUserDetailByPage(t *testing.T) {
	details, _ := mysql.SelectUserDetailByPage(-1, 0, mysql.WriteDB())
	for _, detail := range details {
		fmt.Println(detail)
	}
}