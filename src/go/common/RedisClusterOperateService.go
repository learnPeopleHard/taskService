package common

import (
	cluster "github.com/chasex/redis-go-cluster"
	"log"
	"time"
)

var RedisClusterClient cluster.Cluster

// 初始化链接
func RedisClusterInit() {
	log.Println("redis pool init start: " )
	var err error
	RedisClusterClient,err = cluster.NewCluster(
		&cluster.Options{
			StartNodes: []string{"127.0.0.1:6379", "127.0.0.1:6381", "127.0.0.1:6382"},
			ConnTimeout: 50 * time.Millisecond,
			ReadTimeout: 50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive: 16,
			AliveTime: 60 * time.Second,
		})
	if err!=nil{
		panic("创建redis 集群 失败")
	}
	log.Println("redis pool init success: " )
}
