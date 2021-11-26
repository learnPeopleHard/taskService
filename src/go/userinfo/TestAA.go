package userinfo


func TestAAA(userId int64)  *UserInfoPO {
	return QueryUserInfoById(userId)
}
