package dao

import (
	"context"
	"github.com/oofpgDLD/ctpf/imp-logic/conf"
	"github.com/gomodule/redigo/redis"
	"time"
)

// Dao dao.
type Dao struct {
	c           *conf.Config
	redis       *redis.Pool
	redisExpire int32
}

// New new a dao and return.
func New(c *conf.Config) *Dao {
	d := &Dao{
		c:           c,
		redis:       newRedis(c.Redis),
		redisExpire: int32(time.Duration(c.Redis.Expire) / time.Second),
	}
	return d
}

func newRedis(c *conf.Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.Idle,
		MaxActive:   c.Active,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Duration(c.WriteTimeout)),
				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

// Close close the resource.
func (d *Dao) Close() error {
	return d.redis.Close()
}

// Ping dao ping.
func (d *Dao) Ping(c context.Context) error {
	return d.pingRedis(c)
}