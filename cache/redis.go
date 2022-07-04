package cache

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/stulzq/hexo-statistics/config"
	"github.com/stulzq/hexo-statistics/logger"
)

var cli *redis.Client

func init() {
	var conf RedisConf
	if err := config.GetStruct("redis", &conf); err != nil {
		panic(errors.Wrap(err, "load redis conf err"))
	}

	timout := time.Duration(conf.Timeout) * time.Millisecond
	//opt := &redis.ClusterOptions{Password: conf.Password, Addrs: strings.Split(conf.Address, ","), MaxRetries: 2, ReadTimeout: timout, WriteTimeout: timout, IdleTimeout: timout}
	opt := &redis.Options{Password: conf.Password, Addr: conf.Address, MaxRetries: 2, ReadTimeout: timout, WriteTimeout: timout, IdleTimeout: timout}
	client := redis.NewClient(opt)

	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		logger.Error("[cache] connect to redis err", ping.Err())
	} else {
		logger.Info("connect to redis success: ", conf.Address)
	}

	cli = client
}
