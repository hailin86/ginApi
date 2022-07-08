package main

import (
	"fmt"
	"ginApi/cache"
	"ginApi/common"
	"ginApi/middleware"
	"ginApi/model"
	"ginApi/router"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func main()  {

	v := 1
	value := strings.TrimSpace(fmt.Sprintf("%v", v))

	fmt.Println("type =",reflect.TypeOf(value))
	fmt.Println("value =",value)


	//初始化配置
	configPath := "./conf"
	config := common.InitConf(configPath)

	//初始化mysql
	sqlConn := config.SqlConn
	//fmt.Println("mysql conn:",sqlConn)
	model.InitDB(sqlConn) //gorm

	//初始化redis
	//fmt.Println("redis db num:",config.RedisConf.Db)
	cache.InitRedis(config.RedisConf)

	//初始化日志
	common.InitLogger(config.LoggerConf)


	//初始化路由
	//route := gin.Default()
	route := gin.New()
	route.Use(middleware.GinLogger(),middleware.GinRecovery(false))
	//使用中间件解决跨域问题
	route.Use(middleware.Cors())

	router.InitRouter(route)
	//项目启动
	//gin.SetMode(gin.ReleaseMode)
	//fmt.Println("app listen port:",config.AppConf.Port)


	_ = route.Run(config.AppConf.Port)
}

