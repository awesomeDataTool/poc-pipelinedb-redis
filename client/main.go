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

<<<<<<< HEAD
	group := sync.WaitGroup{}
=======
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	group := sync.WaitGroup{}

>>>>>>> 4a01eab6ea9c9855b2997721e53e74d05af410c7
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
