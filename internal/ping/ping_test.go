package ping

import (
	"context"
	"testing"
	"time"
)

func TestRun_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := Run(ctx, false)
	if err == nil {
		t.Fatal("expected context cancellation error")
	}
}
