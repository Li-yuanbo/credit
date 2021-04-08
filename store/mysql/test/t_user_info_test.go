package test

import (
	"credit_gin/model"
	"credit_gin/store/mysql"
	"fmt"
	"log"
	"testing"
)

func TestInsertUserInfo(t *testing.T) {
	req := model.RegisterUserReq{
		UserName: "李四",
		Password: "lisi",
		Phone:    "13299162870",
	}
	_, err := mysql.InsertUserInfo(req, mysql.WriteDB())
	if err != nil {
		log.Println("insert user err: ", err)
		return
	}
	log.Println("insert user success")
}

func TestUserInfo(t *testing.T) {
	req := model.DeleteUserReq{
		Id: 5,
	}
	mysql.DeleteUserInfoById(req, mysql.WriteDB())
}

func TestSelectUsers(t *testing.T){
	users, err := mysql.SelectUserInfos(-1, 0, mysql.WriteDB())
	if err != nil {
		log.Println("select users err: ", err)
		return
	}
	log.Println("select users success, user num: ", len(users))
	for _, user := range users {
		fmt.Println(user.Name)
	}
}