package main

import (
	"log"

	"github.com/PrajaktaB27/akshar/internal/env"
)

func main() {

	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
	}
	mux := app.mount()
	log.Fatal(app.serve(mux))
}
