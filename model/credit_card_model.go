package model

type CreditCard struct {
	UserId			int		`json:"user_id"`
	CreditCardNum	string	`json:"credit_card_num"`
	Balance			float64	`json:"balance"`
	BankName		string	`json:"bank_name"`
	BankCardPic		string	`json:"bank_card_pic"`
}

type CreditCardReq struct {
	Id				int64`json:"id"`
	BankId			int64		`json:"bank_id"`
	UserId			int64		`json:"user_id"`
	CreditCardNum	string	`json:"credit_card_num"`
	Balance			float64	`json:"balance"`
	CreateTime		int64	`json:"create_time"`
	UpdateTime		int64	`json:"update_time"`
}

type CreditCardFlowReq struct {
	Id					int		`json:"id"`
	CreditCardId		int		`json:"credit_card_id"`
	OperateUserId		int 	`json:"operate_user_id"`
	OperateFlow			string	`json:"operate_flow"`
	Desc				string	`json:"desc"`
	OperateType			int		`json:"operate_type"` //0-存取款 1-转账
	OtherUserId			int		`json:"other_user_id"` //转账收款人id
	OtherCreditCardId	int	`json:"other_credit_card_id"` //收款人信用卡号
}
