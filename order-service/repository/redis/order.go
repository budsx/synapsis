package redis

import (
	"context"
	"fmt"
	"time"
)

func (r *RedisRepository) DeduplicateCreateOrder(ctx context.Context, idempotencyKey string) (bool, error) {
	key := fmt.Sprintf("deduplicate:order:%s", idempotencyKey)
	return r.redis.SetNX(key, idempotencyKey, 24*time.Hour).Result()
}
