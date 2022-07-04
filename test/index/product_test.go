package index

import (
	"ginApi/test"
	"testing"
)

func TestProduct(t *testing.T)  {

	var caseList  = []test.TestCase{
		{
			Method:      "POST",
			Url:         "/api/index/product/getItems",
			ContentType: "application/json",
			Param:      "",
			Code:        200,
			Desc:        "商品列表接口测试",
			ShowBody:    true,
			ErrMsg:      "",
			Ext1:        nil,
			Ext2:        nil,
		},
		{
			Method:      "POST",
			Url:         "/api/index/product/getItem",
			ContentType: "application/json",
			Param:      `{"id":1}`,
			Code:        200,
			Desc:        "商品详情接口测试",
			ShowBody:    true,
			ErrMsg:      "",
			Ext1:        nil,
			Ext2:        nil,
		},
	}



	test.Exec(t,caseList)


	//body := `{"id":1}`
	//_,_,w := test.DoRequest("","/api/index/product/getItem","application/json",body)
	//
	//assert.Equal(t, 200, w.Code)
	//
	//result := test.Result(w.Body.String())
	//fmt.Println("接口返回:", w.Body.String())
	//assert.Equal(t, "success", result["msg"])


}
