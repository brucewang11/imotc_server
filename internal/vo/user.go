package vo


type UserVo struct {
	OldPwd string `form:"old_pwd"`
	Account string `form:"account"`
	Pwd     string `form:"pwd"`
	Code    string `form:"code"`
	Index   int     `form:"index"`
	CodeType string  `form:"code_type"`
	UserId  string
}


type LoadRes struct {
	Account string `json:"account"`
	ID      string `json:"id"`
	Token   string `json:"token"`
}
