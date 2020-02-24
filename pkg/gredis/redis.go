/*
 * @Description: redis
 * @Author: leo
 * @Date: 2020-02-24 12:18:12
 * @LastEditors: leo
 * @LastEditTime: 2020-02-24 20:16:43
 */

package gredis

import (
	"encoding/json"
	"time"

	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,     // 最大空闲连接数
		MaxActive:   setting.RedisSetting.MaxActive,   // 在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		IdleTimeout: setting.RedisSetting.IdleTimeout, // 在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
		Dial: func() (redis.Conn, error) { // 提供创建和配置应用程序连接的一个函数
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 可选的应用程序检查健康功能
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

// Set 设置 key value
func Set(key string, data interface{}, time int) error {
	// 在连接池中获取一个活跃连接
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// 向 Redis 服务器发送命令并返回收到的答复
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断是否存在
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	// 将命令返回转为布尔值
	_, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return true
}

// Get 根据key获取值
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	// 将命令返回转为 Bytes
	value, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return value, nil
}

// Delete 根据key 删除值
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes 模糊删除
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	// 将命令返回转为 []string
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
