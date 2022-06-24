package router

import (
	"ginApi/controller/admin"
	"ginApi/controller/api"
	"ginApi/controller/index"
	"github.com/gin-gonic/gin"
)

func InitRouter(route *gin.Engine)   {
	//业务接口
	index.InitRouter(route)

	//管理后台
	admin.InitRouter(route)

	//对外openApi
	api.InitRouter(route)
}
