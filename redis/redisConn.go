package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"flag"
	"time"
)

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", "180.76.111.158:6379", "")
)

func initPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func getRedisConn() redis.Conn {

	if pool == nil {
		conn := initPool(*redisServer)
		pool = conn
	}
	return pool.Get()
}

func doRedis(i int) {
	flag.Parse()
	conn := getRedisConn()
	defer conn.Close()
	str := fmt.Sprintf("%s,%d", "ileir_", i)
	_, err := conn.Do("SET", "username", str)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(conn.Do("GET", "username"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Got username %v \n", username)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		doRedis(i)
	}
}
