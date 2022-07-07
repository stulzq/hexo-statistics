package uv_renew

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stulzq/hexo-statistics/cache"
	"github.com/stulzq/hexo-statistics/config"
	"github.com/stulzq/hexo-statistics/keys"
	"github.com/stulzq/hexo-statistics/logger"
)

func Work() {
	// get old uv
	cli := cache.GetClient()
	data := config.Get("statistics:site").([]interface{})

	for i := 0; i < len(data); i++ {
		ctx := context.TODO()
		host := data[i].(string)
		k := keys.GenSiteUv(host)
		ak := keys.GenSiteUvArchive(host)

		pipe := cli.Pipeline()
		res := pipe.PFCount(ctx, k)
		archiveRes := pipe.Get(context.TODO(), ak)

		if _, err := pipe.Exec(ctx); err != nil && err != redis.Nil {
			logger.Errorf("[Job][UV_RENEW][Exec] host %s process failed, err: %v", host, err)
			continue
		}

		var currentUV, archiveUV int64
		currentUV = res.Val()
		if archiveRes.Err() == nil {
			archiveUV, _ = archiveRes.Int64()
		}

		archiveUV += currentUV
		pipe = cli.Pipeline()
		pipe.Del(ctx, k)
		pipe.Set(ctx, ak, archiveUV, 0)
		if _, err := pipe.Exec(ctx); err != nil && err != redis.Nil {
			logger.Errorf("[Job][UV_RENEW][Exec] store new value failed, host %s, err: %v", host, err)
		}

		logger.Infof("[Job][UV_RENEW][Exec] host %s archive uv %d", host, archiveUV)
	}

}
