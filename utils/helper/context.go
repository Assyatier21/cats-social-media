package helper

import (
	"context"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	var ctx = context.Background()
	return context.WithTimeout(ctx, time.Second*60)
}
