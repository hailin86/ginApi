package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

type LikeCache struct {

}

//redis zset 使用

func (c *LikeCache) Key() string  {
	return "product:like"
}

//zset 添加一个成员 member score
func (c *LikeCache) ZAdd (score float64,id int) bool {
	key := c.Key()
	temp := redis.Z{
		Score:  score,
		Member: id,
	}
	_,err :=RedisClient.ZAdd(key,temp).Result()
	if err != nil {
		fmt.Println("redis ZAdd error:",err)
		return false
	}
	return true
}

func (c *LikeCache) ZRem (id int) bool {
	key := c.Key()
	temp,err := RedisClient.ZRem(key,id).Result()
	if err != nil {
		fmt.Println("redis ZRem error",err)
		return false
	}
	fmt.Println("redis ZRem result:",temp) // temp 返回的是 1 / 0
	return true
}

//zset 成员 score 增加1
func (c *LikeCache) ZIncrBy (id int) bool {
	key := c.Key()
	idStr := strconv.Itoa(id)
	_,err :=RedisClient.ZIncrBy(key,1,idStr).Result()
	if err != nil {
		fmt.Println("redis ZIncrBy error:",err)
		return false
	}
	return true
}

//返回zset 中成员的分数
func (c *LikeCache) ZScore (id int) (float64,error) {
	key := c.Key()
	idStr := strconv.Itoa(id)
	score,err :=RedisClient.ZScore(key,idStr).Result()
	if err != nil {
		fmt.Println("redis ZScore error:",err)
		return 0,err
	}
	return score,nil
}
