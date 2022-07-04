package admin

import (
	"ginApi/common"
	"github.com/gin-gonic/gin"
)

type OrderController struct {

}

//订单列表
func (this *OrderController) GetItems (c *gin.Context) {
	common.ResultSuccess(c,"admin order list")
}

//订单详情
func (this *OrderController) GetItem (c *gin.Context) {
	common.ResultSuccess(c,"admin order detail")
}