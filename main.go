package goban

import (
    "github.com/garyburd/redigo/redis"
    "flag"
)

var redisPool *redis.Pool

var (
    redisAddress = flag.String("redis-address", ":6379", "Address to the Redis server")
    maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

func start() {
    flag.Parse()
    redisPool = redis.NewPool(func() (redis.Conn, error) {
        c, err := redis.Dial("tcp", *redisAddress)
        if err != nil {
            return nil, err
        }

        return c, err
    }, *maxConnections)
}

func getConn() redis.Conn {
    if redisPool == nil {
        start()
    }
    
    c := redisPool.Get()
    return c
}

func Get(key string) (string, error) {
    c := getConn()
    v,err := redis.String(c.Do("GET", key))
    return v,err
}

func Set(key string, value string) error {
    c := getConn()
    _,err := c.Do("SET", key, value)
    return err
}
