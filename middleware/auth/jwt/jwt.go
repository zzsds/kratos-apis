package jwt

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

// Option is tracing option.
type Option func(*options)

type options struct {
	logger log.Logger
}

// WithLogger with recovery logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Server returns a new server middleware for OpenTelemetry.
func Server(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}

	log := log.NewHelper("middleware/jwt", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Errorf("fdasfdsafds a---这个是测试")
			reply, err = handler(ctx, req)
			return
		}
	}
}
