package main

import (
	"errors"
	"fmt"
	"log"
	"loginService/src/go/biz"
	"loginService/src/go/common"
	"loginService/src/go/dto"
	"loginService/src/go/dto/userinfodto"
	"loginService/src/go/login"
	"loginService/src/go/redis"
	"loginService/src/go/userinfo"
	"net"
	"net/rpc"
	"strconv"
	"time"
)

type Listener int

//type Arith int

func (l *Listener) Login(line []byte, reply *login.LoginResponseDTO) error {
	startTime:= time.Now().UnixMilli()
	rv, _ := login.Reqdecode(line)
	log.Printf("Receive: %v \n", rv)
	common.Try(func() {
		*reply = biz.Login(rv)
	}, func(err interface{}) {
		*reply = login.LoginResponseDTO{Code: 111, Message: "内部异常"}
	})
	log.Printf("userName %s login time is %d ", rv.Name, time.Now().UnixMilli()-startTime)
	return nil
}

func (l *Listener) QueryUserInfo(line []byte, reply *userinfodto.UserInfoDTO) error {
	startTime:= time.Now().UnixMilli()
	rv := string(line)
	log.Printf("Receive: %v \n", rv)
	var err1 error =nil
	common.Try(func() {
		userId,_ := strconv.ParseInt(rv,10,64)
		userInfoPO := userinfo.QueryUserInfoById(userId)
		*reply = userinfodto.UserInfoDTO{Id: userInfoPO.Id,UserId: userInfoPO.UserId,Nickname: userInfoPO.Nickname,ProfilePicture: userInfoPO.ProfilePicture}
	}, func(err interface{}) {
		*reply = userinfodto.UserInfoDTO{}
		err1 = errors.New("查询异常")
	})
	log.Printf("userId %s query time is %d ", rv, time.Now().UnixMilli()-startTime)
	return err1
}

func (l *Listener) UpdateUserInfo(line []byte, reply *common.BaseResponseDTO) error {
	rv, _ := userinfodto.Decode(line)
	fmt.Printf("Receive: %v\n", rv)
	common.Try(func() {
		*reply =  biz.UpdateUserInfo(rv)
	}, func(err interface{}) {
		*reply = common.BaseResponseDTO{Code: 112, Message: "内部异常"}
	})
	return nil
}

func (l *Listener) RedisSetToken(line []byte, reply *common.BaseResponseDTO) error {
	rv,_:= common.Decode(line)
	fmt.Printf("Receive: %v\n", rv)
	common.Try(func() {
		flag := redis.SetkeyExpire(rv.Token,strconv.FormatInt(rv.UserId,10),30*60)
		if flag==nil{
			*reply = common.BaseResponseDTO{Code: 100, Message: "设置token成功"}
		} else {
			*reply = common.BaseResponseDTO{Code: 111, Message: "设置token失败"}
		}
	}, func(err interface{}) {
		*reply = common.BaseResponseDTO{Code: 112, Message: "内部异常"}
	})
	return nil
}

func (l *Listener) RedisGetToken(line []byte, reply *common.TokenDTO) error {
	rv:=string(line)
	fmt.Printf("Receive: %v\n", rv)
	common.Try(func() {
		value,err := redis.Getkey(rv)
		user := common.TokenDTO{}
		if err==nil &&  value!="" {
			user.UserId,_ = strconv.ParseInt(value,10,64)
		}
		*reply=user
	}, func(err interface{}) {
		*reply = common.TokenDTO{}
	})
	return nil
}

func (l *Listener) RedisGetAndSetToken(line []byte, reply *common.TokenDTO) error {
	rv,_:=dto.CheckDecode(line)
	startTime:=time.Now().UnixMilli()
	fmt.Printf("RedisGetAndSetToken Receive: %v   time:%d \n", rv,startTime)
	common.Try(func() {
		value,err := redis.Getkey(rv.Token)
		user := common.TokenDTO{}
		if err==nil &&  value!="" {
			user.UserId,_ = strconv.ParseInt(value,10,64)
			if rv.FlagGetAndSet{
				common.Try(func() {
					redis.Setkey(rv.Token,value)
				}, func(err interface{}) {
					log.Printf("get and set token error %s \n",err)
				})
			}
		}
		*reply=user
	}, func(err interface{}) {
		*reply = common.TokenDTO{}
	})
	return nil
}

func main() {
	fmt.Printf("service start: ")
	addy, err := net.ResolveTCPAddr("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	common.MysqlInit()
	common.RedisInit()

	listener := new(Listener)
	err1 := rpc.Register(listener)
	if err1 != nil {
		return
	}
	rpc.Accept(inbound)

	//l, err := net.Listen("tcp", "127.0.0.1:12345")
	//if err != nil {
	//	log.Fatal("listen error:", err)
	//}
	//
	//arith := new(Arith)
	//rpc.Register(arith)
	//
	//for {
	//	conn, err := l.Accept()
	//	if err != nil {
	//		log.Fatal("accept error:", err)
	//	}
	//
	//	timeout := 10*time.Second
	//	conn.SetDeadline(time.Now().Add(timeout))
	//
	//	// 注意这一行
	//	go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	//}
}