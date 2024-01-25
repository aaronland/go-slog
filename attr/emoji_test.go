package attr

import (
	"context"
	"testing"
	"log/slog"
	"os"
)

func TestEmojiLevelFunc(t *testing.T) {

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: LevelTrace,			
		ReplaceAttr: EmojiLevelFunc(),
	})
			
	logger := slog.New(th)
	ctx := context.Background()
	logger.Log(ctx, LevelEmergency, "missing pilots")
	logger.Error("failed to start engines", "err", "missing fuel")
	logger.Warn("falling back to default value")
	logger.Log(ctx, LevelNotice, "all systems are running")
	logger.Info("initiating launch")
	logger.Debug("starting background job")
	logger.Log(ctx, LevelTrace, "button clicked")
}

