package api

import (
	"ginApi/test"
	"testing"
)

func TestProduct(t *testing.T)  {

	var caseList  = []test.TestCase{
		{
			Method:      "POST",
			Url:         "/api/api/product/getItem",
			ContentType: "application/json",
			Param:       "",
			Code:        1100,
			Desc:        "api商品详情不传id",
			ShowBody:    true,
			ErrMsg:      "",
			Ext1:        nil,
			Ext2:        nil,
		},
		{
			Method:      "POST",
			Url:         "/api/api/product/getItem",
			ContentType: "application/json",
			Param:      `{"id":1}`,
			Code:        200,
			Desc:        "api商品详情传id",
			ShowBody:    true,
			ErrMsg:      "",
			Ext1:        nil,
			Ext2:        nil,
		},
	}

	test.ExecApi(t,caseList)
}
