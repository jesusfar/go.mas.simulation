package main

import (
	"github.com/jesusfar/go.mas/agent"
	"github.com/garyburd/redigo/redis"
	"github.com/jesusfar/go.mas.simulation/environment"
)

func main() {
	var env environment.Env
	var redisPool *redis.Pool

	redisPool = env.SetupDataBaseConnPool()

	connPubSub := redis.PubSubConn{Conn: redisPool.Get()}

	agent := agent.New("Inventory-agent", &connPubSub)

	agent.Run()
}
