package test

import (
	"fmt"
	"github.com/agiledragon/gomonkey"
	"loginService/src/go/biz"
	"loginService/src/go/login"
	"loginService/src/go/userinfo"
	"testing"
)

func Test_Login_f(t *testing.T)  {
	patches := gomonkey.ApplyFunc(userinfo.QueryUser, func(name string)*userinfo.UserPO{
		return &userinfo.UserPO{Id: 1, Name: "zhangsan", Pwd: "1"}
	})
	defer patches.Reset()
	req:=login.LoginRequestDTO{Name:"zhangsan",Password: "1"}
	resp := biz.Login(req)
	if resp.Code != 100 {
		//t.Errorf("error expected %v, got %v", 100, resp.Code)
		fmt.Printf("expected %v, got %v\n success ", 100, resp.Code)
	}

	req2:=login.LoginRequestDTO{Name:"",Password: "1"}
	resp2 := biz.Login(req2)
	if resp2.Message!="参数不能为空"{
		t.Errorf("error expected %v, got %v", "参数不能为空", resp2.Message)
	}
}

func Test_Login_s(t *testing.T)  {
	patches := gomonkey.ApplyFunc(userinfo.QueryUser, func(name string)*userinfo.UserPO{
		return &userinfo.UserPO{Id: 1, Name: "zhangsan", Pwd: "c4ca4238a0b923820dcc509a6f75849b"}
	})
	defer patches.Reset()
	req:=login.LoginRequestDTO{Name:"zhangsan",Password: "1"}
	resp := biz.Login(req)
	if resp.Code != 100 {
		t.Errorf("error expected %v, got %v", 100, resp.Code)
	}
	if resp.UserId!=1{
		t.Errorf("error expected userid %v, got %v", 1, resp.UserId)
	}
}
