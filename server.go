package httpserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// Server is an interface for http server
type Server interface {
	// ListenAndServe is a wrapper for http.Server.ListenAndServe method, that
	// returns nil when Server is closed. Otherwise, it always returns non-nil error
	ListenAndServe() error

	// Run runs http server and gracefully shuts it down on context cancellation. It
	// works well in conjunction with custom os signals handling via context (signal.NotifyContext)
	Run(ctx context.Context) error
}

// New creates and returns new Server with given options. Default options:
// - read timeout 30s
// - write timeout 30s
// - shutdown timeout 5s
func New(opts ...opt) Server {
	srv := &server{
		httpServer: &http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		shutdownTimeout: 5 * time.Second,
	}

	for _, fn := range opts {
		fn(srv)
	}

	return srv
}

// server implements Server
type server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

// ListenAndServe implements Server.ListenAndServe
func (s *server) ListenAndServe() error {
	err := s.httpServer.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

// stop stops server. At first it tries to shut down server gracefully
// with timeout. If timeout exceeded, it closes the server immediately
func (s *server) stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil && errors.Is(err, context.DeadlineExceeded) {
		return s.httpServer.Close()
	}
	return err
}

// Run implements Server.Run
func (s *server) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// if ListenAndServe returns non-nil error, Go cancels ctx
	g.Go(s.ListenAndServe)

	g.Go(func() error {
		<-ctx.Done()
		return s.stop()
	})

	return g.Wait()
}

// ListenAndServe is a shortcut for New(opts...).ListenAndServe()
func ListenAndServe(opts ...opt) error {
	return New(opts...).ListenAndServe()
}

// Run is a shortcut for New(opts...).Run(ctx)
func Run(ctx context.Context, opts ...opt) error {
	return New(opts...).Run(ctx)
}
