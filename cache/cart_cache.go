package cache

import (
	"fmt"
	"strconv"
)

type CartCache struct {

}

//redis hash 使用
//实现购物车功能

func (c *CartCache) Key (userId int) string {
	return "shoppingCart:userId:" + strconv.Itoa(userId)
}

//以客户id作为key，每位客户创建一个hash存储结构存储对应的购物车信息
//插入一条hash数据
func (c *CartCache) HSet (userId int,goodsId,num int) bool  {
	key := c.Key(userId)
	field := "goodsId:" + strconv.Itoa(goodsId)
	flag,err := RedisClient.HSet(key,field,num).Result()
	if err != nil {
		fmt.Println("redis HSet error:",err)
	}
	fmt.Println("redis HSet result:",flag) //当已经存在的时候 flag 为false
	return flag
}

//查看hash中指定字段是否存在
func (c *CartCache) HExists (userId,goodsId int) bool  {
	key := c.Key(userId)
	field := "goodsId:" + strconv.Itoa(goodsId)
	flag,err := RedisClient.HExists(key,field).Result()
	if err != nil {
		fmt.Println("redis hExists error:",err)
	}
	fmt.Println("redis HExists result:",flag)
	return flag
}

//获取hash中指定字段的值
func (c *CartCache) HGet (userId ,goodsId int)string {
	key := c.Key(userId)
	field := "goodsId:" + strconv.Itoa(goodsId)
	res,err := RedisClient.HGet(key,field).Result()
	if err != nil {
		//当field 不存在时候
		fmt.Println("redis HGet error:",err)
		return ""
	}
	fmt.Println("redis HGet result:",res)
	return res
}

//为哈希表 key 中的指定字段的整数值加上增量 increment
func (c *CartCache) HIncrBy (userId ,goodsId int,step int64) bool {
	key := c.Key(userId)
	field := "goodsId:" + strconv.Itoa(goodsId)
	temp,err := RedisClient.HIncrBy(key,field,step).Result()
	if err != nil {
		fmt.Println("redis HIncrBy error:",err)
		return false
	}
	fmt.Println("redis HIncrBy result:",temp) //temp 为增加后的值
	return true
}




