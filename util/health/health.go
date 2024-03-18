package health

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"time"
)

func HealthCheckPostgres(db *sql.DB, timeout time.Duration) error {
	pgDbContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := db.PingContext(pgDbContext)
	if err != nil {
		return err
	}

	return nil
}

func HealthCheckRedis(redis *redis.Client, timeout time.Duration) error {
	redisCtx, redisCancel := context.WithTimeout(context.Background(), timeout)
	defer redisCancel()
	err := redis.Ping(redisCtx).Err()
	return err
}
