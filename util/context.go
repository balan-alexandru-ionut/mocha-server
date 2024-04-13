package util

import (
	"context"
	"time"
)

func NewBackgroundContext(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}
