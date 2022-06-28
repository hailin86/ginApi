package index

import (
	"fmt"
	"ginApi/cache"
	"ginApi/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {

}

func (this *TestController) Test(c *gin.Context)  {
	common.Success(c,"我是test啊")
	return



	//hash 操作
	cc := cache.CartCache{}
	//cc.HSet(1,1,3)
	//cc.HSet(1,2,1)
	//cc.HSet(1,3,3)
	//cc.HSet(2,4,6)
	//cc.HSet(2,5,2)
	//cc.HSet(2,6,4)
	flag1 := cc.HExists(1,1)
	fmt.Println("flag1:",flag1) //true
	flag2 :=cc.HExists(1,5)
	fmt.Println("flag2:",flag2) //false
	res1 := cc.HGet(1,1)
	fmt.Println("res1:",res1) //false
	res2 := cc.HGet(1,6)
	fmt.Println("res2:",res2) //false

	flag3 := cc.HIncrBy(1,1,8)
	fmt.Println("flag3:",flag3) //false


	//list 操作
	//cc := &cache.CommentCache{}
	//for i := 1;i <= 20; i++ {
	//	flag := cc.LPush(1,i)
	//	fmt.Println("lpush res:",flag)
	//}
	//res := cc.LRange(1,0,10)
	//
	//fmt.Println("redis lrange result:",res)


	//ZSet 操作
	//lc := cache.LikeCache{}
	//flag := lc.ZAdd(95,2)
	//fmt.Println("zadd flag:",flag)
	//for i:= 1;i <= 10;i++ {
	//	flag2 := lc.ZIncrBy(2)
	//	fmt.Println("zIncrBy flag:",flag2)
	//}
	//
	//s1,err := lc.ZScore(1)
	//if err == nil {
	//	fmt.Println("s1:",s1)
	//}
	//s2,err := lc.ZScore(2)
	//if err == nil {
	//	fmt.Println("s2:",s2)
	//}
	//s3,err := lc.ZScore(3)
	//if err == nil {
	//	fmt.Println("s3:",s3)
	//}
	//
	//flag4 := lc.ZRem(1)
	//fmt.Println("flag4:",flag4)
	//flag5 := lc.ZRem(3)
	//fmt.Println("flag5:",flag5)










	// Set 操作
	//ipC :=cache.BlackListCache{}
	//for i:=1;i< 10;i++ {
	//	flag := ipC.SAdd( "123.124.145.2"+strconv.Itoa(i))
	//	fmt.Println(i,"---flag:",flag)
	//}
	//ip := "123.124.145.24"
	//
	////检查ip是否存在
	//flag1 := ipC.SIsMember(ip)
	//fmt.Println("是否存在:",flag1)
	////删除 该成员
	//flag2 := ipC.SRem(ip)
	//fmt.Println("del flag:",flag2)

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":nil,
	})

}
