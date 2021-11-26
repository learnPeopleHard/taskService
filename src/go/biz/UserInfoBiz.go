package biz

import (
	"loginService/src/go/common"
	"loginService/src/go/dto/userinfodto"
	"loginService/src/go/userinfo"
)

func UpdateUserInfo( rv userinfodto.UserInfoDTO) common.BaseResponseDTO {
	userInfoPO := userinfo.QueryUserInfoById(rv.UserId)
	var flag bool
	if userInfoPO==nil{
		flag = userinfo.InsertUserInfo(rv.UserId,rv.Nickname,rv.ProfilePicture,rv.UserName)
	}else{
		//判断属性一致不用deepEqual这里主要是对应的属性判断
		if userInfoPO.Nickname==rv.Nickname &&  userInfoPO.ProfilePicture== rv.ProfilePicture{
			return common.BaseResponseDTO{Code: 222, Message: "值相同不用修改"}
		}
		flag = userinfo.UpdateUserInfo(rv.UserId,rv.Nickname,rv.ProfilePicture)
	}
	if flag{
		return common.BaseResponseDTO{Code: 100, Message: "修改成功"}
	} else {
		return common.BaseResponseDTO{Code: 111, Message: "修改失败"}
	}
}
