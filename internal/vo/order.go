package vo

type OrderVo struct {
	ID      string   `form:"order_id"`
	AskId 	string `form:"ask_id"`
	Amount  float64 `form:"amount"`
	Price	float64 `form:"price"`
	UserId  string `form:"user_id"`
	UserName  string `form:"user_name"`
	Index   int  `form:"index"`
	Length int `form:"length"`
	Status int `form:"status"`
	UnitPrice float64 `form:"unit_price"`
}
