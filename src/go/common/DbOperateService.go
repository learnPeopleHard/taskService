package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	PASS_WORD = "Twh123456789"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "test"
	CHARSET   = "utf8mb4"
)

// 初始化链接
func MysqlInit() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	log.Println("mysql init start: " + dbDSN)
	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(5000)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(30)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100*time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
	log.Println("mysql init success: " + dbDSN)
}
