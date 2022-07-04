package admin

import (
	"ginApi/common"
	"github.com/gin-gonic/gin"
)

type CasbinController struct {

}

//add rule

func (this *CasbinController) Add(c *gin.Context)  {
	common.Enforcer.LoadPolicy()
	common.ResultSuccess(c,nil)
}
