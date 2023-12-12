package httpserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type opt func(*server)

func WithAddress(host string, port uint16) opt {
	return func(s *server) {
		s.httpServer.Addr = fmt.Sprintf("%s:%d", host, port)
	}
}

func WithHandler(handler http.Handler) opt {
	return func(s *server) {
		s.httpServer.Handler = handler
	}
}

func WithTLSConfig(conf *tls.Config) opt {
	return func(s *server) {
		s.httpServer.TLSConfig = conf
	}
}

func WithReadTimeout(d time.Duration) opt {
	return func(s *server) {
		s.httpServer.ReadTimeout = d
	}
}

func WithReadHeaderTimeout(d time.Duration) opt {
	return func(s *server) {
		s.httpServer.ReadHeaderTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) opt {
	return func(s *server) {
		s.httpServer.WriteTimeout = d
	}
}

func WithIdleTimeout(d time.Duration) opt {
	return func(s *server) {
		s.httpServer.IdleTimeout = d
	}
}

func WithMaxHeaderBytes(i int) opt {
	return func(s *server) {
		s.httpServer.MaxHeaderBytes = i
	}
}

func WithShutdownTimeout(d time.Duration) opt {
	return func(s *server) {
		s.shutdownTimeout = d
	}
}
