package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"loginService/src/go/common"
	"time"
)

func Getkey(key string)  (string,error){
	startTime:=time.Now().UnixMilli()
	c := common.RedisClient.Get()
	defer c.Close()
	fmt.Printf("redis get pool: %s   time:%d \n", key,time.Now().UnixMilli()-startTime)
	result, err := redis.String(c.Do("Get", key))
	if err != nil {
		fmt.Printf("Getkey redis key %s err:%v \n",key,err)
		return "",err
	}
	fmt.Printf("redis do set : %s   time:%d \n", key,time.Now().UnixMilli()-startTime)
	return result, err
}


func Setkey(key string,value string) error {
	c := common.RedisClient.Get()
	defer c.Close()

	_, err := c.Do("Set", key, value)
	if err != nil {
		fmt.Printf("Setkey redis key %s err:%v \n",key,err)
		return err
	}
	return nil
}

func SetkeyExpire(key string,value string,expire int) error {
	c := common.RedisClient.Get()
	defer c.Close()

	_, err := c.Do("setex", key, expire,value)
	if err != nil {
		fmt.Printf("SetkeyExpire redis key %s err:%v \n", key,err)
		return err
	}
	return nil
}
