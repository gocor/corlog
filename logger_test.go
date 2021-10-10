package corlog

import (
	"context"
	"testing"

	"github.com/gocor/corlog/logcfg"
)

func TestLogger(t *testing.T) {
	logcfg.InitConfig(logcfg.Config{
		IsDebug: true,
		App:     "UnitTestApp",
		Env:     "local",
	})

	ctx := context.Background()
	log := New(ctx)
	log.Info("Test Info Message")
	log.Warn("Test Warn Message")
	log.Error("Test Error Message")
}
