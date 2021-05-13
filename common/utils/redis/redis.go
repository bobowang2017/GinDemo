package redis

import (
	"GinDemo/common/utils/log"
	"GinDemo/config"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
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
			c, err := redis.Dial("tcp", config.RedisSetting.Host,
				redis.DialConnectTimeout(config.RedisSetting.ConnectTimeout*time.Millisecond),
				redis.DialReadTimeout(config.RedisSetting.ReadTimeout*time.Millisecond),
				redis.DialWriteTimeout(config.RedisSetting.WriteTimeout*time.Millisecond))
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
		log.Logger.Error(err.Error(), "| Redis Marshal Error")
		return errors.New(err.Error())
	}
	if time == -1 {
		_, err = conn.Do("SET", key, value)
	} else {
		_, err = conn.Do("SET", key, value, "EX", time)
	}
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis Set Error")
		return errors.New(err.Error())
	}
	return nil
}

func SetNx(key string, data interface{}, time int) error {
	/**
	Redis String setnx操作
	*/
	conn := pool.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = redis.String(conn.Do("SET", key, value, "EX", time, "NX"))
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis SetNx Error")
		return errors.New(err.Error())
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
		log.Logger.Error(err.Error(), "| Redis Expire Error")
		return errors.WithMessage(err, "| Redis Expire Error")
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
		log.Logger.Error(err.Error(), "| Redis Get Error")
	}
	return res, nil
}

func Delete(key string) (bool, error) {
	/**
	Redis String delete操作
	*/
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.Bool(conn.Do("DEL", key))
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis Delete Error")
	}
	return res, err
}

func LPush(key string, data []string) error {
	/**
	Redis String LPush操作
	*/
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(data)...)
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis LPush Error")
		return errors.WithMessage(err, "| Redis LPush Error")
	}
	return nil
}

func HSet(key, field, value string) error {
	/**
	Redis String Hset操作
	*/
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", key, field, value)
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis HSet Error")
		return errors.WithMessage(err, "| Redis HSet Error")
	}
	return nil
}

func HGetAll(key string) (map[string]string, error) {
	conn := pool.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("hgetall", key))
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis HGetAll Error")
		return nil, err
	}
	data := make(map[string]string)
	for i := 0; i < len(result); i += 2 {
		key := string(result[i].([]byte))
		value := string(result[i+1].([]byte))
		data[key] = value
	}
	return data, nil
}

func HGet(key, field string) (string, error) {
	/**
	Redis String Hget操作
	*/
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("hget", key, field))
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis HGet Error")
	}
	return res, err
}

func IncrBy(key string, cnt int) (int64, error) {
	/**
	Redis String IncrBy
	*/
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.Int64(conn.Do("INCRBY", key, cnt))
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis IncrBy Error")
	}
	return res, nil
}

func DecrBy(key string, cnt int) (int64, error) {
	/**
	Redis String IncrBy
	*/
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.Int64(conn.Do("DecrBy", key, cnt))
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis DecrBy Error")
	}
	return res, nil
}

func MDelete(keys ...string) error {
	/**
	Redis String MDelete
	*/
	conn := pool.Get()
	defer conn.Close()
	_ = conn.Send("MULTI")
	for _, v := range keys {
		if len(v) == 0 {
			continue
		}
		conn.Send("DEL", v)
	}
	_, err := conn.Do("EXEC")
	if err != nil {
		log.Logger.Error(err.Error(), "| Redis MDelete Error")
	}
	return err
}
