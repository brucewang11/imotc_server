package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/model"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/brucewang11/frame/utils"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"time"
)
var salt = "wangbruce"
var codeKey = "code_"
func Register(userVo *vo.UserVo) *CodeModel{
	if userVo.Account =="" || (userVo.Code=="") || (userVo.Pwd==""){
		return &CodeModel{Code:1001}
	}
	userDao := &dao.UserDao{BaseDao:dao.NewBaseDao(nil)}
	param := &model.User{
		Account:     userVo.Account,
	}
	_,err := userDao.First(&model.User{},param)
	if err==nil {
		return &CodeModel{Code:1002}
	}

	key := codeKey+userVo.Account+"_1"
	redisCode,err := dao.RedisClient.Get(key).Result()
	if redisCode != userVo.Code {
		return &CodeModel{Code:1007}
	}

	sha := sha256.New()
	sha.Write([]byte(userVo.Pwd))
	sha.Write([]byte(salt))
	pass := sha.Sum(nil)
	userDB := &model.User{
		ID:           uuid.NewV4().String(),
		Account:     userVo.Account,
		UsdtBalance: 0,
		Pwd:hex.EncodeToString(pass),
		CreateTime:time.Now().UnixNano() / 1e6,
	}
	err = userDao.Create(userDB)
	if err!=nil {
		fmt.Println("register",err)
		return &CodeModel{Code:1003}
	}

	return  &CodeModel{Code:0}
}

func Load (userVo *vo.UserVo)*CodeModel{
	if userVo.Account == "" || (userVo.Pwd == ""){
		return &CodeModel{Code:1001}
	}
	sha := sha256.New()
	sha.Write([]byte(userVo.Pwd))
	sha.Write([]byte(salt))
	pass := sha.Sum(nil)
	param := &model.User{
		Account:     userVo.Account,
		Pwd:         hex.EncodeToString(pass),

	}
	fmt.Println(param)
	userDao := &dao.UserDao{BaseDao:dao.NewBaseDao(nil)}
	user := &model.User{}
	_,err := userDao.First(user,param)
	if err!=nil {
		return &CodeModel{Code:1004}
	}

	jwts := utils.NewJWT()
	token,err := jwts.CreateToken(utils.CustomClaims{
		ID:             user.ID,
		Account:        user.Account,
		UserType:       2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:time.Now().Add(24 * time.Hour).Unix(),
		},
	})
	if token == "" || (err!=nil) {
		return &CodeModel{Code:1005}
	}
	loadRes := vo.LoadRes{
		Account:user.Account,
		ID:user.ID,
		Token:token,
	}

	return &CodeModel{Code:0,Data:loadRes}
}

func FindUserById(userVo *vo.UserVo)*CodeModel{
	userDao := &dao.UserDao{BaseDao:dao.NewBaseDao(nil)}
	user := model.User{}
	_,err := userDao.First(&user,model.User{ID:userVo.UserId})
	if err!=nil {
		return &CodeModel{Code:1001}
	}
	return &CodeModel{Code:0,Data:map[string]interface{}{"user_id":userVo.UserId,"balance":user.UsdtBalance}}
}

//发送验证码
func SendCode(userVo *vo.UserVo) *CodeModel{
	if userVo.Account=="" || (userVo.CodeType==""){
		return &CodeModel{Code:1001}
	}
	code := utils.GenValidateCode(6)
	client := utils.New("475611044@qq.com", "omxpmtpnejglbiej", "ImOTC", "smtp.qq.com", 465, true)
	if err := client.SendEmail([]string{userVo.Account},"ImOTC验证码","您的验证码是"+code);err !=nil{
		fmt.Println("发送验证码错误",err)
	}


	//err := utils.SendToMail(userVo.Account,"ImOTC验证码","您的验证码是："+code,"")
	//if err !=nil {
	//	fmt.Println("发送验证码错误",err)
	//	return &CodeModel{Code:1010}
	//}
	dao.RedisClient.Set(codeKey+userVo.Account+"_"+userVo.CodeType,code,time.Minute)
	return  &CodeModel{Code:0}
}


//---------------------------后台管理----------------

func UserList(userVo *vo.UserVo)*CodeModel{
	UserDB := []model.User{}
	orderDao := &dao.OrderDao{BaseDao:dao.NewBaseDao(nil)}
	count,err := orderDao.Count(model.User{},model.User{})
	err = orderDao.ListByPage(model.User{},&UserDB,&dao.Page{PageSize:10,PageIndex:userVo.Index},"create_time desc")
	if err!=nil {
		fmt.Println("UserList",err)
		return &CodeModel{Code:1003}
	}
	return &CodeModel{Code:0,Data:map[string]interface{}{"count":count,"data":UserDB}}
}



