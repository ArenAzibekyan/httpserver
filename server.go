package httpserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// Server is a http server
type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

// New creates new Server with given options. Default options:
// - read timeout 30s
// - write timeout 30s
// - shutdown timeout 15s
func New(opts ...Opt) *Server {
	srv := &Server{
		httpServer: &http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		shutdownTimeout: 15 * time.Second,
	}

	for _, fn := range opts {
		fn(srv)
	}

	return srv
}

// Run runs http server and gracefully shuts it down on context cancellation. It
// works well in conjunction with custom os signals handling via context (signal.NotifyContext)
func (s *Server) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// if listenAndServe returns non-nil error, ctx will be canceled
	g.Go(s.listenAndServe)

	g.Go(func() error {
		<-ctx.Done()
		return s.stop()
	})

	return g.Wait()
}

// listenAndServe is a wrapper for http.Server.ListenAndServe method, that
// returns nil when Server is closed. Otherwise, it always returns non-nil error
func (s *Server) listenAndServe() error {
	err := s.httpServer.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

// stop stops server. At first it tries to shut down server gracefully
// with timeout. If timeout exceeded, it closes the server immediately
func (s *Server) stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil && errors.Is(err, context.DeadlineExceeded) {
		return s.httpServer.Close()
	}
	return err
}

// Run is a shortcut for New(...).Run(ctx). Port and handler parameters
// take precedence over the WithAddress and WithHandler options
// respectively if they are passed too
func Run(ctx context.Context, port uint16, handler http.Handler, opts ...Opt) error {
	opts = append(opts,
		WithAddress("", port),
		WithHandler(handler),
	)

	return New(opts...).Run(ctx)
}
