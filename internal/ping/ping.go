package ping

import (
	"context"
	"fmt"
	"time"
)

func Run(ctx context.Context, verbose bool) error {
	fmt.Println("Starting ping... (press Ctrl+C)")

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("pong")
		return nil
	case <-ctx.Done():
		fmt.Println("Ping cancelled")
		return ctx.Err()
	}

}
