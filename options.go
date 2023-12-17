package httpserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type Opt func(*Server)

func WithAddress(host string, port uint16) Opt {
	return func(s *Server) {
		s.httpServer.Addr = fmt.Sprintf("%s:%d", host, port)
	}
}

func WithHandler(handler http.Handler) Opt {
	return func(s *Server) {
		s.httpServer.Handler = handler
	}
}

func WithTLSConfig(conf *tls.Config) Opt {
	return func(s *Server) {
		s.httpServer.TLSConfig = conf
	}
}

func WithReadTimeout(d time.Duration) Opt {
	return func(s *Server) {
		s.httpServer.ReadTimeout = d
	}
}

func WithReadHeaderTimeout(d time.Duration) Opt {
	return func(s *Server) {
		s.httpServer.ReadHeaderTimeout = d
	}
}

func WithWriteTimeout(d time.Duration) Opt {
	return func(s *Server) {
		s.httpServer.WriteTimeout = d
	}
}

func WithIdleTimeout(d time.Duration) Opt {
	return func(s *Server) {
		s.httpServer.IdleTimeout = d
	}
}

func WithMaxHeaderBytes(i int) Opt {
	return func(s *Server) {
		s.httpServer.MaxHeaderBytes = i
	}
}

func WithShutdownTimeout(d time.Duration) Opt {
	return func(s *Server) {
		s.shutdownTimeout = d
	}
}
