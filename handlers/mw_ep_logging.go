package handlers

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	global "myservices/books/global"
	"time"
)

func LoggingEndpointMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				global.Logger.InfoWithFields(func() *global.LogFields {
					return global.NewLogFields().
						Append("headers", ctx.Value("headers")).
						Append("err:", err).
						Append("duration", time.Since(begin))
				}, "out endpoint")
			}(time.Now())

			return next(ctx, request)
		}
	}
}
