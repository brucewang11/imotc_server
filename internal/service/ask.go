package service

import (
	"fmt"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/model"
	"github.com/brucewang11/frame/internal/vo"
)

func AskList() *CodeModel{
	AskDao := &dao.AskDao{BaseDao:dao.NewBaseDao(nil)}
	askList := []model.Ask{}
	err := AskDao.List(model.Ask{},&askList)
	if err!=nil {
		fmt.Println(err)
		return &CodeModel{Code:3002}
	}

	return &CodeModel{Code:0,Data:askList}
}


func EditAsk(ask *vo.AskVo)*CodeModel{
	askDao := &dao.AskDao{BaseDao:dao.NewBaseDao(nil)}
	askDB := model.Ask{

		UserName:    ask.UserName,
		Amount:      ask.Amount,
		MinLimit:    ask.MinLimit,
		MaxLimit:    ask.MaxLimit,
		Price:       ask.Price,
		BankCard:    ask.BankCard,
		BankName:    ask.BankName,
		BankAddress: ask.BankAddress,
	}
	err := askDao.Update(model.Ask{ID:ask.ID},askDB)
	if err!=nil {
		fmt.Println(err)
		return &CodeModel{Code:3003}
	}
	return &CodeModel{Code:0}
}