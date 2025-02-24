package redis

import "fmt"

func (rdb *Redisconf) pubredis(message string) {
	err:=rdb.rdb.Publish(rdb.ctx, rdb.channel, message).Err()
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return
	}
}