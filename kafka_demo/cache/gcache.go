package cache

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

var Cache gcache.Cache

const CachSize = 100

func InitGCache() {
	Cache = gcache.New(CachSize).LRU().Build()
	fmt.Println("Init gcache done!")
}

func Test() {
	Cache.SetWithExpire("1", "value 1", time.Second*2)
	value, err := Cache.Get("1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	time.Sleep(time.Second * 3)
	value, err = Cache.Get("1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
