package dao

import (
	"context"
	"fmt"
	"gitee.com/QunXiongZhuLu/KongMing/log"
	"gitee.com/QunXiongZhuLu/kratos/pkg/cache/redis"
	"gitee.com/QunXiongZhuLu/kratos/pkg/conf/paladin"
	"os"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	var (
		cfg redis.Config
		ct  paladin.Map
	)
	if err = paladin.Get("redis.yaml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	cfg.Addr = getRedisAddr()
	r = redis.NewRedis(&cfg)

	if _, err = r.Do(context.Background(), "SET", "ping", "pong"); err != nil {
		log.Error("====== DB Connect Error ======", err)
		return
	}

	cf = func() {
		_ = r.Close()
	}
	return
}

func getRedisAddr() (addr string) {
	var (
		ipStr   string
		portStr string
		ipKey   = "ARSENAL_SVC_SECURITY_HXREDIS_HXREDIS_TCP_IP"
		portKey = "ARSENAL_SVC_SECURITY_HXREDIS_HXREDIS_TCP_PORT"
		hostKey = "ARSENAL_SVC_SECURITY_HXREDIS_HXREDIS_TCP_HOST"
	)
	ipStr = os.Getenv(ipKey)
	portStr = os.Getenv(portKey)

	addr = fmt.Sprintf("%v:%v", ipStr, portStr)
	log.Info("getRedisAddr", ipStr, portStr, ipStr, portStr, addr)
	if ipStr == "" || portStr == "" {
		addr = os.Getenv(hostKey)
	}
	if addr == "" {
		addr = "127.0.0.1:908"
	}
	return
}

func (d *dao) SetRedisSecond(ctx context.Context, key interface{}, value interface{}, expire int64) (err error) {
	if _, err = d.redis.Do(ctx, "SETEX", key, expire, value); err != nil {
		log.TLog.Error("conn.Set(%v) error(%v)", key, err)
		return err
	}
	log.TLog.Info("conn.Set(%v) (%v)", key, value)
	return
}

func (d *dao) GetRedis(ctx context.Context, key interface{}) (value string, err error) {
	var res []byte
	if res, err = redis.Bytes(d.redis.Do(ctx, "GET", key)); err != nil {
		log.TLog.Error("conn.Get(%v) error(%v)", key, err)
		return "", err
	}
	log.TLog.Info("conn.Get(%v): %v", key, res)
	return string(res), nil
}

func (d *dao) DelRedis(ctx context.Context, key interface{}) (err error) {
	_, err = d.redis.Do(ctx, "Del", key)
	log.TLog.Info("Delete Key(%v): error:%v", key, err)
	return err
}

func (d *dao) ExistRedis(ctx context.Context, key interface{}) (value int, err error) {
	value, err = redis.Int(d.redis.Do(ctx, "EXISTS", key))
	log.TLog.Info("Exist Key(%v): error:%v", key, err)
	return
}

func (d *dao) PublishRedis(ctx context.Context, channel string, msg string) (count int, err error) {
	count, err = redis.Int(d.redis.Do(ctx, "PUBLISH", channel, msg))
	log.TLog.Info(fmt.Sprintf("Redis Publish Channel:%v,msg:%v, comsumer count:%v, err:%v", channel, msg, count, err))

	return count, err
}

func (d *dao) AddZSetRedis(ctx context.Context, key string, score int64, mem string) (count int, err error) {
	count, err = redis.Int(d.redis.Do(ctx, "ZADD", key, score, mem))
	return
}

func (d *dao) DelZSetByScore(ctx context.Context, key string, min interface{}, max interface{}) (count int, err error) {
	count, err = redis.Int(d.redis.Do(ctx, "ZREMRANGEBYSCORE", key, min, max))
	if err != nil {
		log.TLog.Error(fmt.Sprintf("DelZSetByScore Error, key:%v,min:%v, max:%v, count:%v,err:%v", key, min, max, count, err))
		return
	}

	log.TLog.Debug(fmt.Sprintf("DelZSetByScore, key:%v,min:%v, max:%v, count:%v,err:%v", key, min, max, count, err))
	return
}

func (d *dao) DelZSetByVal(ctx context.Context, key string, val string) (count int, err error) {
	count, err = redis.Int(d.redis.Do(ctx, "ZREM", key, val))
	return
}

func (d *dao) GetZSetByScore(ctx context.Context, key string, min interface{}, max interface{}) (data []string, err error) {
	data, err = redis.Strings(d.redis.Do(ctx, "zrangebyscore", key, min, max))
	if err != nil {
		log.TLog.Error(fmt.Sprintf("GetZSetByScore Error, key:%v,min:%v, max:%v, data:%v,err:%v", key, min, max, data, err))
		return
	}

	log.TLog.Debug(fmt.Sprintf("GetZSetByScore, key:%v,min:%v, max:%v, data:%v,err:%v", key, min, max, data, err))
	return
}
