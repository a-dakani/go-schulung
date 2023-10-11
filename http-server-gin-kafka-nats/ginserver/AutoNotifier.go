package ginserver

import "golang.org/x/net/context"

type AutoNotifier interface {
	NewAutoCreated(ctx context.Context, auto Auto) error
}
