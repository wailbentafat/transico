package redis

import (

	"fmt"
	"log"
	

	
)


func (rdb *Redisconf)subredis() {
	pubsub := rdb.rdb.Subscribe(rdb.ctx, rdb.channel)
	defer pubsub.Close()
	fmt.Println("Waiting for messages...")
	val, err := rdb.rdb.Get(rdb.ctx, "key").Result()
	  if err != nil {
		log.Fatalf("could not get key: %v", err)
	  }
	  fmt.Println(val)
	  for {
		msg, err := pubsub.ReceiveMessage(rdb.ctx)
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		fmt.Printf("Received message: %s\n", msg.Payload)
		//get array of messages 
	}
	
	
}

