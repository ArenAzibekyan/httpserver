package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/ArenAzibekyan/httpserver/v2"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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

	srv := httpserver.New(
		httpserver.WithAddress("", 8080),
		httpserver.WithHandler(r.Handler()),
	)

	err := srv.Run(ctx)
	log.Err(err).Msg("http server stopped")
}
