// cmd/elitetoken/main.go
package main

import (
	"flag"
	"log"
	"os"

	"elitetoken/internal/elitetoken"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := elitetoken.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
