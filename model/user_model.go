package model

type UserInfoModel struct{
	Id			int64	`json:"id"`
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"password"`
	Phone 		string	`json:"phone"`
	IdCard		string	`json:"id_card"`
	Level		int64	`json:"level"`
	Gender		int64	`json:"gender"`
	Birthday	string	`json:"birthday"`
	Desc		string	`json:"desc"`
	CreateTime	int64	`json:"create_time"`
	UpdateTime	int64	`json:"update_time"`
}

type DeleteUserReq struct {
	Id		int64	`json:"id"`
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

type GetUserInfoReq struct {
	Id	int64	`json:"id"`
}

type UserCard struct {
	BankId			interface{}	`json:"bank_id"`
	BankName		interface{}	`json:"bank_name"`
	CreditCardNum	interface{}	`json:"credit_card_num"`
	Balance			interface{}	`json:"balance"`
}

type GetUserInfoResp struct {
	User	UserInfoModel `json:"user"`
	Card	[]*UserCard	  `json:"card"`
}

type GetUserInfosReq struct {
	Limit	int64	`json:"limit"`
	Offset	int64	`json:"offset"`
	CurPage	int64	`json:"cur_page"`
}

type GetUserInfosResp struct {
	TotalPage	int64				`json:"total_page"`
	CurPage		int64				`json:"cur_page"`
	Users		[]*UserInfoModel	`json:"users"`
}

type LoginUserReq struct{
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"password"`
}

type ExitLoginUserReq struct {
	Id		int64	`json:"id"`
}