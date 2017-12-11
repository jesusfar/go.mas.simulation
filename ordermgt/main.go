package main

import (
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jesusfar/go.mas.simulation/environment"
	"github.com/jesusfar/go.mas.simulation/ordermgt/order"
	"os"
)

func healthCheckController(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func main() {
	var g inject.Graph

	var env environment.Env
	var controller order.Controller
	var service order.Service
	var orderRepository order.RedisRepository
	var redisPool *redis.Pool

	redisPool = env.SetupDataBaseConnPool()

	err := g.Provide(
		&inject.Object{Value: &controller},
		&inject.Object{Value: &service},
		&inject.Object{Value: &orderRepository},
		&inject.Object{Value: redisPool},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/health", healthCheckController)

	r.POST("/order", controller.CreateOrderController)
	r.GET("/order/:id", controller.GetOrderById)

	r.Run(":9000")
}
