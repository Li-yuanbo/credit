package model


type BankModel struct {
	Id			int64	`json:"id"`
	BankName	string	`json:"bank_name"`
	BankPic		string	`json:"bank_pic"`
	BankCardPic string	`json:"bank_card_pic"`
	BankAddress	string	`json:"bank_address"`	//详细地址
	Desc		string	`json:"desc"`			//银行简介
	Province	string	`json:"province"`		//省
	Town		string	`json:"town"`			//市
	Country		string	`json:"country"`		//县
	CreateTime	int64	`json:"create_time"`
	UpdateTime	int64	`json:"update_time"`
}

type BankReq struct {
	Id			int64	`json:"id"`
	BankName	string	`json:"bank_name"`
	BankPic		string	`json:"bank_pic"`
	BankCardPic string	`json:"bank_card_pic"`
	BankAddress	string	`json:"bank_address"`	//详细地址
	Desc		string	`json:"desc"`			//银行简介
	Province	string	`json:"province"`		//省
	Town		string	`json:"town"`			//市
	Country		string	`json:"country"`		//县
}

type PublishNewsReq struct {
	BankId		int64	`json:"bank_id"`
	NewsTitle	string	`json:"news_title"`
	NewsContent	string	`json:"news_content"`
}