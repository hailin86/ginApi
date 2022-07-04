package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ginApi/common"
	"ginApi/middleware"
	"ginApi/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	Method       string      //请求类型
	Url          string      //链接
	ContentType string      //
	Param        string      //参数
	Code         int         //状态码
	Desc         string      //描述
	ShowBody     bool        //是否展示返回
	ErrMsg       string      //错误信息
	Ext1         interface{} //自定义1
	Ext2         interface{} //自定义2
}

func init()  {
	common.InitLogger()
}

func NewBufferString(body string) io.Reader {
	return bytes.NewBufferString(body)
}

func DoRequest(method, url, contentType string, body string) (c *gin.Context, r *http.Request, w *httptest.ResponseRecorder) {

	route := gin.New()
	route.Use(middleware.GinLogger(),middleware.GinRecovery(false))
	route.Use(middleware.Cors())
	router.InitRouter(route)

	if method =="" {
		method = "POST" //因为我的请求大部分都是post 这个为了偷懒
	}
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	r = httptest.NewRequest(method, url, NewBufferString(body))
	c.Request = r
	c.Request.Header.Set("Content-Type", contentType)
	c.Request.Header.Set("GINAPPKEY","100000")
	route.ServeHTTP(w, r)
	return
}

func Exec (t *testing.T,caseList []TestCase) {
	for k, v := range caseList {
		_, _, w := DoRequest(v.Method, v.Url, v.ContentType, v.Param)
		fmt.Println()
		fmt.Printf("第%d个测试用例:%s;", k+1, v.Desc)
		if v.ShowBody {
			fmt.Printf("接口返回%s", w.Body.String())
			fmt.Println()
		}
		assert.Equal(t, 200, w.Code)
		//result := Result(w.Body.String())
		//fmt.Println("接口返回:", w.Body.String())
		//assert.Equal(t, "success", result["msg"])
	}
}



func Result(body string) gin.H  {
	var result gin.H
	json.Unmarshal([]byte(body),&result)
	return result
}