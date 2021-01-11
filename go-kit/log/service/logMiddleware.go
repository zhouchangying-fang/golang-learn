package service

import (
	kitlog "github.com/go-kit/kit/log"
	"time"
)

type ServiceMiddleware func(Service) Service

type loggingMiddleware struct {
	next   Service
	logger kitlog.Logger
}

func LoggingMiddleware(log kitlog.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next,
			log,
		}
	}
}
func (st loggingMiddleware) Add(a, b int) int {
	defer func(begin time.Time) {
		st.logger.Log(
			"function", "Add",
			"a", a,
			"b", b,
			"took", time.Since(begin),
		)
	}(time.Now())
	return st.next.Add(a, b)
}
