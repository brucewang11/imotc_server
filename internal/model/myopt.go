package model

type SysUser struct {
	ID       string    `gorm:"column:id"`
	Account     string `gorm:"column:account"`
	Pwd string `gorm:"column:pwd"`
}

func (sysUser SysUser) TableName() string {
	return "tb_sys_user"
}

type User struct {
	ID       string    `gorm:"column:id"`
	Account     string `gorm:"column:account"`
	Pwd string `gorm:"column:pwd"`
	UsdtBalance float64 `gorm:"column:usdt_balance"`
	CreateTime int64  `gorm:"column:create_time"`
}
func (user User) TableName() string {
	return "tb_user"
}

type Ask struct {
	ID       string    `gorm:"column:id"`
	UserName string       `gorm:"column:user_name"`
	Amount  float64    `gorm:"column:amount"`
	MinLimit float64   `gorm:"column:min_limit"`
	MaxLimit float64   `gorm:"column:max_limit"`
	Price   float64     `gorm:"column:price"`
	BankCard string    `gorm:"column:bank_card"`
	CoinType int     `gorm:"column:coin_type"`
	PayType  string   `gorm:"column:pay_type"`
	BankName string  `gorm:"column:bank_name"`
	BankAddress string  `gorm:"column:bank_address"`
}

func (ask Ask) TableName() string {
	return "tb_ask"
}

type CoinOut struct {
	Id string `gorm:"column:id"`
	Address string `gorm:"column:address"`
	CoinType int `gorm:"column:coin_type"`
	Amount float64 `gorm:"column:amount"`
	CreateTime int64 `gorm:"column:create_time"`
	UpdateTime int64 `gorm:"column:update_time"`
	UserId  string `gorm:"column:user_id"`
	Status int `gorm:"status"`
	Account string `gorm:"account"`
}
func (coinOut CoinOut) TableName() string {
	return "tb_coin_out"
}


type Order struct {
	ID       string    `gorm:"column:id"`
	UserId   string     `gorm:"column:user_id"`
	UserName string       `gorm:"column:user_name"`
	Amount  float64    `gorm:"column:amount"`
	AskId   string     `gorm:"column:ask_id"`
	Price   float64     `gorm:"column:price"`
	Status int    `gorm:"column:status"`
	CoinType int     `gorm:"column:coin_type"`
	VerifyCode string  `gorm:"column:verify_code"`
	CreateTime int64 `gorm:"column:create_time"`
	UnitPrice   float64 `gorm:"column:unit_price"`
	Fee        float64  `gorm:"column:fee"`

}

func (order Order) TableName() string {
	return "tb_order"
}







