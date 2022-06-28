package main

import (
	"fmt"
	"ginApi/cache"
	"ginApi/conf"
	"ginApi/model"
	"ginApi/router"
	"github.com/gin-gonic/gin"
)

func main()  {
	//初始化配置
	config := conf.InitConf()

	//初始化路由
	route := gin.Default()
	router.InitRouter(route)

	//初始化mysql
	sqlConn := config.SqlConn
	fmt.Println("mysql conn:",sqlConn)
	model.InitDB(sqlConn) //gorm


	//初始化redis
	fmt.Println("redis db num:",config.RedisConf.Db)
	cache.InitRedis(config.RedisConf)

	//初始化日志


	//项目启动
	fmt.Println("listen port:",config.AppConf.Port)
	_ = route.Run(config.AppConf.Port)
}

