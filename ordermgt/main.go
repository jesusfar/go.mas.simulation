package main

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jesusfar/go.mas.simulation/ordermgt/order"
)

var (
	redisAddress = flag.String("redis-address", ":6379", "Redis address")
	maxConn      = flag.Int("max-conn", 10, "Max connection")
)

func getRedisPool() *redis.Pool {
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

func healthCheckController(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func main() {
	redisPool := getRedisPool()
	defer redisPool.Close()

	r := gin.Default()

	r.GET("/health", healthCheckController)

	r.POST("/order", order.CreateOrderController)

	r.Run(":9000")
}
