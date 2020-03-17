package service

import (
	"fmt"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/model"
	"github.com/brucewang11/frame/internal/vo"
	uuid "github.com/satori/go.uuid"
	"time"
)

func CreateCoinOut(vo *vo.CoinOutVo)*CodeModel{
	if vo.Amount <=0 ||(vo.Address==""){
		return &CodeModel{Code:4001}
	}
	tx := dao.GetTransaction()

	coinOutDao := &dao.CoinOutDao{BaseDao:dao.NewBaseDao(tx)}
	userDao := &dao.UserDao{BaseDao:dao.NewBaseDao(tx)}
	user := model.User{}
	_,err := userDao.First(&user,model.User{ID:vo.UserId})
	if err!=nil {
		fmt.Println("create",err)
		tx.Rollback()
		return &CodeModel{Code:4002}
	}
	if user.UsdtBalance <vo.Amount {
		tx.Rollback()
		return &CodeModel{Code:4003}
	}
	err = coinOutDao.Create(&model.CoinOut{
		Id:         uuid.NewV4().String(),
		Address:    vo.Address,
		CoinType:   1,
		Amount:     vo.Amount,
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: 0,
		UserId:vo.UserId,
		Status:1,
		Account:vo.Account,
	})
	if err!=nil {
		fmt.Println("create",err)
		tx.Rollback()
		return &CodeModel{Code:4002}
	}
	err = tx.Exec("update tb_user set usdt_balance = ? where id = ?",user.UsdtBalance - vo.Amount,user.ID).Error
	if err!=nil {
		fmt.Println("create",err)
		tx.Rollback()
		return &CodeModel{Code:4002}
	}
	tx.Commit()
	return &CodeModel{Code:0}

}

func CoinOutList(outVo *vo.CoinOutVo)*CodeModel{
	coinOutDao := &dao.CoinOutDao{BaseDao:dao.NewBaseDao(nil)}
	list := []model.CoinOut{}
	count,err := coinOutDao.Count(model.CoinOut{UserId:outVo.UserId},model.CoinOut{})
	if err !=nil {
		fmt.Println("CoinOutList",err)
		return &CodeModel{Code:4004}
	}
	err = coinOutDao.ListByPage(model.CoinOut{UserId:outVo.UserId},&list,&dao.Page{PageIndex:outVo.Index,PageSize:10},"create_time desc")
	if err!=nil {
		fmt.Println("CoinOutList",err)

		return &CodeModel{Code:4004}
	}
	return &CodeModel{Code:0,Data:map[string]interface{}{
		"count":count,
		"data":list,

	}}

}


func EditCoinOut(outVo *vo.CoinOutVo)*CodeModel{
	if outVo.Status != 2 && (outVo.Status!=3){
		return &CodeModel{Code:4001}
	}
	if outVo.Id == "" {
		return &CodeModel{Code:4001}
	}

	tx:= dao.GetTransaction()
	coinOutDao := &dao.CoinOutDao{BaseDao:dao.NewBaseDao(tx)}
	//userDao := &dao.UserDao{BaseDao:dao.NewBaseDao(tx)}
	coinOut := model.CoinOut{}
	_,err := coinOutDao.First(&coinOut,model.CoinOut{Id:outVo.Id})
	if err!=nil {
		fmt.Println("EditCoinOut",err)
		tx.Rollback()
		return &CodeModel{Code:4005}
	}
	if coinOut.Status !=1 {
		fmt.Println("EditCoinOut",err)
		tx.Rollback()
		return &CodeModel{Code:4007}
	}

	err = coinOutDao.Update(model.CoinOut{Id:coinOut.Id},model.CoinOut{Status:outVo.Status,UpdateTime:time.Now().UnixNano() / 1e6})
	if err!=nil {
		fmt.Println("EditCoinOut",err)
		tx.Rollback()
		return &CodeModel{Code:4005}
	}
	if outVo.Status == 2 {
		tx.Commit()
		return &CodeModel{Code:0}
	}

	err = tx.Exec("update tb_user set usdt_balance = ? + usdt_balance where id = ?",coinOut.Amount,coinOut.UserId).Error
	if err!=nil {
		fmt.Println("EditCoinOut",err)
		tx.Rollback()
		return &CodeModel{Code:4005}
	}

	//user := model.User{}
	//_,err=userDao.First(&user,model.User{ID:coinOut.UserId})
	//if err!=nil {
	//	fmt.Println("EditCoinOut",err)
	//	tx.Callback()
	//	return &CodeModel{Code:4005}
	//}
	//if user.UsdtBalance < coinOut.Amount {
	//	tx.Callback()
	//	return &CodeModel{Code:4003}
	//}
	//balance := user.UsdtBalance - coinOut.Amount
	//err = tx.Exec("update tb_user set usdt_balance = ? where id = ?",balance,user.ID).Error
	////err = userDao.Update(model.User{ID:user.ID},model.User{UsdtBalance:balance})
	//if err!=nil {
	//	fmt.Println("EditCoinOut",err)
	//	tx.Callback()
	//	return &CodeModel{Code:4005}
	//}
	tx.Commit()
	return &CodeModel{Code:0}




}

