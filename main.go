package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"joubertredrat/flylang/cmd/americandes"
	"joubertredrat/flylang/cmd/avionca"
	"joubertredrat/flylang/cmd/copana"
	"joubertredrat/flylang/cmd/latangos"
	"joubertredrat/flylang/cmd/skylux"

	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		name   string
		runner func(ctx context.Context) error
	}
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	servers := []Server{
		{
			name:   "americandes",
			runner: americandes.Run,
		},
		{
			name:   "avionca",
			runner: avionca.Run,
		},
		{
			name:   "copana",
			runner: copana.Run,
		},
		{
			name:   "latangos",
			runner: latangos.Run,
		},
		{
			name:   "skylux",
			runner: skylux.Run,
		},
	}

	for _, server := range servers {
		go func(name string, runner func(ctx context.Context) error) {
			if err := runner(ctx); err != nil {
				log.Printf("Error on running server %s: %v", name, err)
			}
		}(server.name, server.runner)
	}

	<-quit
	log.Println("Stopping main...")
	cancel()

	time.Sleep(1 * time.Second)
	log.Println("Finished.")
}
