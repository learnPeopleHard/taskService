package test

import (
	"fmt"
	"loginService/src/go/common"
	"loginService/src/go/userinfo"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func Test_token(t *testing.T)  {
	common.MysqlInit()
	for i:=1 ;i<200;i++{
		userInfoPO := userinfo.QueryUser("zhangsan"+strconv.Itoa(i))
		token,_:= common.BuildWebToken(userInfoPO.Id,userInfoPO.Name,time.Now().UnixMilli())
		fmt.Println(strconv.FormatInt(userInfoPO.Id,10)+","+url.QueryEscape(token))
	}


}
