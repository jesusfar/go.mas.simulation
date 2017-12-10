package environment

import (
	"flag"
	"github.com/garyburd/redigo/redis"
)

var (
	redisAddress = flag.String("redis-address", ":6379", "Redis address")
	maxConn      = flag.Int("max-conn", 10, "Max connection")
)

type Env struct {
	dbConnPool interface{}
}

type DataBaseInterface interface {
	SetupDataBaseConnPool() interface{}
}

func (e *Env) SetupDataBaseConnPool() *redis.Pool {
	flag.Parse()

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}
		return c, err
	}, *maxConn)

	return redisPool
}