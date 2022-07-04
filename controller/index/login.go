package index

import (
	"ginApi/common"
	"github.com/gin-gonic/gin"
)

type LoginController struct {

}

type LoginParam struct {
	Name string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}


func (this *LoginController) LoginByForm (c *gin.Context) {
	var param LoginParam
	if err := c.Bind(&param); err != nil {
		common.LogError(c,"loginByForm param err:",err)
		common.ResultFailed(c,"param error",1,nil)
		return
	}
	common.ResultSuccess(c,param)
}



func (this *LoginController) LoginByJson (c *gin.Context) {
	var param LoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		common.LogError(c,"loginByJson param err:",err)
		common.ResultFailed(c,"param error",1,nil)
		return
	}
	if param.Name =="" {
		common.ResultFailed(c,"param name required",100,nil)
		return
	}


	common.ResultSuccess(c,param)
}


func (this *LoginController) LoginByUrl (c *gin.Context) {


	common.ResultSuccess(c,nil)
}