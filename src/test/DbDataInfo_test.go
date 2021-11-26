package test

import (
	"fmt"
	"github.com/msterzhang/gpool"
	"loginService/src/go/common"
	"loginService/src/go/userinfo"
	"strconv"
	"testing"
)

func test_dnInit(t *testing.T)  {
	common.MysqlInit()
	size:=20
	pool := gpool.New(size)
	for i :=723554;i<10000000;i++{
		pool.Add(1)
		go RunTest(
			i,
			pool,
		)
	}
	pool.Wait()
}

func Test_userinfo(t *testing.T)  {
	common.MysqlInit()
	for i:=1; i<100000;i++{
		userPO := userinfo.QueryUser("zhangsan"+strconv.Itoa(i))
		if userPO!=nil{
			userInfoPO := userinfo.QueryUserInfoById(userPO.Id)
			if userInfoPO==nil{
				userinfo.InsertUserInfo(userPO.Id,"123213","123213",userPO.Name)
			}
		}
	}

}
func RunTest(i int,pool *gpool.Pool)  {
	var name = "zhangsan"+strconv.Itoa(i)
	userinfo.InsertUser(name,"e10adc3949ba59abbe56e057f20f883e")
	fmt.Println("执行====",i)
	pool.Done()
}
