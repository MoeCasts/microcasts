package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NovelsService) NovelsService

type loggingMiddleware struct {
	logger log.Logger
	next   NovelsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NovelsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NovelsService) NovelsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Browse(ctx context.Context, request interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Browse", "request", request, "rs", fmt.Sprintf("%+v", rs), "err", err)
	}()
	return l.next.Browse(ctx, request)
}
func (l loggingMiddleware) Read(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Read", "pk", pk, "request", request, "rs", rs, "err", err)
	}()
	return l.next.Read(ctx, pk, request)
}
func (l loggingMiddleware) Add(ctx context.Context, request interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Add", "request", request, "rs", rs, "err", err)
	}()
	return l.next.Add(ctx, request)
}
func (l loggingMiddleware) Edit(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Edit", "pk", pk, "request", request, "rs", rs, "err", err)
	}()
	return l.next.Edit(ctx, pk, request)
}
func (l loggingMiddleware) Trash(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Trash", "pk", pk, "rs", rs, "err", err)
	}()
	return l.next.Trash(ctx, pk)
}
func (l loggingMiddleware) Restore(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Restore", "pk", pk, "rs", rs, "err", err)
	}()
	return l.next.Restore(ctx, pk)
}
func (l loggingMiddleware) Destroy(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "Destroy", "pk", pk, "rs", rs, "err", err)
	}()
	return l.next.Destroy(ctx, pk)
}
func (l loggingMiddleware) BatchAdd(ctx context.Context, request interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "BatchAdd", "request", request, "rs", rs, "err", err)
	}()
	return l.next.BatchAdd(ctx, request)
}
func (l loggingMiddleware) BatchEdit(ctx context.Context, request interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "BatchEdit", "request", request, "rs", rs, "err", err)
	}()
	return l.next.BatchEdit(ctx, request)
}
func (l loggingMiddleware) BatchTrash(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "BatchTrash", "pks", pks, "rs", rs, "err", err)
	}()
	return l.next.BatchTrash(ctx, pks)
}
func (l loggingMiddleware) BatchRestore(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "BatchRestore", "pks", pks, "rs", rs, "err", err)
	}()
	return l.next.BatchRestore(ctx, pks)
}
func (l loggingMiddleware) BatchDestroy(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	defer func() {
		l.logger.Log("method", "BatchDestroy", "pks", pks, "rs", rs, "err", err)
	}()
	return l.next.BatchDestroy(ctx, pks)
}
