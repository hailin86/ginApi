package cache

import "fmt"

type BlackListCache struct {

}

//redis  set 的 使用 set 查找 添加 删除的复杂度都是 O(1)

//模拟一个ip黑名单 使用redis的set数据类型 ，因为set集合 成员不可以重复
func (c *BlackListCache) Key() string {
	return "BlackList:ip"
}

//set 中添加一个成员
func (c *BlackListCache) SAdd(ip string) bool  {
	key := c.Key()
	_,err :=RedisClient.SAdd(key,ip).Result()
	if err != nil {
		fmt.Println("redis SAdd error:",err)
		return false
	}
	return true
}

//set 中删除一个成员
func (c *BlackListCache) SRem (ip string) bool {
	key := c.Key()
	_,err :=RedisClient.SRem(key,ip).Result()
	if err != nil {
		fmt.Println("redis SRem error:",err)
		return false
	}
	return true

}

//set 中是否存在某个成员
func (c *BlackListCache) SIsMember(ip string) bool  {
	key := c.Key()
	_,err := RedisClient.SIsMember(key,ip).Result()
	if err != nil {
		fmt.Println("redis SAdd error:",err)
		return false
	}
	return true
}
