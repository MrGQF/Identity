package dao

import (
	"context"
	"gitee.com/QunXiongZhuLu/kratos/pkg/cache/redis"
)

type dao struct {
	redis *redis.Redis
}

type DaoImpl interface {
	// redis
	SetRedisSecond(ctx context.Context, key interface{}, value interface{}, expire int64) (err error)
	GetRedis(ctx context.Context, key interface{}) (value string, err error)
	DelRedis(ctx context.Context, key interface{}) (err error)
	ExistRedis(ctx context.Context, key interface{}) (value int, err error)

	// ZSet
	AddZSetRedis(ctx context.Context, key string, score int64, mem string) (count int, err error)
	DelZSetByScore(ctx context.Context, key string, min interface{}, max interface{}) (count int, err error)
	DelZSetByVal(ctx context.Context, key string, val string) (count int, err error)
	GetZSetByScore(ctx context.Context, key string, min interface{}, max interface{}) (data []string, err error)
}

func New(r *redis.Redis) (d DaoImpl, cf func(), err error) {
	return NewDao(r)
}

func NewDao(r *redis.Redis) (d *dao, cf func(), err error) {

	d = &dao{
		redis: r,
	}
	cf = func() {}

	return
}
