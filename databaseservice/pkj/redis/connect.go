package redis

import (
	"context"
	"databaseservice/pkj/utils"


	"github.com/redis/go-redis/v9"
	
)
type Redisconf struct {
	rdb *redis.Client
	ctx context.Context
	channel string
	logger *utils.MyLogger
	

}

func NewRedisconfig(redisAddress string, channel string, logger *utils.MyLogger) (*Redisconf, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.LogError().Msg("Failed to connect to Redis")
		return nil, err
	}

	logger.LogInfo().Msg("Connected to Redis successfully")
	return &Redisconf{rdb: rdb, ctx: ctx, channel: channel, logger: logger}, nil
}
