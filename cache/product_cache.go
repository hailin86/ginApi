package cache

import (
	"encoding/json"
	"fmt"
	"ginApi/model"
	"strconv"
	"time"
)

type ProductCache struct {

}

// redis string 使用

func (c *ProductCache) Key(id int) string {
	return "product:pk:id" + strconv.Itoa(id)
}

func (c *ProductCache) Set(id int,model *model.Product,expire time.Duration) bool {
	key := c.Key(id)
	bytes ,_ := json.Marshal(model)
	_,err := RedisClient.Set(key,bytes,120*time.Second).Result()
	if err != nil {
		fmt.Println("redis set error:",err)
		return false
	}
	return true
}

func (c *ProductCache) Get (id int) *model.Product {
	key := c.Key(id)
	res,err := RedisClient.Get(key).Result()
	if err != nil {
		return nil
	}
	var data *model.Product
	_ = json.Unmarshal([]byte(res),&data)
	return data
}

func (c *ProductCache) Del (id int) bool  {
	_,err := RedisClient.Del(c.Key(id)).Result()
	if err != nil {
		return false
	}
	return  true
}