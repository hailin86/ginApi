package index

import (
	"ginApi/common"
	"ginApi/model"
	"github.com/gin-gonic/gin"
	"time"
)

type ProductController struct {

}

//商品列表
func (this *ProductController) GetItems (c *gin.Context)  {
	list := []model.Product{
		{
			Id: 1,
			Name:     "test001",
			Price:    100,
			Status:   1,
			Created:  time.Now(),
			Modified: time.Now(),
		},
		{
			Id: 2,
			Name:     "test002",
			Price:    65,
			Status:   1,
			Created:  time.Now(),
			Modified: time.Now(),
		},
		{
			Id: 3,
			Name:     "test003",
			Price:    95,
			Status:   1,
			Created:  time.Now(),
			Modified: time.Now(),
		},
	}
	common.Success(c,list)
}

type ItemParam struct {
	Id int `json:"id"`
}

//商品详情
func (this *ProductController) GetItem (c *gin.Context)  {
	var IdParam ItemParam
	if err := c.ShouldBindJSON(&IdParam); err != nil {
		common.Failed(c,"param id error",100,nil)
		return
	}
	var p model.Product
	if IdParam.Id ==1 {
		p = model.Product{
			Id:       1,
			Name:     "test",
			Price:    100,
			Status:   1,
			Created:  time.Now(),
			Modified: time.Now(),
		}
	}else {
		p = model.Product{
			Id:       100,
			Name:     "100test",
			Price:    100,
			Status:   1,
			Created:  time.Now(),
			Modified: time.Now(),
		}
	}



	common.Success(c,p)
}