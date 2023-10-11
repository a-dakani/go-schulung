package Logging

import (
	"context"
	"github.com/sirupsen/logrus"
)

func Logger(ctx context.Context) *logrus.Entry {
	logger := logrus.StandardLogger()
	entry := logrus.NewEntry(logger)

	if v := ctx.Value("x-trace-id"); v != "" {
		entry = entry.WithField("x-trace-id", ctx.Value("x-trace-id"))
	}
	return entry
}
