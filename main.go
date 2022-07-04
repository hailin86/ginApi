package main

import (
	"ginApi/cache"
	"ginApi/common"
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
	//fmt.Println("mysql conn:",sqlConn)
	model.InitDB(sqlConn) //gorm



	//初始化redis
	//fmt.Println("redis db num:",config.RedisConf.Db)
	cache.InitRedis(config.RedisConf)

	//初始化日志
	common.InitLogger()

	_ = common.InitCasbinEnforcer()


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

