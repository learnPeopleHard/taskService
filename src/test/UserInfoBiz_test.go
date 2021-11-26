package test

import (
	"github.com/agiledragon/gomonkey"
	"loginService/src/go/biz"
	"loginService/src/go/dto/userinfodto"
	"loginService/src/go/userinfo"
	"testing"
)

func Test_UpdateUserInfo(t *testing.T)  {

	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{&userinfo.UserInfoPO{Id: 1, UserId: 1, UserName: "1", Nickname: "1", ProfilePicture: "1"}}},// 模拟函数的第1次输出
		{Values: gomonkey.Params{&userinfo.UserInfoPO{Id: 1, UserId: 1, UserName: "1", Nickname: "1", ProfilePicture: "1"}}},// 模拟函数的第2次输出
		{Values: gomonkey.Params{nil}},// 模拟函数的第3次输出
	}

	//outputs2 := []gomonkey.OutputCell{
	//	{Values: gomonkey.Params{func(userId int64) *userinfo.UserInfoPO{ return &userinfo.UserInfoPO{Id: 1, UserId: userId, UserName: "1", Nickname: "1", ProfilePicture: "1"}},nil}},// 模拟函数的第1次输出
	//	{Values: gomonkey.Params{func(userId int64) *userinfo.UserInfoPO{ return &userinfo.UserInfoPO{Id: 1, UserId: userId, UserName: "1", Nickname: "1", ProfilePicture: "1"}},nil}},// 模拟函数的第1次输出
	//	{Values: gomonkey.Params{nil,nil}},// 模拟函数的第3次输出
	//}
	patches := gomonkey.ApplyFuncSeq(userinfo.QueryUserInfoById, outputs)
	defer patches.Reset()
	//patches := gomonkey.ApplyFunc(userinfo.QueryUserInfoById, func(userId int64)*userinfo.UserInfoPO{
	//	return &userinfo.UserInfoPO{Id: 1, UserId: userId, UserName: "1", Nickname: "1", ProfilePicture: "1"}
	//})
	//defer patches.Reset()

	patches2 := gomonkey.ApplyFunc(userinfo.UpdateUserInfo, func(userId int64,nickName string,ProfilePicture string) bool{
		return true
	})
	defer patches2.Reset()

	patches3 := gomonkey.ApplyFunc(userinfo.InsertUserInfo, func(userId int64,nickName string,ProfilePicture string,userName string) bool{
		return true
	})
	defer patches3.Reset()
	resp := biz.UpdateUserInfo(userinfodto.UserInfoDTO{UserId: 100,Nickname: "1",ProfilePicture: "1"})
	if resp.Code!=222 {
		t.Errorf("expected %v, got %v", 222, resp.Code)
	}

	resp = biz.UpdateUserInfo(userinfodto.UserInfoDTO{UserId: 100,Nickname: "2",ProfilePicture: "1"})
	if resp.Code!=100 {
		t.Errorf("expected %v, got %v", 100, resp.Code)
	}

	resp = biz.UpdateUserInfo(userinfodto.UserInfoDTO{UserId: 100,Nickname: "2",ProfilePicture: "1"})
	if resp.Code!=100 {
		t.Errorf("expected %v, got %v", 100, resp.Code)
	}

}
