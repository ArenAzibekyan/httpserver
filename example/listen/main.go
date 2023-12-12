package main

import (
	"net/http"

	"github.com/ArenAzibekyan/httpserver"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	r := gin.New()
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})

	err := httpserver.ListenAndServe(
		httpserver.WithAddress("", 8080),
		httpserver.WithHandler(r.Handler()),
	)
	log.Err(err).Msg("http server stopped")
}
