package service

import (
	"log"
	"time"
)

type Middleware func(HelloService) HelloService

func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next HelloService) HelloService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   HelloService
	logger *log.Logger
}

func (mw *loggingMiddleware) Hello(name string) string {
	defer func(begin time.Time) {
		mw.logger.Printf("method=hello, name=%s, took=%v\n", name, time.Since(begin))
	}(time.Now())

	return mw.next.Hello(name)
}
