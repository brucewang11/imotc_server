package service

import (
	"fmt"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/model"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/brucewang11/frame/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
)

func CreateOrder(vo *vo.OrderVo) *CodeModel{
	if vo.UnitPrice <= 0 || (vo.AskId=="") ||(vo.Price <= 0) {
		return &CodeModel{Code:2001}
	}
	askDao := &dao.AskDao{BaseDao:dao.NewBaseDao(nil)}
	ask := model.Ask{}
	_,err:= askDao.First(&ask,model.Ask{ID:vo.AskId})
	if err!=nil {
		fmt.Println("createorder",err)
		return &CodeModel{Code:2013}
	}
	if vo.Price < ask.MinLimit || (vo.Price>ask.MaxLimit){
		tem := map[string]interface{}{
			"min":decimal.NewFromFloat(ask.MinLimit),
			"max":decimal.NewFromFloat(ask.MaxLimit),
		}
		return &CodeModel{Code:2014,TemplateData:tem}
	}

	fee := vo.Price*0.09
	amount := vo.Price*0.91/vo.UnitPrice
	amount = utils.DecimalTwo(amount)
	//amount = math.Floor(amount)
	orderDB := &model.Order{
		ID:         uuid.NewV4().String(),
		UserName:   vo.UserName,
		Amount:     amount,
		AskId:      vo.AskId,
		Price:      vo.Price,
		CoinType:   1,
		VerifyCode: utils.GetRandomString(6),
		Status:1,
		UserId:vo.UserId,
		CreateTime:time.Now().UnixNano() / 1e6,
		Fee:fee,
		UnitPrice:vo.UnitPrice,

	}
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(nil)}
	err = orderDao.Create(orderDB)
	if err!=nil {
		fmt.Println("createorder",err)
		return &CodeModel{Code:2002}
	}
	return &CodeModel{Code:0,Data:map[string]string{"order_id":orderDB.ID}}
}

func OrderList(vo *vo.OrderVo) *CodeModel{
	orderDB := []model.Order{}
	where := model.Order{UserId:vo.UserId}
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(nil)}
	count,err := orderDao.Count(where,model.Order{})
	if err!=nil {
		fmt.Println("orderlist",err)
		return &CodeModel{Code:2003}
	}
	err = orderDao.ListByPage(where,&orderDB,&dao.Page{PageSize:10,PageIndex:vo.Index},"create_time desc")
	if err!=nil {
		fmt.Println("orderlist",err)
		return &CodeModel{Code:2003}
	}

	return &CodeModel{Code:0,Data:map[string]interface{}{"count":count,"data":orderDB}}
}

func OrderStatus(vo *vo.OrderVo)*CodeModel{
	if vo.ID =="" {
		return &CodeModel{Code:2001}
	}
	if vo.Status !=2 && (vo.Status !=4){
		return &CodeModel{Code:2001}
	}
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(nil)}
	err := orderDao.Update(model.Order{ID:vo.ID},model.Order{Status:vo.Status})
	if err!=nil {
		fmt.Println("OrderStatus",err)
		return &CodeModel{Code:2004}
	}
	return &CodeModel{Code:0}
}
type Orders struct {
	Order model.Order
	Ask   model.Ask
}
func OrderInfo(vo *vo.OrderVo)*CodeModel{
	if vo.ID == ""{
		return &CodeModel{Code:2001}
	}
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(nil)}
	order := model.Order{}
	_,err := orderDao.First(&order,model.Order{ID:vo.ID})
	if err!=nil {
		fmt.Println("OrderStatus",err)
		return &CodeModel{Code:2005}
	}
	ask := model.Ask{}
	askDao := &dao.AskDao{BaseDao:dao.NewBaseDao(nil)}
	_,err = askDao.First(&ask,model.Ask{ID:order.AskId})
	if err!=nil {
		fmt.Println("OrderStatus",err)
		return &CodeModel{Code:2005}
	}
	return &CodeModel{Code:0,Data:Orders{
		Order: order,
		Ask:   ask,
	}}

}




//---------------后台管理-----------------

func EditOrder(vo *vo.OrderVo) *CodeModel{
	if vo.ID == "" {
		return &CodeModel{Code:2001}
	}
	if vo.Status !=3 && (vo.Status!=4) {
		return &CodeModel{Code:2001}
	}
	tx := dao.GetTransaction()
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(tx)}
	order := model.Order{}
	_,err := orderDao.First(&order,model.Order{ID:vo.ID})
	if err!=nil {
		tx.Rollback()
		return &CodeModel{Code:2004}
	}
	if order.Status != 2{
		return  &CodeModel{Code:2012}
	}

	err = orderDao.Update(model.Order{ID:vo.ID},model.Order{Status:vo.Status})

	if err!=nil {
		tx.Rollback()
		fmt.Println("editorder",err)
		return &CodeModel{Code:2004}
	}
	if vo.Status == 4 {
		tx.Commit()
		return &CodeModel{Code:0}
	}
	err = tx.Exec("update tb_user set usdt_balance = usdt_balance + ? where id = ?",order.Amount,order.UserId).Error
	if err!=nil {
		tx.Rollback()
		fmt.Println("editorder",err)
		return &CodeModel{Code:2004}
	}
	tx.Commit()
	return &CodeModel{Code:0}
}

func OrderListAdmin(vo *vo.OrderVo)*CodeModel{
	orderDB := []model.Order{}
	where := model.Order{UserId:vo.UserId}
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(nil)}
	count,err := orderDao.Count(where,model.Order{})
	if err!=nil {
		fmt.Println("orderlist",err)
		return &CodeModel{Code:2003}
	}
	err = orderDao.ListByPage(where,&orderDB,&dao.Page{PageSize:10,PageIndex:vo.Index},"create_time desc")
	if err!=nil {
		fmt.Println("orderlist",err)
		return &CodeModel{Code:2003}
	}

	return &CodeModel{Code:0,Data:map[string]interface{}{"count":count,"data":orderDB}}
}

