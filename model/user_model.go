package model

type UserDetailModel struct {
	Id 			int64 		`json:"id"`
	UserId		int64 		`json:"user_id"`
	Gender		int64		`json:"gender"`    //0-男 1-女  2223333333333333333333333333333
	Birthday	string		`json:"birthday"`
	UserPic		string		`json:"user_pic"`
	UserDesc	string		`json:"user_desc"`
	CreateTime	int64		`json:"create_time"`
	UpdateTime	int64		`json:"update_time"`
}

type UserInfoReq struct{
	Id 			int64	`json:"id"`
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"password"`
	Phone 		string	`json:"phone"`
}

type UserDetailReq struct {
	Id			int64		`json:"id"`
	UserId		int64		`json:"user_id"`
	Gender		int64		`json:"gender"`		//0-男 1-女
	Birthday	string		`json:"birthday"`
	UserPic		string		`json:"user_pic"`
	UserDesc	string		`json:"user_desc"`
	CreateTime	int64		`json:"create_time"`
	UpdateTime	int64		`json:"update_time"`
}

type UserResp struct {
	UserId		int64			`json:"user_id"`
	UserName 	string			`json:"user_name"`
	Phone 		string			`json:"phone"`
	Gender		int64			`json:"gender"`		//0-男 1-女
	Birthday	string			`json:"birthday"`
	UserPic		string			`json:"user_pic"`
	UserDesc	string			`json:"user_desc"`
	UserCards	[]*CreditCard	`json:"user_cards"`
}