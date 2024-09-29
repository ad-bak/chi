package main

import (
	"log"

	"github.com/ad-bak/chi/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("API_ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
