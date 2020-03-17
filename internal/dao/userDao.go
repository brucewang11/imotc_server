package dao

import "github.com/brucewang11/frame/internal/model"

type UserDao struct {
	BaseDao
}


func (t *UserDao) EditBalance(user *model.User) error{
	return  t.Conn.Exec("update tb_user set usdt_balance = ? where id = ?",user.UsdtBalance,user.ID).Error

}