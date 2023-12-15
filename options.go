package httpserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type opt func(*Server)

func WithAddress(host string, port uint16) opt {
	return func(s *Server) {
		s.httpServer.Addr = fmt.Sprintf("%s:%d", host, port)
	}
}

func WithHandler(handler http.Handler) opt {
	return func(s *Server) {
		s.httpServer.Handler = handler
	}
}

func WithTLSConfig(conf *tls.Config) opt {
	return func(s *Server) {
		s.httpServer.TLSConfig = conf
	}
}

func WithReadTimeout(d time.Duration) opt {
	return func(s *Server) {
		s.httpServer.ReadTimeout = d
	}
}

func WithReadHeaderTimeout(d time.Duration) opt {
	return func(s *Server) {
		s.httpServer.ReadHeaderTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) opt {
	return func(s *Server) {
		s.httpServer.WriteTimeout = d
	}
}

func WithIdleTimeout(d time.Duration) opt {
	return func(s *Server) {
		s.httpServer.IdleTimeout = d
	}
}

func WithMaxHeaderBytes(i int) opt {
	return func(s *Server) {
		s.httpServer.MaxHeaderBytes = i
	}
}

func WithShutdownTimeout(d time.Duration) opt {
	return func(s *Server) {
		s.shutdownTimeout = d
	}
}
