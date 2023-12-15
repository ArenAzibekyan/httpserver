package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/ArenAzibekyan/httpserver/v2"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	r := gin.New()
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return httpserver.Run(ctx, 8080, r.Handler())
	})

	g.Go(func() error {
		return httpserver.Run(ctx, 8081, r.Handler())
	})

	g.Go(func() error {
		return httpserver.Run(ctx, 8082, r.Handler())
	})

	err := g.Wait()
	log.Err(err).Msg("http server stopped")
}
