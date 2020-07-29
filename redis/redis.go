/**
 * 基于 redigo 做了部分封装, 仍可直接使用原 redigo 命令
 */

package redis

import (
	"encoding/json"
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	*redis.Pool
	options Options
}

type Options struct {
	Network     string                                 // 连接协议, 默认 tcp
	Addr        string                                 // 连接地址, 默认 127.0.0.1:6379
	Password    string                                 // 鉴权密码, 默认空
	Db          int                                    // 连接数据库, 默认 0
	MaxActive   int                                    // 最大活动连接数, 0 为不限制, 默认 runtime.NumCPU() * 10
	MaxIdle     int                                    // 最大空闲连接数, 0 为不限制, 默认 runtime.NumCPU() * 100
	IdleTimeout int                                    // 空闲连接的超时时间, 单位秒, 默认 300
	Prefix      string                                 // 键名前缀
	Marshal     func(v interface{}) ([]byte, error)    // 序列化函数, 默认 json.Marshal
	Unmarshal   func(data []byte, v interface{}) error // 反序列化函数, 默认 json.Unmarshal
}

func New(options Options) (*Redis, error) {
	if options.Network == "" {
		options.Network = "tcp"
	}
	if options.Addr == "" {
		options.Addr = "127.0.0.1:6379"
	}
	if options.MaxActive == 0 {
		options.MaxActive = runtime.NumCPU() * 10
	}
	if options.MaxIdle == 0 {
		options.MaxIdle = runtime.NumCPU() * 100
	}
	if options.IdleTimeout == 0 {
		options.IdleTimeout = 300
	}
	if options.Marshal == nil {
		options.Marshal = json.Marshal
	}
	if options.Unmarshal == nil {
		options.Unmarshal = json.Unmarshal
	}

	r := &Redis{
		&redis.Pool{
			MaxActive:   options.MaxActive,
			MaxIdle:     options.MaxIdle,
			IdleTimeout: time.Duration(options.IdleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial(options.Network, options.Addr)
				if err != nil {
					return nil, err
				}
				if options.Password != "" {
					if _, err := conn.Do("AUTH", options.Password); err != nil {
						conn.Close()
						return nil, err
					}
				}
				if _, err := conn.Do("SELECT", options.Db); err != nil {
					conn.Close()
					return nil, err
				}
				return conn, err
			},
			TestOnBorrow: func(conn redis.Conn, t time.Time) error {
				_, err := conn.Do("PING")
				return err
			},
		},
		options,
	}

	conn, err := r.Dial()
	if err != nil {
		return nil, err
	}
	conn.Close()

	return r, nil
}

/**
 * Inner
 */

func (r *Redis) do(command string, args ...interface{}) (interface{}, error) {
	conn := r.Get()
	err := conn.Err()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.Do(command, args...)
}

func (r *Redis) prefix(key string) string {
	return r.options.Prefix + key
}

func (r *Redis) encode(val interface{}) (interface{}, error) {
	var result interface{}
	switch v := val.(type) {
	case string, int, uint, int8, int16, int32, int64, float32, float64, bool:
		result = v
	default:
		b, err := r.options.Marshal(v)
		if err != nil {
			return nil, err
		}
		result = string(b)
	}
	return result, nil
}

func (r *Redis) decode(str interface{}, err error, val interface{}) error {
	return r.options.Unmarshal([]byte(str.(string)), val)
}

/**
 * Common
 */

func (r *Redis) SELECT(db int) error {
	_, err := r.do("SELECT", db)
	return err
}

func (r *Redis) FLUSHDB() error {
	_, err := r.do("FLUSHDB")
	return err
}

func (r *Redis) GET(key string) (string, error) {
	return redis.String(r.do("GET", r.prefix(key)))
}

func (r *Redis) SET(key string, val interface{}, expire int64) error {
	value, err := r.encode(val)
	if err != nil {
		return err
	}
	if expire > 0 {
		_, err := r.do("SETEX", r.prefix(key), expire, value)
		return err
	}
	_, err = r.do("SET", r.prefix(key), value)
	return err
}

func (r *Redis) EXISTS(key string) bool {
	resp, err := redis.Bool(r.do("EXISTS", r.prefix(key)))
	return err == nil && resp
}

func (r *Redis) DEL(key string) error {
	_, err := r.do("DEL", r.prefix(key))
	return err
}

func (r *Redis) TTL(key string) (ttl int64, err error) {
	return redis.Int64(r.do("TTL", r.prefix(key)))
}

func (r *Redis) EXPIRE(key string, expire int64) error {
	_, err := r.do("EXPIRE", r.prefix(key), expire)
	return err
}

/**
 * String
 */

func (r *Redis) INCR(key string) (val int64, err error) {
	return redis.Int64(r.do("INCR", r.prefix(key)))
}

func (r *Redis) INCRBY(key string, amount int64) (val int64, err error) {
	return redis.Int64(r.do("INCRBY", r.prefix(key), amount))
}

func (r *Redis) DECR(key string) (val int64, err error) {
	return redis.Int64(r.do("DECR", r.prefix(key)))
}

func (r *Redis) DECRBY(key string, amount int64) (val int64, err error) {
	return redis.Int64(r.do("DECRBY", r.prefix(key), amount))
}

/**
 * List
 */

func (r *Redis) LPUSH(key string, val string) error {
	_, err := r.do("LPUSH", r.prefix(key), val)
	return err
}

func (r *Redis) RPUSH(key string, val string) error {
	_, err := r.do("RPUSH", r.prefix(key), val)
	return err
}

func (r *Redis) LPOP(key string) (string, error) {
	return redis.String(r.do("LPOP", r.prefix(key)))
}

func (r *Redis) RPOP(key string) (string, error) {
	return redis.String(r.do("RPOP", r.prefix(key)))
}

func (r *Redis) LLEN(key string) int64 {
	resp, err := redis.Int64(r.do("LLEN", r.prefix(key)))
	if err != nil {
		return 0
	}
	return resp
}

// todo: hash, set, sorted set
