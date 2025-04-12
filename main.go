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
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := latangos.Run(ctx); err != nil {
			log.Printf("Error on running server latangos: %v", err)
		}
	}()
	go func() {
		if err := americandes.Run(ctx); err != nil {
			log.Printf("Error on running server americandes: %v", err)
		}
	}()
	go func() {
		if err := copana.Run(ctx); err != nil {
			log.Printf("Error on running server copana: %v", err)
		}
	}()
	go func() {
		if err := skylux.Run(ctx); err != nil {
			log.Printf("Error on running server skylux: %v", err)
		}
	}()
	go func() {
		if err := avionca.Run(ctx); err != nil {
			log.Printf("Error on running server avionca: %v", err)
		}
	}()

	<-quit
	log.Println("Stopping...")
	cancel()

	time.Sleep(1 * time.Second)
	log.Println("Stopped.")
}
