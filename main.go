package main

import (
	"fmt"
	"ginApi/conf"
	"ginApi/router"
	"github.com/gin-gonic/gin"
)

func main()  {
	//初始化配置
	config := conf.InitConf()
	//初始化mysql
	sqlConn := config.SqlConn
	fmt.Println(sqlConn)

	//初始化redis
	fmt.Println("redis db num:",config.RedisConf.Db)

	//初始化日志

	//初始化路由
	route := gin.Default()

	router.InitRouter(route)

	//项目启动
	fmt.Println("list port:",config.AppConf.Port)
	_ = route.Run(config.AppConf.Port)
}

