package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := rds.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		ctx, cacnel := context.WithTimeout(context.Background(), time.Second)
		defer cacnel()
		sub := rds.Subscribe(ctx, "test-1")
		_, err := sub.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		ch := sub.Channel()
		for msg := range ch {
			log.Println(msg.Channel, msg.Pattern, msg.Payload)
		}
	}()

	go func() {
		ctx, cacnel := context.WithTimeout(context.Background(), time.Second)
		defer cacnel()
		sub := rds.Subscribe(ctx, "test-1")
		_, err := sub.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		ch := sub.Channel()
		for msg := range ch {
			log.Println(msg.Channel, msg.Pattern, msg.Payload)
		}
	}()

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		err := rds.Publish(context.Background(), "test-1", time.Now().Unix()).Err()
		if err != nil {
			log.Println(err)
		}
	}
}
