package service

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"time"
)

type ServiceMiddleware func(service Service) Service

type loggingMiddleware struct {
	Service
	log.Logger
}

func LoggingMiddleware(log log.Logger) ServiceMiddleware {
	return func(service Service) Service {
		return loggingMiddleware{
			service,
			log,
		}
	}
}
func (st loggingMiddleware) Add(a, b int) int {
	defer func(begin time.Time) {
		fmt.Println("------------------")
	}(time.Now())
	return st.Service.Add(a, b)
}
