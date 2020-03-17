package vo


type CoinOutVo struct {
	Id string `form:"coinout_id"`
	Address string `form:"address"`
	Amount float64 `form:"amount"`
	UserId string `form:"user_id"`
	Index int `form:"index"`
	Status int `form:"status"`
	Account string `form:"account"`
}