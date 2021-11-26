package test

import (
	"github.com/agiledragon/gomonkey"
	"log"
	"loginService/src/go/userinfo"
	"testing"
)

func Test_QueryUserInfoById(t *testing.T)  {
	patches := gomonkey.ApplyFunc(userinfo.QueryUserInfoById, func(userId int64)*userinfo.UserInfoPO{
		return &userinfo.UserInfoPO{Id: 1, UserId: 1, UserName: "1", Nickname: "1", ProfilePicture: "1"}
	})
	defer patches.Reset()

	info := userinfo.TestAAA(123)
	if info.UserId != 1 {
		t.Errorf("expected %v, got %v", 1, info.UserId)
	}
	log.Println("ceshi =  ",info)
}
