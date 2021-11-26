package biz

import "loginService/src/go/login"

func Login( rv login.LoginRequestDTO) login.LoginResponseDTO {
	if rv.Name=="" ||rv.Password==""{
		return login.LoginResponseDTO{Code: 111, Message: "参数不能为空"}
	}
	flag,message,user := login.UserNameCheck(rv.Name,rv.Password)
	if flag{
		return login.LoginResponseDTO{Code: 100, Message: message,UserId: user.Id,UserName: user.Name}
	} else {
		return login.LoginResponseDTO{Code: 111, Message: message}
	}
}
