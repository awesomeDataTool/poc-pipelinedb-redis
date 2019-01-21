package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		sub := client.Subscribe("products")
		_, err := sub.Receive()
		if err != nil {
			log.Fatal(err)
		}

		ch := sub.Channel()

		for msg := range ch {
			fmt.Println(msg.Channel, msg.Payload)
		}
	}()

	group.Wait()
}
