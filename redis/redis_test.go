package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRedis() (*Redis, error) {
	r, err := New(Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		Prefix:   "higo_",
	})
	return r, err
}

func TestCommon(t *testing.T) {
	r, err := getRedis()
	assert.Nil(t, err)
	if err != nil {
		t.FailNow()
	}
	// redis.Pool 方法 Stats, 可直接调用
	assert.NotNil(t, r.Stats())

	// select 9
	assert.Nil(t, r.SELECT(9))

	key, val, expire := "hello", "world", int64(1000)

	// set
	assert.Nil(t, r.SET(key, val, expire))

	// get
	getResp, getErr := r.GET(key)
	assert.Nil(t, getErr)
	assert.Equal(t, val, getResp)

	// exist
	assert.True(t, r.EXISTS(key))

	// expire
	assert.Nil(t, r.EXPIRE(key, int64(2000)))

	// ttl
	ttlResp, ttlErr := r.TTL(key)
	assert.Nil(t, ttlErr)
	assert.LessOrEqual(t, int64(0), ttlResp)
	assert.GreaterOrEqual(t, int64(2000), ttlResp)

	// del
	assert.Nil(t, r.DEL(key))
	assert.False(t, r.EXISTS(key))

	//flushdb
	assert.Nil(t, r.FLUSHDB())
}

func TestString(t *testing.T) {
	r, err := getRedis()
	assert.Nil(t, err)
	if err != nil {
		t.FailNow()
	}
	assert.Nil(t, r.SELECT(9))

	key, val, expire := "incr", 1, int64(1000)
	assert.Nil(t, r.SET(key, val, expire))

	// incr
	incrResp, incrErr := r.INCR(key)
	assert.Nil(t, incrErr)
	assert.Equal(t, int64(2), incrResp)

	// incrby
	incrByResp, incrByErr := r.INCRBY(key, 10)
	assert.Nil(t, incrByErr)
	assert.Equal(t, int64(12), incrByResp)

	// decr
	decrResp, decrErr := r.DECR(key)
	assert.Nil(t, decrErr)
	assert.Equal(t, int64(11), decrResp)

	// incrby
	decrByResp, decrByErr := r.DECRBY(key, 10)
	assert.Nil(t, decrByErr)
	assert.Equal(t, int64(1), decrByResp)

	assert.Nil(t, r.DEL(key))
}

func TestList(t *testing.T) {
	r, err := getRedis()
	assert.Nil(t, err)
	if err != nil {
		t.FailNow()
	}
	assert.Nil(t, r.SELECT(9))

	key := "list"

	// push
	assert.Nil(t, r.LPUSH(key, "1"))
	assert.Equal(t, int64(1), r.LLEN(key))
	assert.Nil(t, r.RPUSH(key, "2"))
	assert.Equal(t, int64(2), r.LLEN(key))
	assert.Nil(t, r.RPUSH(key, "3"))
	assert.Nil(t, r.LPUSH(key, "0"))
	assert.Equal(t, int64(4), r.LLEN(key))

	// pop
	lpopResp, lpopErr := r.LPOP(key)
	assert.Nil(t, lpopErr)
	assert.Equal(t, "0", lpopResp)
	rpopResp, rpopErr := r.RPOP(key)
	assert.Nil(t, rpopErr)
	assert.Equal(t, "3", rpopResp)

	assert.Nil(t, r.DEL(key))
}

func TestCoverage(t *testing.T) {
	_, err1 := New(Options{
		Addr: "127.0.0.1:63790",
	})
	assert.NotNil(t, err1)

	_, err2 := New(Options{
		Password: "1234567",
	})
	assert.NotNil(t, err2)

	_, err3 := New(Options{
		Password: "123456",
		Db:       16,
	})
	assert.NotNil(t, err3)

	r, err := New(Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		Prefix:   "higo_",
	})
	assert.Nil(t, err)
	if err != nil {
		t.FailNow()
	}
	assert.Nil(t, r.SELECT(9))
	assert.Nil(t, r.SET("encode", []int{1, 2, 3, 4}, 0))
	assert.Empty(t, r.LLEN("queue"))
	// assert.Nil(t, r.FLUSHDB())
}
