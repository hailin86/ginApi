package main

import (
	"fmt"
	"ginApi/conf"
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

	//初始化中间件 权限校验 跨域问题解决等

	//项目启动
	_ = route.Run(config.AppConf.Port)
}

