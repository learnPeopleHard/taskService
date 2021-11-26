package redis

import (
	"fmt"
	cluster "github.com/chasex/redis-go-cluster"
	"loginService/src/go/common"
	"time"
)

func ClusterGetkey(key string)  (string,error){
	startTime:=time.Now().UnixMilli()
	c := common.RedisClient.Get()
	defer c.Close()
	fmt.Printf("redis get pool: %s   time:%d \n", key,time.Now().UnixMilli()-startTime)
	result, err := cluster.String(common.RedisClusterClient.Do("Get",key))
	if err != nil {
		fmt.Printf("Getkey redis key %s err:%v \n",key,err)
		return "",err
	}
	fmt.Printf("redis do set : %s   time:%d \n", key,time.Now().UnixMilli()-startTime)
	return result, err
}


func ClusterSetkey(key string,value string) error {
	c := common.RedisClient.Get()
	defer c.Close()

	_, err := c.Do("Set", key, value)
	if err != nil {
		fmt.Printf("Setkey redis key %s err:%v \n",key,err)
		return err
	}
	return nil
}

func ClusterSetkeyExpire(key string,value string,expire int) error {
	c := common.RedisClient.Get()
	defer c.Close()

	_, err := c.Do("setex", key, expire,value)
	if err != nil {
		fmt.Printf("SetkeyExpire redis key %s err:%v \n", key,err)
		return err
	}
	return nil
}