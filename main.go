package main

import (
	"fmt"
	"ginApi/cache"
	"ginApi/conf"
	"ginApi/middleware"
	"ginApi/model"
	"ginApi/router"
	"github.com/gin-gonic/gin"
)

func main()  {
	//初始化配置
	config := conf.InitConf()

	//初始化mysql
	sqlConn := config.SqlConn
	fmt.Println("mysql conn:",sqlConn)
	model.InitDB(sqlConn) //gorm

	//初始化redis
	//fmt.Println("redis db num:",config.RedisConf.Db)
	cache.InitRedis(config.RedisConf)

	//初始化日志


	//初始化路由
	route := gin.Default()
	//使用中间件解决跨域问题
	route.Use(middleware.Cors())
	//校验接口调用者身份
	//route.Use(middleware.CheckAppInstance())
	router.InitRouter(route)
	//项目启动
	//gin.SetMode(gin.ReleaseMode)
	fmt.Println("app listen port:",config.AppConf.Port)
	_ = route.Run(config.AppConf.Port)
}

