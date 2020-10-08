package redis

import (
	"GinDemo/config"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func SetUp() {
	pool = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     config.RedisSetting.MaxIdle,
		MaxActive:   config.RedisSetting.MaxActive,
		IdleTimeout: 2 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisSetting.Host)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			// 选择db
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}

func Set(key string, data interface{}, time int) error {
	/**
	Redis String set操作
	*/
	conn := pool.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value, "EX", 100)
	if err != nil {
		return err
	}
	return nil
}

func Expire(key string, time int) error {
	/**
	Redis Expire操作
	*/
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	/**
	Redis String get操作
	*/
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get error:", err)
	}
	return res, nil
}

func Delete(key string) (bool, error) {
	/**
	Redis String delete操作
	*/
	conn := pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

func Lpush(key string, data []string) error {
	/**
	Redis String Lpush操作
	*/
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(data)...)
	if err != nil {
		fmt.Println("redis lpush error:", err)
		return err
	}
	return nil
}

func Hset(key, field, value string) error {
	/**
	Redis String Hset操作
	*/
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", key, field, value)
	if err != nil {
		fmt.Println("redis hset error:", err)
		return err
	}
	return nil
}

func Hget(key, field string) (string, error) {
	/**
	Redis String Hget操作
	*/
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("hget", key, field))
	if err != nil {
		fmt.Println("redis hget error", err)
	}
	return res, err
}
