package main

import (
	"fmt"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		err         error
		ctx         = gctx.New()
		cache       = gcache.New()
		redisConfig = &gredis.Config{
			Address: "127.0.0.1:15001",
			Pass:    "123456",
			Db:      9,
		}
		cacheKey   = `key`
		cacheValue = `value`
	)
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	// Set and Get using cache object.
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// Get using redis client.
	fmt.Println(redis.MustDo(ctx, "GET", cacheKey).String())
}
