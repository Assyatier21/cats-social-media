package redis

//TODO: Implement Redis

// [^^ ONLY BOILERPLATE ^^]

// type RedisClient interface {
// 	Set(ctx context.Context, key string, val []byte, expiration time.Duration) (err error)
// 	Get(ctx context.Context, key string) (result string, err error)
// 	Del(ctx context.Context, key string) (err error)
// 	Ping(ctx context.Context) (err error)
// }

// func (c *client) Set(ctx context.Context, key string, val []byte, expiration time.Duration) (err error) {
// 	key = fmt.Sprintf("%s-%s", c.redisKeyPrefix, key)
// 	err = c.redisClient.Set(ctx, key, string(val), expiration).Err()
// 	return
// }

// func (c *client) Get(ctx context.Context, key string) (result string, err error) {
// 	key = fmt.Sprintf("%s-%s", c.redisKeyPrefix, key)
// 	result, err = c.redisClient.Get(ctx, key).Result()
// 	return
// }

// func (c *client) Del(ctx context.Context, key string) (err error) {
// 	key = fmt.Sprintf("%s-%s", c.redisKeyPrefix, key)
// 	err = c.redisClient.Del(ctx, key).Err()
// 	return
// }

// func (c *client) Ping(ctx context.Context) (err error) {
// 	err = c.redisClient.Ping(ctx).Err()
// 	return
// }
