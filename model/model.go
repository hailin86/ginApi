package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB(sqlConn string)  {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,         // 禁用彩色打印
			IgnoreRecordNotFoundError: true, //禁用 当查询 record not not 的时候不打印日志
		},
	)

	db, err := gorm.Open(mysql.Open(sqlConn), &gorm.Config{

		Logger:newLogger,
		//禁用表名复数
		NamingStrategy:schema.NamingStrategy{
			SingularTable: true,
		},

		//创建并缓存预编译语句
		PrepareStmt:true,
		//跳过默认事务
		SkipDefaultTransaction:true,

	})

	// Error
	if sqlConn == "" || err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	//设置连接池 一般情况上面的两个函数是一起使用的，而且最大连接数的设置，必须要大于最大可空闲连接数
	//最大空闲数
	sqlDB.SetMaxIdleConns(10)
	//最大链接数
	sqlDB.SetMaxOpenConns(100)

	DB = db

}
