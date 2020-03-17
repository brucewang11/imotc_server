package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/model"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/brucewang11/frame/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func LoadAdmin(userVo *vo.UserVo) *CodeModel{
	if userVo.Account == "" || (userVo.Pwd == ""){
		return &CodeModel{Code:1001}
	}
	sha := sha256.New()
	sha.Write([]byte(userVo.Pwd))
	sha.Write([]byte(salt))
	pass := sha.Sum(nil)
	param := &model.SysUser{
		Account:     userVo.Account,
		Pwd:         hex.EncodeToString(pass),

	}

	userDao := &dao.SysUserDao{BaseDao:dao.NewBaseDao(nil)}
	user := &model.SysUser{}
	_,err := userDao.First(user,param)
	if err!=nil {
		return &CodeModel{Code:1004}
	}

	jwts := utils.NewJWT()
	token,err := jwts.CreateToken(utils.CustomClaims{
		ID:             user.ID,
		Account:        user.Account,
		UserType:       1,
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

func saltUtils(pwd string) string{
	sha := sha256.New()
	sha.Write([]byte(pwd))
	sha.Write([]byte(salt))
	pass := sha.Sum(nil)
	return hex.EncodeToString(pass)
}


func EditPwd(vo *vo.UserVo) *CodeModel{
	if vo.Pwd == "" ||(vo.OldPwd == ""){
		return &CodeModel{Code:1001}
	}
	userDao := &dao.SysUserDao{BaseDao:dao.NewBaseDao(nil)}
	user := model.SysUser{}
	_,err := userDao.First(&user,model.SysUser{ID:"1"})
	if err!=nil {
		return &CodeModel{Code:1001}
	}
	old := saltUtils(vo.OldPwd)
	if old != user.Pwd {
		return &CodeModel{Code:1006}
	}
	sha := sha256.New()
	sha.Write([]byte(vo.Pwd))
	sha.Write([]byte(salt))
	pass := sha.Sum(nil)
	err = userDao.Update(model.SysUser{ID:"1"},model.SysUser{Pwd:hex.EncodeToString(pass)})
	if err!=nil {
		return &CodeModel{Code:1011}
	}
	return &CodeModel{Code:0}

}

