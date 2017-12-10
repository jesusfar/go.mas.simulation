package order

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type OrderRepository interface {
	save(order order)
	getById(orderId string) *order
}

type RedisRepository struct {
	DBPoolConn *redis.Pool `inject:""`
}

func (repository RedisRepository) save(order *order) error {

	conn := repository.DBPoolConn.Get()
	defer conn.Close()

	serialized, err := json.Marshal(order)

	if err != nil {
		return err
	}

	_, err = conn.Do("SET", order.Id, serialized)

	return err
}

func (repository RedisRepository) getById(orderId string) (*order, error) {
	conn := repository.DBPoolConn.Get()
	defer conn.Close()

	var order order

	res, err := redis.Bytes(conn.Do("GET", orderId))

	err = json.Unmarshal(res, &order)

	return &order, err
}
