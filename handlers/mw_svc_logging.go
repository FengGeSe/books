package handlers

import (
	"context"
	global "myservices/books/global"
	pb "myservices/books/pb/details"
	"time"
)

func LoggingSvcMiddleware() SvcMiddleware {
	return func(next pb.BookDetailsServer) pb.BookDetailsServer {
		return loggingSvcMiddleware{next}
	}
}

type loggingSvcMiddleware struct {
	Next pb.BookDetailsServer
}

func (this loggingSvcMiddleware) Detail(ctx context.Context, in *pb.DetailReq) (resp *pb.DetailResp, err error) {
	defer func(begin time.Time) {
		global.Logger.InfoWithFields(func() *global.LogFields {
			return global.NewLogFields().
				Append("method", "detail").
				Append("input", in).
				Append("err", err).
				Append("duration", time.Since(begin))
		}, "out svc")
	}(time.Now())

	return this.Next.Detail(ctx, in)
}
