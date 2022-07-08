package test

import (
	"bytes"
	"crypto/sha256"
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
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

type TestCase struct {
	Method       string      //请求类型
	Url          string      //链接
	ContentType string      //
	Param        string      //参数
	Code         float64         //状态码
	Desc         string      //描述
	ShowBody     bool        //是否展示返回
	ErrMsg       string      //错误信息
	Ext1         interface{} //自定义1
	Ext2         interface{} //自定义2
}




var route *gin.Engine


func init()  {
	path ,_ := os.Getwd()
	configPath := path + "/../../conf"
	//fmt.Println("conf path:",configPath)
	config := common.InitConf(configPath)
	//跟main.go 保持一致，用到什么就初始化什么
	//初始日志
	//fmt.Println(config.LoggerConf.Filename)
	common.InitLogger(config.LoggerConf)

	route = gin.New()
	route.Use(middleware.GinLogger(),middleware.GinRecovery(false))
	route.Use(middleware.Cors())

	router.InitRouter(route)
}

func NewBufferString(body string) io.Reader {
	return bytes.NewBufferString(body)
}

func DoRequest(method, url, contentType string, body string) (c *gin.Context, r *http.Request, w *httptest.ResponseRecorder) {
	if method =="" {
		method = "POST" //因为我的请求大部分都是post 这个为了偷懒
	}
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	r = httptest.NewRequest(method, url, NewBufferString(body))
	c.Request = r
	c.Request.Header.Set("Content-Type", contentType)
	c.Request.Header.Set("GINAPPKEY","100000")
	//按照自己需求设置自己的 header
	//...

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

func DoApiRequest(method, url, contentType string, body string) (c *gin.Context, r *http.Request, w *httptest.ResponseRecorder) {

	if method =="" {
		method = "POST" //因为我的请求大部分都是post 这个为了偷懒
	}
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	//fmt.Println("body:",body)
	r = httptest.NewRequest(method, url, NewBufferString(body))
	c.Request = r
	c.Request.Header.Set("Content-Type", contentType)
	route.ServeHTTP(w, r)
	return
}

func BuildApiParam (testCase *TestCase) {
	var appKey = "100000"
	var appSecret = "123456"
	var params = make(map[string]interface{})

	_ = json.Unmarshal([]byte(testCase.Param),&params)

	//fmt.Println("p1:",params)

	params["app_key"] = appKey
	params["timestamp"] =strconv.FormatInt(time.Now().Unix(), 10)
	params["nonce_str"] ="随机字符串"

	//fmt.Println("p2:",params)

	sign := GenerateSign(appSecret,params)
	params["sign"] = sign

	//fmt.Println("p3:",params)

	tempBytes,_ := json.Marshal(params)
	testCase.Param = string(tempBytes)

	//fmt.Println("p4:",testCase.Param)
}

func ExecApi (t *testing.T,caseList []TestCase) {

	for k, v := range caseList {
		BuildApiParam(&v)

		_, _, w := DoApiRequest(v.Method, v.Url, v.ContentType, v.Param)
		fmt.Println()
		fmt.Printf("第%d个测试用例:%s;", k+1, v.Desc)
		if v.ShowBody {
			fmt.Printf("接口返回%s", w.Body.String())
			fmt.Println()
		}
		//assert.Equal(t, 200, w.Code)
		result := Result(w.Body.String())
		//fmt.Println("result:", result)
		assert.Equal(t, v.Code, result["code"])
	}
}


func GenerateSign(appSecret string, params map[string]interface{}) string  {
	var keys []string
	for key,_:= range params {
		keys = append(keys,key)
	}

	sort.Strings(keys)

	signStr := ""
	for _,key := range keys {
		value := strings.TrimSpace(fmt.Sprintf("%v", params[key]))
		if value != "" &&  key != "sign" {
			signStr += key + "=" + value + "&"
		}
	}

	signStr += "secretKey="+appSecret

	h := sha256.New()
	h.Write([]byte(signStr))
	return strings.ToUpper(fmt.Sprintf("%x",h.Sum(nil)))
}


func Result(body string) gin.H  {
	var result gin.H
	_ = json.Unmarshal([]byte(body),&result)
	return result
}