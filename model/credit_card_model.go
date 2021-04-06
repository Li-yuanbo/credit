package model

type CreditCard struct {
	UserId			int		`json:"user_id"`
	CreditCardNum	string	`json:"credit_card_num"`
	Balance			float64	`json:"balance"`
	BankName		string	`json:"bank_name"`
	BankCardPic		string	`json:"bank_card_pic"`
}

type CreditCardReq struct {
	Id				int64		`json:"id"`
	BankId			int64		`json:"bank_id"`
	UserId			int64		`json:"user_id"`
	CreditCardNum	string		`json:"credit_card_num"`
	Balance			float64		`json:"balance"`
	CreateTime		int64		`json:"create_time"`
	UpdateTime		int64		`json:"update_time"`
}