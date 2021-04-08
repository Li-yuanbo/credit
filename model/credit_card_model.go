package model

type CreditCardFlowModel struct {
	Id 					int64 	`json:"id"`
	UserId				int64	`json:"user_id"`
	BankId				int64	`json:"bank_id"`
	CreditCardNum		string	`json:"credit_card_num"`
	OperateDesc			string	`json:"operate_desc"`
	OperateFlow			string	`json:"operate_flow"`
	OperateType			int64	`json:"operate_type"`   //0-存款 1-取款 2-转账 3-被转账
	Money				float64	`json:"money"`
	MoneyType			int64	`json:"money_type"`     //0-收入 1-支出
	AfterBalance		float64	`json:"after_balance"`
	CreateTime			int64	`json:"create_time"`
	UpdateTime			int64	`json:"update_time"`
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

type OperateBalanceReq struct {
	UserId			int64	`json:"user_id"`
	BankId 			int64	`json:"bank_id"`
	CreditCardNum	string	`json:"credit_card_num"`
	OperateDesc		string	`json:"operate_desc"`
	OperateType		int64	`json:"operate_type"` //0-存款 1-取款 2-转账 3-被转账
	Money			float64	`json:"money"`
}