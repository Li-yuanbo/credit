package main

import (
	. "credit_gin/handler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func handler(r *gin.Engine){
	//user_info
	r.POST("/user", RegisterUser)  //注册用户
	r.DELETE("/user", DeleteUser)  //注销用户
	r.GET("/user/login", LoginUser) //用户登录
	r.GET("/user", GetUserInfo)    //根据id查询用户的信息
	r.PUT("/user/detail", UpdateUserDetail) //修改用户个人信息

	//bank
	r.POST("/bank", UpdateBank) //新增/修改银行
	r.DELETE("/bank", DeleteBank) //删除银行

	//credit_card
	r.POST("/card", BindCard) //用户绑定信用卡
	r.PUT("/card", OperateCard) //存取款， 转账
}

func main(){
	//默认路由
	r := gin.Default()
	store := cookie.NewStore([]byte("user"))
	store.Options(sessions.Options{
		MaxAge: 24 * 60 * 60,
		Path:   "/",
	})
	//在路由上加入session中间件
	r.Use(sessions.Sessions("mysession", store))
	//加载handler
	handler(r)
	//启动项目
	r.Run(":8080")
}