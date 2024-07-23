package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// 连接到Redis服务器
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 没有密码时留空
		DB:       0,                // 默认DB
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("无法连接到Redis:", err)
		return
	}
	fmt.Println("连接成功:", pong)

	// 设置键值对
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("设置键值对失败:", err)
		return
	}
	fmt.Println("键值对设置成功")

	// 获取值
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("获取值失败:", err)
		return
	}
	fmt.Println("获取到的值:", val)

	// 其他Redis操作
	// 自增操作
	err = rdb.Incr(ctx, "counter").Err()
	if err != nil {
		fmt.Println("自增操作失败:", err)
		return
	}
	counter, err := rdb.Get(ctx, "counter").Result()
	if err != nil {
		fmt.Println("获取自增值失败:", err)
		return
	}
	fmt.Println("自增后的值:", counter)

	// 列表操作
	err = rdb.RPush(ctx, "list", "item1").Err()
	if err != nil {
		fmt.Println("列表操作失败:", err)
		return
	}
	err = rdb.RPush(ctx, "list", "item2").Err()
	if err != nil {
		fmt.Println("列表操作失败:", err)
		return
	}
	list, err := rdb.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		fmt.Println("获取列表值失败:", err)
		return
	}
	fmt.Println("列表中的值:", list)
}
