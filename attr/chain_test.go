package attr

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func TestReplaceAttrChain(t *testing.T) {

	replace_attr := ReplaceAttrChain([]func(groups []string, a slog.Attr) slog.Attr{
		EmojiLevelFunc(),
	})

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       LevelTrace,
		ReplaceAttr: replace_attr,
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
