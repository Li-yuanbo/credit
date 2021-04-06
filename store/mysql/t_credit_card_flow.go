package mysql

type CreditCardFlow struct {
	Id 					int64 	`gorm:"column:id"`
	UserId				int64	`gorm:"column:user_id"`
	BankId				int64	`gorm:"column:bank_id"`
	CreditCardNum		string	`gorm:"column:credit_card_num"`
	OperateDesc			string	`gorm:"column:operate_desc"`
	OperateFlow			string	`gorm:"column:operate_flow"`
	OperateType			int64	`gorm:"column:operate_type"`  //0-存款 1-取款 2-转账 3-被转账
	Money				float64	`gorm:"column:money"`
	MoneyType			int64	`gorm:"column:money_type"`    //0-收入 1-支出
	AfterBalance		float64	`gorm:"column:after_balance"`
	CreateTime			int64	`gorm:"column:create_time"`
	UpdateTime			int64	`gorm:"column:update_time"`
}

func (*CreditCardFlow) TableName() string {
	return "credit_card_flow"
}