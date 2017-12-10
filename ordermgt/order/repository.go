package order

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type OrderRepository interface {
	save(order order)
}

type RedisRepository struct {
	DBPoolConn *redis.Pool `inject:""`
}

func (repository RedisRepository) save(order order)  (error) {

	conn := repository.DBPoolConn.Get()
	defer conn.Close()

	res, err := conn.Do("SET", order.id, order)

	fmt.Println(res)
	fmt.Println(order)
	return err
}
