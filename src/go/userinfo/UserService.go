package userinfo

import (
	"fmt"
	"loginService/src/go/common"
)

// 查询数据，取所有字段
func QueryUser(name string) *UserPO {

	var user = new(UserPO)
	row := common.MysqlDb.QueryRow("select id, name, pwd from user where name=? limit 1",name)
	if err :=row.Scan(&user.Id,&user.Name,&user.Pwd); err != nil{
		return nil
	}
	return user
}

func InsertUser(userName string,pwd string ) bool{
	ret,_ := common.MysqlDb.Exec("insert into user(name, pwd) values (?,?) ",userName,pwd)
	upd_nums,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",upd_nums)
	return upd_nums>0
}
