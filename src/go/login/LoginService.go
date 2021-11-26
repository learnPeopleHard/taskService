package login

import (
	"loginService/src/go/common"
	"loginService/src/go/userinfo"
)

// 查询数据，取所有字段
func UserNameCheck(name string,pwd string) (bool,string,*userinfo.UserPO) {
	user := userinfo.QueryUser(name)
	if user==nil{
		return false,"用户名||或者密码 不对",nil
	}
	if user.Pwd!=common.Md5V(pwd) {
		return false,"用户名||或者密码 不对",nil
	}
	return  true,"登录成功",user
}
