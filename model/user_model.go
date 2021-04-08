package model

type UserInfoModel struct{
	Id			int64	`json:"id"`
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"password"`
	Phone 		string	`json:"phone"`
	IdCard		string	`json:"id_card"`
	Gender		int64	`json:"gender"`
	Birthday	string	`json:"birthday"`
	Desc		string	`json:"desc"`
	CreateTime	int64	`json:"create_time"`
	UpdateTime	int64	`json:"update_time"`
}

type RegisterUserReq struct{
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"password"`
	Phone 		string	`json:"phone"`
	IdCard		string	`json:"id_card"`
	Gender		int64	`json:"gender"`
	Birthday	string	`json:"birthday"`
	Desc		string	`json:"desc"`
}

type LoginUserReq struct{
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"password"`
}