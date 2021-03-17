package auth

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/zzsds/kratos-apis/auth"
)

var base auth.Auth

func NewAuth(a auth.Auth) {
	base = a
}

// Option is tracing option.
type Option func(*options)

type options struct {
	auth    auth.Auth
	logger  log.Logger
	header  string
	prefix  string
	exclude []string
}

// WithAuth with auth interface{}.
func WithAuth(auth auth.Auth) Option {
	return func(o *options) {
		o.auth = auth
	}
}

// WithLogger with recovery logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithHeader with recovery header.
func WithHeader(header string) Option {
	return func(o *options) {
		o.header = header
	}
}

// WithPrefix with recovery prefix.
func WithPrefix(prefix string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

// WithExclude with recovery exclude.
func WithExclude(exclude ...string) Option {
	return func(o *options) {
		o.exclude = exclude
	}
}

// Server returns a new server middleware for OpenTelemetry.
func Server(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
		header: "Authorization",
		prefix: auth.BearerScheme,
		auth:   base,
	}
	for _, o := range opts {
		o(&options)
	}

	_ = log.NewHelper("middleware/auth", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				path      string
				method    string
				params    string
				component string
			)

			if info, ok := http.FromServerContext(ctx); ok {
				method = info.Request.Method
				request := info.Request
				_, ok := request.Header[options.header]
				if !ok {
					return nil, errors.Unauthorized("Unauthorized", "Header 参数名 %s 不存在", options.header)
				}
				header := request.Header.Get(options.header)
				if !strings.HasPrefix(header, options.prefix) {
					return nil, errors.Unauthorized("Unauthorized", "invalid authorization header. expected Bearer schema")
				}
				if options.auth == nil {
					return nil, errors.DataLoss("Auth Nut initialization ", "token parsing failed")
				}

				account, err := options.auth.Inspect(strings.TrimPrefix(header, options.prefix))
				if err != nil {
					return nil, errors.Unauthorized("Unauthorized", "token parsing failed")
				}
				ctx = auth.ContextWithAccount(ctx, account)
			} else if info, ok := grpc.FromServerContext(ctx); ok {
				path = info.FullMethod
				method = "POST"
				fmt.Println(info.FullMethod, "GRPC")
			}

			fmt.Println(path, method, params, component)
			reply, err = handler(ctx, req)
			return
		}
	}
}

func GetID(ctx context.Context) int32 {
	var id int
	account, ok := auth.AccountFromContext(ctx)
	if ok {
		id, _ = strconv.Atoi(account.ID)
	}
	return int32(id)
}
