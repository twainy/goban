package goban

import (
    "github.com/garyburd/redigo/redis"
    "os"
    "encoding/json"
)

var redisPool *redis.Pool
var config Config

type Config struct {
    Server string
    Maxconnection int
}

func Setup(conffilepath string) {
    file,err := os.Open(conffilepath)
    if  err != nil {
        panic(err)
    }
    decoder := json.NewDecoder(file)
    config = Config{}
    err = decoder.Decode(&config)
    if  err != nil {
        panic(err)
    }
}

func start() {
    redisPool = redis.NewPool(func() (redis.Conn, error) {
        c, err := redis.Dial("tcp", config.Server)
        if err != nil {
            return nil, err
        }

        return c, err
    }, config.Maxconnection)
}

func getConn() redis.Conn {
    if redisPool == nil {
        start()
    }
    
    c := redisPool.Get()
    return c
}

func Delete(key string) (error) {
    c := getConn()
    _,err := c.Do("DELETE", key)
    return err
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
func Zadd(key string, value int, name string) error {
    c := getConn()
    _,err := c.Do("ZADD", key, value, name)
    return err
}
func Zrevrank(key string, name string) (int,error) {
    c := getConn()
    val,err := redis.Int(c.Do("ZREVRANK", key, name))
    return val,err
}

