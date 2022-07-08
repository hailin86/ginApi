package common

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
)

func Success(c *gin.Context,data interface{})  {
	Failed(c,"success",200,data)
}

func Failed(c *gin.Context,msg string,code int,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"msg":msg,
		"code":code,
		"data":data,
		"trace_id":c.GetString(TraceKey),
	})
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
