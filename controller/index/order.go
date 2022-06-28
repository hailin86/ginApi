package index

import "github.com/gin-gonic/gin"

type OrderController struct {

}


//订单列表
func (this *OrderController) GetItems (c *gin.Context) {

	userId,flag := c.Get("userId")
	if flag {
		c.JSON(200,gin.H{"code":200,"msg":"success","data":userId})
	}else {
		c.JSON(200,gin.H{"code":200,"msg":"success","data":0})
	}


}

//订单详情
func (this *OrderController) GetItem (c *gin.Context) {

}
