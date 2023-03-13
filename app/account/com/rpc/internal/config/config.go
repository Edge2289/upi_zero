package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Db struct {
		Host string
		User string
		Password string
		DbName string
		DeBug int
		MaxLifetime int
		MaxOpenConn int
		MaxIdleConn int
	}

	Cache cache.CacheConf
}
