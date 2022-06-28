package cache

import (
	"fmt"
	"strconv"
)

type CommentCache struct {

}

//redis list 使用
//使用list 实现 一个商品的最新评论

func (c *CommentCache) Key(id int) string  {
	return "comment:product:id:" + strconv.Itoa(id)
}

//list 头部插入数据
func (c *CommentCache) LPush (id,commentId int) bool {
	key := c.Key(id)
	temp,err := RedisClient.LPush(key,commentId).Result()
	if err != nil {
		fmt.Println("redis lpush error:",err)
		return false
	}
	fmt.Println("redis lpush result:",temp)
	return true
}

//获取list 最新头部插入的 前 n 条数据
func (c *CommentCache ) LRange(id int,start,stop int64) []string  {
	key := c.Key(id)
	temp,err := RedisClient.LRange(key,start,stop).Result()
	if err != nil {
		fmt.Println("redis lrange error:",err)
		return nil
	}
	return temp
}