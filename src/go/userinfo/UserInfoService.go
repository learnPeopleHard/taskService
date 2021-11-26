package userinfo

import (
	"fmt"
	"loginService/src/go/common"
)

// QueryUserInfoById 查询数据，取所有字段
func QueryUserInfoById(userId int64) *UserInfoPO {
	user := new(UserInfoPO)
	row := common.MysqlDb.QueryRow("select * from user_info where user_id=? limit 1",userId)
	if err :=row.Scan(&user.Id,&user.UserId,&user.UserName,&user.Nickname,&user.ProfilePicture); err != nil{
		return nil
	}
	return user
}

// UpdateUserInfo 更新数据
func UpdateUserInfo(userId int64,nickName string,ProfilePicture string ) bool{
	ret,_ := common.MysqlDb.Exec("UPDATE user_info set nickname=?,profile_picture=? where user_id=?",nickName,ProfilePicture,userId)
	upd_nums,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",upd_nums)
	return upd_nums>0
}

// InsertUserInfo 更新数据
func InsertUserInfo(userId int64,nickName string,ProfilePicture string,userName string ) bool{
	ret,_ := common.MysqlDb.Exec("insert into user_info(user_id,nickname,profile_picture,username) values (?,?,?,?) ",userId,nickName,ProfilePicture,userName)
	upd_nums,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",upd_nums)
	return upd_nums>0
}