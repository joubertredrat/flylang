package main

import (
	"flag"
	"joubertredrat/flylang/cmd/aeronyx"
	"joubertredrat/flylang/shared"
	"log"
)

func main() {
	apiFlag := flag.String("api", "", "Specify the API to run: 'airports' or 'aeronyx'")
	flag.Parse()

	switch *apiFlag {
	case "airports":
		shared.Run()
	case "aeronyx":
		aeronyx.Run()
	default:
		log.Fatal("Invalid or missing --api flag. Use '--api=airports' or '--api=aeronyx'")
	}
}
