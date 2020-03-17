package vo

type AskVo struct{
	ID string  `form:"ask_id"`
	UserName string `form:"user_name"`
	Amount float64 `form:"amount"`
	MinLimit float64 `form:"min_limit"`
	MaxLimit float64 `form:"max_limit"`
	Price float64 `form:"price"`
	BankCard string `form:"bank_card"`
	BankAddress string `form:"bank_address"`
	BankName string  `form:"bank_name"`

}